package handlers

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gorilla/securecookie"
)

// HTTPError is a structure for an http error
type HTTPError struct {
	Title   string `json:"title"`
	Details string `json:"details"`
}

// NewHTTPError returns a new instance of a http error
func NewHTTPError(title string, details string) HTTPError {
	return HTTPError{
		Title:   title,
		Details: details,
	}
}

// HTTPErrorInvalidParams is a structure for an http invalid parameters error
type HTTPErrorInvalidParams struct {
	HTTPError
	InvalidParams map[string]string `json:"invalid_params"`
}

// NewHTTPErrorInvalidParams returns an error for invalid body parameters
func NewHTTPErrorInvalidParams(invalidParams map[string]string) HTTPErrorInvalidParams {
	return HTTPErrorInvalidParams{
		HTTPError: HTTPError{
			Title:   "Invalid parameters",
			Details: "Some parameters were invalid. Check for empty fields or invalid ones",
		},
		InvalidParams: invalidParams,
	}
}

// NewBadRequest returns an http error for a bad request
func NewBadRequest(details string) HTTPError {
	return HTTPError{
		Title:   "Bad Request",
		Details: details,
	}
}

// JSONResponse is a structure for a json response
type JSONResponse struct {
	Data interface{} `json:"data"`
}

// ErrorResponse is a structure for an error response
type ErrorResponse struct {
	Error interface{} `json:"error"`
}

// SendJSON responds with a json message to the http client
func SendJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	// w.Header().Add("Access-Control-Allow-Origin", "http://localhost")
	// w.Header().Add("Access-Control-Allow-Credentials", "true")
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	res := JSONResponse{Data: data}
	json.NewEncoder(w).Encode(res)
}

// SendError responds with a json message describing an error to the http error
func SendError(w http.ResponseWriter, statusCode int, err interface{}) {
	w.Header().Add("Content-Type", "application/json")
	// w.Header().Add("Access-Control-Allow-Origin", "http://localhost")
	// w.Header().Add("Access-Control-Allow-Credentials", "true")
	w.WriteHeader(statusCode)
	res := ErrorResponse{Error: err}
	json.NewEncoder(w).Encode(res)
}

var hashSecret = os.Getenv("HASH_SECRET")
var blockSecret = os.Getenv("BLOCK_SECRET")
var hashKey = []byte(hashSecret)
var blockKey = []byte(blockSecret)
var cookieHandler = securecookie.New(
	hashKey,
	nil)

// SetCookie sends a set header cookie to the client
func SetCookie(w http.ResponseWriter, values map[string]string) {
	if encoded, err := cookieHandler.Encode("cookie", values); err == nil {
		cookie := &http.Cookie{
			Name:  "cookie",
			Value: encoded,
			Path:  "/",
		}
		http.SetCookie(w, cookie)
	}
}

// GetCookieValue returns a cookie value from the client's cookie
func GetCookieValue(r *http.Request, valueName string) (value string) {
	if cookie, err := r.Cookie("cookie"); err == nil {
		cookieValue := make(map[string]string)
		if err = cookieHandler.Decode("cookie", cookie.Value, &cookieValue); err == nil {
			value = cookieValue[valueName]
		}
	}

	return value
}

// ClearCookie clears the client's cookie
func ClearCookie(w http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   "cookie",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(w, cookie)
}
