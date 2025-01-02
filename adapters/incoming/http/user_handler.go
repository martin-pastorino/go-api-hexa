package http

import (
	"api/core/ports/incoming"
	"encoding/json"
	"net/http"
)

type UserHandler struct {
	userService incoming.UserService
}

func NewUserHandler(userService incoming.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

// Provider for UserHandler
func NewUserHandlerProvider(userService incoming.UserService) *UserHandler {
	return NewUserHandler(userService)
}

// // CreateUser godoc
func (uh *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var userRequest struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	ctx := r.Context()

	err := json.NewDecoder(r.Body).Decode(&userRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userID, err := uh.userService.CreateUser(ctx, userRequest.Name, userRequest.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"id": userID})
}

// // GetUser godoc
func (uh *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")
	ctx := r.Context()

	user, err := uh.userService.GetUser(ctx, email)
	if err != nil || user.ID == "" {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func (uh *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")
	ctx := r.Context()
	err := uh.userService.DeleteUser(ctx, email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
