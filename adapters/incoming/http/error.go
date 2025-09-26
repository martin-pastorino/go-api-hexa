package http

import (
	"github.com/go-chi/render"
	"net/http"
)

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

func ErrNotFound(err error) render.Renderer {
	return &errResponse{
		Error: err.Error(),
		Code:  404,
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
