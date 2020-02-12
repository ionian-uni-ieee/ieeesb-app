package rest

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func (h *Handler) PostRegister(w http.ResponseWriter, r *http.Request) {
	type body struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Email    string `json:"email"`
		Fullname string `json:"fullname"`
	}

	b := body{}

	err := json.NewDecoder(r.Body).Decode(&b)

	if err != nil {
		sendError(w, 400, httpError{"Bad body format", "Body format is not in the proper shape"})
		return
	}

	invalidParams := httpErrorInvalidParams{InvalidParams: map[string]string{}}

	if b.Username == "" {
		varName := "username"
		invalidParams.InvalidParams[varName] = "Empty string"
	}

	if b.Password == "" {
		varName := "password"
		invalidParams.InvalidParams[varName] = "Empty string"
	}

	if b.Email == "" {
		varName := "email"
		invalidParams.InvalidParams[varName] = "Empty string"
	}

	if b.Fullname == "" {
		varName := "fullname"
		invalidParams.InvalidParams[varName] = "Empty string"
	}

	if len(invalidParams.InvalidParams) > 0 {
		invalidParams.Title = "Invalid parameters"
		invalidParams.Details = "Some parameters were either empty or invalid"
		sendError(w, 400, invalidParams)
		return
	}

	userID, err := h.authController.Register(b.Username, b.Password, b.Email, b.Fullname)

	if err != nil {
		sendError(w, 400, httpError{"Bad Request", err.Error()})
		return
	}

	sendJSON(w, 200, userID)
}

func (h *Handler) PostLogin(w http.ResponseWriter, r *http.Request) {
	type body struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	b := body{}

	err := json.NewDecoder(r.Body).Decode(&b)

	if err != nil {
		sendError(w, 400, httpError{"Bad body format", "Body format is not in the proper shape"})
		return
	}

	invalidParams := httpErrorInvalidParams{InvalidParams: map[string]string{}}

	if b.Username == "" {
		varName := "username"
		invalidParams.InvalidParams[varName] = "Empty string"
	}

	if b.Password == "" {
		varName := "password"
		invalidParams.InvalidParams[varName] = "Empty string"
	}

	if len(invalidParams.InvalidParams) > 0 {
		invalidParams.Title = "Invalid parameters"
		invalidParams.Details = "Some parameters were either empty or invalid"
		sendError(w, 400, invalidParams)
		return
	}

	sessionID, err := h.authController.Login(b.Username, b.Password)

	if err != nil {
		sendError(w, 400, httpError{"Bad request", err.Error()})
		return
	}

	cookieValue := map[string]string{
		"sessionID": sessionID,
		"username":  b.Username,
		"Expires":   strconv.Itoa(30 * 24 * 60 * 60 * 1000),
	}
	setCookie(w, cookieValue)

	sendJSON(w, 200, "Logged in")
}

func (h *Handler) PostLogout(w http.ResponseWriter, r *http.Request) {
	sessionID := getCookieValue(r, "sessionID")

	err := h.authController.Logout(sessionID)

	if err != nil {
		sendError(w, 400, httpError{"Bad request", err.Error()})
		return
	}

	clearCookie(w)
	sendJSON(w, 200, "Logged out")
}

func (h *Handler) GetProfile(w http.ResponseWriter, r *http.Request) {
	sessionID := getCookieValue(r, "sessionID")

	user, err := h.authController.Profile(sessionID)

	if err != nil {
		sendError(w, 400, httpError{"Not logged in", "No stored session exists that corresponds to your sessionID"})
		return
	}

	sendJSON(w, 200, user)
}
