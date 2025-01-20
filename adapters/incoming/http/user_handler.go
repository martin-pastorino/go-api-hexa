package http

import (
	"api/adapters/dtos"
	core_errors "api/core/errors"
	"api/core/ports/incoming"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-chi/render"
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
	var userRequest dtos.User

	if err := render.Bind(r, &userRequest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userID, err := uh.userService.CreateUser(r.Context(), userRequest.ToUserDomainModel())
	if err != nil {
		var alreadyExists *core_errors.AlreadyExists
		if errors.As(err, &alreadyExists) {
			http.Error(w, err.Error(), alreadyExists.Code)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"id": userID})
}

// // GetUser godoc
func (uh *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")

	user, err := uh.userService.GetUser(r.Context(), email)

	if err != nil || user.ID == "" {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(dtos.ToUser(user))
}

func (uh *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")

	result, err := uh.userService.DeleteUser(r.Context(), email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	type Result struct {
		Email string `json:"email"`
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Result{Email: result})
}

func (uh *UserHandler) Search(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")

	users, err := uh.userService.Search(r.Context(), email)

	if err != nil || len(users) == 0 {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(dtos.ToUsers(users))
}