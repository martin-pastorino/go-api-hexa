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
	userHandler := app.InitializeUsersHandler()
	productHandler := app.InitializeProductsHandler()

	// API routes
	r := router.New(userHandler, productHandler)



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
