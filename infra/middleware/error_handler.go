package middleware

import (
	core_errors "api/core/errors"
	"errors"
	"net/http"

	"github.com/go-chi/render"
)

func ErrorHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				var appErr *core_errors.GenericError
				var alreadyExists *core_errors.AlreadyExists

				if errors.As(err.(error), &appErr) {
					http.Error(w, appErr.Error(), appErr.Code)
					return
				}

				if errors.As(err.(error), &alreadyExists) {
					http.Error(w, alreadyExists.Error(), alreadyExists.Code)
					return
				}

				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

type errResponse struct {
	Error string `json:"error"`
	Code  int    `json:"code"`
}

func (e *errResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.Code)
	return nil
}

func ErrInvalidRequest(err error) render.Renderer {
	return &errResponse{
		Error: err.Error(),
		Code:  400,
	}
}

func ErrAlreadyExists(err error) render.Renderer {
	return &errResponse{
		Error: err.Error(),
		Code:  409,
	}
}

func ErrInternalServer(err error) render.Renderer {
	return &errResponse{
		Error: err.Error(),
		Code:  500,
	}
}
