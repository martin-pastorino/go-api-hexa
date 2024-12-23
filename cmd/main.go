package main

import (
	"api/internal/models"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

var port int16
func main() {
	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world!"))
	})

	r.Get("/person", func(w http.ResponseWriter, r *http.Request) {
		p :=  models.Person{Name: "Martin", Age: 30}
		render.Render(w, r,p)
	})

	http.ListenAndServe( fmt.Sprintf(":%d",port), r)
}
