package http

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/securecookie"
	"gitlab.com/gphub/app/internal/app/controllers/collections"
	"gitlab.com/gphub/app/internal/app/controllers/users"
	"gitlab.com/gphub/app/internal/app/repositories"
)

type httpError struct {
	Title   string `json:"title"`
	Details string `json:"details"`
}

type httpErrorInvalidParams struct {
	httpError
	InvalidParams map[string]string `json:"invalid_params"`
}

type jsonResponse struct {
	Data interface{} `json:"data"`
}

type errorResponse struct {
	Error interface{} `json:"error"`
}

type Handler struct {
	usersController       *users.UsersController
	collectionsController *collections.CollectionsController
}

func MakeNewHandler(repositories *repositories.Repositories) *Handler {
	return &Handler{
		users.MakeNewUsersController(repositories),
		collections.MakeNewCollectionsController(repositories),
	}
}

func (h *Handler) GetPing(w http.ResponseWriter, r *http.Request) {
	sendJSON(w, 200, "Pong")
}

func sendJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	res := jsonResponse{Data: data}
	json.NewEncoder(w).Encode(res)
}

func sendError(w http.ResponseWriter, statusCode int, err interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	res := errorResponse{Error: err}
	json.NewEncoder(w).Encode(res)
}

var hashKey = []byte("very-secret")
var blockKey = []byte("a-lot-secret")
var cookieHandler = securecookie.New(
	hashKey,
	nil)

func setCookie(w http.ResponseWriter, values map[string]string) {
	if encoded, err := cookieHandler.Encode("cookie", values); err == nil {
		cookie := &http.Cookie{
			Name:  "cookie",
			Value: encoded,
			Path:  "/",
		}
		http.SetCookie(w, cookie)
	}
}

func getCookieValue(r *http.Request, valueName string) (value string) {
	if cookie, err := r.Cookie("cookie"); err == nil {
		cookieValue := make(map[string]string)
		if err = cookieHandler.Decode("cookie", cookie.Value, &cookieValue); err == nil {
			value = cookieValue[valueName]
		}
	}

	return value
}

func clearCookie(w http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   "cookie",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(w, cookie)
}
