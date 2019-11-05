package http

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["user_id"]

	invalidParams := httpErrorInvalidParams{InvalidParams: map[string]string{}}

	if _, err := primitive.ObjectIDFromHex(userID); err != nil {
		varName := "user_id"
		invalidParams.InvalidParams[varName] = "Invalid ObjectID"
	}

	if userID == "" {
		varName := "user_id"
		invalidParams.InvalidParams[varName] = "Empty string"
	}

	if len(invalidParams.InvalidParams) > 0 {
		invalidParams.Title = "Invalid parameters"
		invalidParams.Details = "Some parameters were either empty or invalid"
		sendError(w, 400, invalidParams)
		return
	}

	err := h.usersController.Delete(userID)

	if err != nil {
		sendError(w, 400, httpError{"Bad request", err.Error()})
		return
	}

	sendJSON(w, 200, "User removed")
}

func (h *Handler) PutUserPassword(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["user_id"]

	type body struct {
		OldPassword string `json:"old_password"`
		NewPassword string `json:"new_password"`
	}

	b := body{}

	err := json.NewDecoder(r.Body).Decode(&b)

	if err != nil {
		sendError(w, 400, httpError{"Bad body format", "Body format is not in the proper shape"})
		return
	}

	invalidParams := httpErrorInvalidParams{InvalidParams: map[string]string{}}

	if _, err := primitive.ObjectIDFromHex(userID); err != nil {
		varName := "user_id"
		invalidParams.InvalidParams[varName] = "Invalid ObjectID"
	}

	if userID == "" {
		varName := "user_id"
		invalidParams.InvalidParams[varName] = "Empty string"
	}

	if b.OldPassword == "" {
		varName := "old_password"
		invalidParams.InvalidParams[varName] = "Empty string"
	}

	if b.NewPassword == "" {
		varName := "new_password"
		invalidParams.InvalidParams[varName] = "Empty string"
	}

	if len(invalidParams.InvalidParams) > 0 {
		invalidParams.Title = "Invalid parameters"
		invalidParams.Details = "Some parameters were either empty or invalid"
		sendError(w, 400, invalidParams)
		return
	}

	err = h.usersController.ChangePassword(userID, b.OldPassword, b.NewPassword)

	if err != nil {
		sendError(w, 400, httpError{"Bad request", err.Error()})
		return
	}

	sendJSON(w, 200, "Password changed")
}

func (h *Handler) PutUser(w http.ResponseWriter, r *http.Request) {

}
