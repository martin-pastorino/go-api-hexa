package main

import (
	"api/cmd/app"
	"api/infra/router"
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	localMiddleware "api/infra/middleware"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

const (
	defaultPort     = 8080
	shutdownTimeout = 10 * time.Second
)

func main() {
	port := getPort()
	server := setupServer()

	// Start server in a goroutine
	go func() {
		fmt.Printf("Server running on port %d\n", port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("Server error: %v\n", err)
			os.Exit(1)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		fmt.Printf("Server forced to shutdown: %v\n", err)
	}
	fmt.Println("Server gracefully stopped")
}

func setupServer() *http.Server {
	r := chi.NewRouter()

	// Middleware stack
	r.Use(localMiddleware.Recovery)
	r.Use(localMiddleware.RequiredHeaders)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	// Health check endpoint
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		render.JSON(w, r, map[string]string{"status": "ok"})
	})

	// API routes
	r.Mount("/users", router.UsersAPIRouter(app.InitializeUsersHandler()))
	r.Mount("/products", router.ProductAPIRouter(app.InitializeProductsHandler()))

	return &http.Server{
		Addr:         fmt.Sprintf(":%d", getPort()),
		Handler:      r,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
}

func getPort() int {
	portStr := os.Getenv("PORT")
	if portStr == "" {
		return defaultPort
	}
	port, err := strconv.Atoi(portStr)
	if err != nil || port < 1 || port > 65535 {
		return defaultPort
	}
	return port
}
