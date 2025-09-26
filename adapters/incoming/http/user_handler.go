package http

import (
	"api/adapters/dtos"
	core_errors "api/core/errors"
	"api/core/ports/incoming"
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
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	userID, err := uh.userService.CreateUser(r.Context(), userRequest.ToUserDomainModel())
	if err != nil {
		var alreadyExists *core_errors.AlreadyExists
		if errors.As(err, &alreadyExists) {
			render.Render(w, r, ErrAlreadyExists(err))
			return
		}
		render.Render(w, r, ErrInternalServer(err))
		return
	}

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, map[string]string{"id": userID})
}

// // GetUser godoc
func (uh *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")

	user, err := uh.userService.GetUser(r.Context(), email)

	if err != nil {
		render.Render(w, r, ErrInternalServer(err))
		return
	}

	if user.ID == "" {
		render.Render(w, r, ErrNotFound(errors.New("user not found")))
		return
	}

	render.JSON(w, r, dtos.ToUser(user))
}

func (uh *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")

	result, err := uh.userService.DeleteUser(r.Context(), email)
	if err != nil {
		render.Render(w, r, ErrInternalServer(err))
		return
	}

	type Result struct {
		Email string `json:"email"`
	}

	render.JSON(w, r, Result{Email: result})
}

func (uh *UserHandler) Search(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")

	users, err := uh.userService.Search(r.Context(), email)

	if err != nil {
		render.Render(w, r, ErrInternalServer(err))
		return
	}

	if len(users) == 0 {
		render.Render(w, r, ErrNotFound(errors.New("no users found")))
		return
	}

	render.JSON(w, r, dtos.ToUsers(users))
}
