package http

import (
	"api/adapters/dtos"
	"api/core/domain"
	"api/core/errors"
	"api/mocks"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUserHandler_CreateUser(t *testing.T) {
	t.Run("should return 201 when user is created", func(t *testing.T) {
		userService := new(mocks.UserService)
		userHandler := NewUserHandler(userService)

		userDto := dtos.User{
			Name:    "any_name",
			Email:   "any_email",
			Phone:   "any_phone",
			Address: "any_address",
		}

		userService.On("CreateUser", mock.Anything, mock.Anything).Return("any_id", nil)

		body, _ := json.Marshal(userDto)

		req := httptest.NewRequest("POST", "/users", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		rr := httptest.NewRecorder()

		userHandler.CreateUser(rr, req)

		assert.Equal(t, http.StatusCreated, rr.Code)
	})

	t.Run("should return 400 when body is invalid", func(t *testing.T) {
		userService := new(mocks.UserService)
		userHandler := NewUserHandler(userService)

		req := httptest.NewRequest("POST", "/users", bytes.NewReader([]byte("invalid")))
		req.Header.Set("Content-Type", "application/json")
		rr := httptest.NewRecorder()

		userHandler.CreateUser(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
	})

	t.Run("should return 409 when user already exists", func(t *testing.T) {
		userService := new(mocks.UserService)
		userHandler := NewUserHandler(userService)

		userDto := dtos.User{
			Name:    "any_name",
			Email:   "any_email",
			Phone:   "any_phone",
			Address: "any_address",
		}

		userService.On("CreateUser", mock.Anything, mock.Anything).Return("", &errors.AlreadyExists{Code: http.StatusConflict, Message: "User already exists"})

		body, _ := json.Marshal(userDto)

		req := httptest.NewRequest("POST", "/users", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		rr := httptest.NewRecorder()

		userHandler.CreateUser(rr, req)

		assert.Equal(t, http.StatusConflict, rr.Code)
	})

	t.Run("should return 500 when something goes wrong", func(t *testing.T) {
		userService := new(mocks.UserService)
		userHandler := NewUserHandler(userService)

		userDto := dtos.User{
			Name:    "any_name",
			Email:   "any_email",
			Phone:   "any_phone",
			Address: "any_address",
		}

		userService.On("CreateUser", mock.Anything, mock.Anything).Return("", assert.AnError)

		body, _ := json.Marshal(userDto)

		req := httptest.NewRequest("POST", "/users", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		rr := httptest.NewRecorder()

		userHandler.CreateUser(rr, req)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
	})
}

func TestUserHandler_GetUser(t *testing.T) {
	t.Run("should return 200 when user is found", func(t *testing.T) {
		userService := new(mocks.UserService)
		userHandler := NewUserHandler(userService)

		user := domain.User{
			ID:      "any_id",
			Name:    "any_name",
			Email:   "any_email",
			Phone:   "any_phone",
			Address: "any_address",
		}

		userService.On("GetUser", mock.Anything, "any_email").Return(user, nil)

		req := httptest.NewRequest("GET", "/users?email=any_email", nil)
		rr := httptest.NewRecorder()

		userHandler.GetUser(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
	})

	t.Run("should return 404 when user is not found", func(t *testing.T) {
		userService := new(mocks.UserService)
		userHandler := NewUserHandler(userService)

		userService.On("GetUser", mock.Anything, "any_email").Return(domain.User{}, assert.AnError)

		req := httptest.NewRequest("GET", "/users?email=any_email", nil)
		rr := httptest.NewRecorder()

		userHandler.GetUser(rr, req)

		assert.Equal(t, http.StatusNotFound, rr.Code)
	})
}

func TestUserHandler_DeleteUser(t *testing.T) {
	r := chi.NewRouter()
	userService := new(mocks.UserService)
	userHandler := NewUserHandler(userService)
	r.Delete("/users", userHandler.DeleteUser)

	t.Run("should return 200 when user is deleted", func(t *testing.T) {
		userService.On("DeleteUser", mock.Anything, "any_email").Return("any_email", nil).Once()

		req := httptest.NewRequest("DELETE", "/users?email=any_email", nil)
		rr := httptest.NewRecorder()

		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
	})

	t.Run("should return 500 when something goes wrong", func(t *testing.T) {
		userService.On("DeleteUser", mock.Anything, "any_email").Return("", assert.AnError).Once()

		req := httptest.NewRequest("DELETE", "/users?email=any_email", nil)
		rr := httptest.NewRecorder()

		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
	})
}

func TestUserHandler_Search(t *testing.T) {
	r := chi.NewRouter()
	userService := new(mocks.UserService)
	userHandler := NewUserHandler(userService)
	r.Get("/users/search", userHandler.Search)

	t.Run("should return 200 when users are found", func(t *testing.T) {
		users := []domain.User{
			{
				ID:      "any_id",
				Name:    "any_name",
				Email:   "any_email",
				Phone:   "any_phone",
				Address: "any_address",
			},
		}

		userService.On("Search", mock.Anything, "any_email").Return(users, nil).Once()

		req := httptest.NewRequest("GET", "/users/search?email=any_email", nil)
		rr := httptest.NewRecorder()

		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
	})

	t.Run("should return 404 when no users are found", func(t *testing.T) {
		userService.On("Search", mock.Anything, "any_email").Return([]domain.User{}, assert.AnError).Once()

		req := httptest.NewRequest("GET", "/users/search?email=any_email", nil)
		rr := httptest.NewRecorder()

		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusNotFound, rr.Code)
	})
}
