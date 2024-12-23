package http

import (
	"encoding/json"
	"net/http"
	"api/core/ports/incoming"

)

type UserHandler struct {
	userService incoming.UserService
}

func NewUserHandler(userService incoming.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (uh *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var userRequest struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	err := json.NewDecoder(r.Body).Decode(&userRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userID, err := uh.userService.CreateUser(userRequest.Name, userRequest.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"id": userID})
}