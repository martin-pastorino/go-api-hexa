package models

import (
	"net/http"

	"github.com/go-chi/render"
)

type Person struct{
	Name string
	Age int
}

func (p Person) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, http.StatusOK)
	return nil
}