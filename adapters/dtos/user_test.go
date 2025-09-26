package dtos

import (
	"api/core/domain"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUser_Bind(t *testing.T) {
	t.Run("should return error when address is empty", func(t *testing.T) {
		u := User{
			Address: "",
		}
		err := u.Bind(&http.Request{})
		assert.NotNil(t, err)
		assert.Equal(t, "address is required", err.Error())
	})

	t.Run("should return error when email is empty", func(t *testing.T) {
		u := User{
			Address: "any_address",
			Email:   "",
		}
		err := u.Bind(&http.Request{})
		assert.NotNil(t, err)
		assert.Equal(t, "email is required", err.Error())
	})

	t.Run("should return error when name is empty", func(t *testing.T) {
		u := User{
			Address: "any_address",
			Email:   "any_email",
			Name:    "",
		}
		err := u.Bind(&http.Request{})
		assert.NotNil(t, err)
		assert.Equal(t, "name is required", err.Error())
	})

	t.Run("should return error when phone is empty", func(t *testing.T) {
		u := User{
			Address: "any_address",
			Email:   "any_email",
			Name:    "any_name",
			Phone:   "",
		}
		err := u.Bind(&http.Request{})
		assert.NotNil(t, err)
		assert.Equal(t, "phone is required", err.Error())
	})

	t.Run("should return nil when all fields are valid", func(t *testing.T) {
		u := User{
			Address: "any_address",
			Email:   "any_email",
			Name:    "any_name",
			Phone:   "any_phone",
		}
		err := u.Bind(&http.Request{})
		assert.Nil(t, err)
	})
}

func TestUser_ToUserDomainModel(t *testing.T) {
	t.Run("should convert User to domain.User", func(t *testing.T) {
		u := User{
			ID:      "any_id",
			Name:    "any_name",
			Email:   "any_email",
			Phone:   "any_phone",
			Address: "any_address",
		}
		domainUser := u.ToUserDomainModel()
		assert.Equal(t, u.ID, domainUser.ID)
		assert.Equal(t, u.Name, domainUser.Name)
		assert.Equal(t, u.Email, domainUser.Email)
		assert.Equal(t, u.Phone, domainUser.Phone)
		assert.Equal(t, u.Address, domainUser.Address)
	})
}

func TestToUser(t *testing.T) {
	t.Run("should convert domain.User to User", func(t *testing.T) {
		domainUser := domain.User{
			ID:      "any_id",
			Name:    "any_name",
			Email:   "any_email",
			Phone:   "any_phone",
			Address: "any_address",
		}
		u := ToUser(domainUser)
		assert.Equal(t, domainUser.ID, u.ID)
		assert.Equal(t, domainUser.Name, u.Name)
		assert.Equal(t, domainUser.Email, u.Email)
		assert.Equal(t, domainUser.Phone, u.Phone)
		assert.Equal(t, domainUser.Address, u.Address)
	})
}

func TestToUsers(t *testing.T) {
	t.Run("should convert []domain.User to []User", func(t *testing.T) {
		domainUsers := []domain.User{
			{
				ID:      "any_id",
				Name:    "any_name",
				Email:   "any_email",
				Phone:   "any_phone",
				Address: "any_address",
			},
		}
		users := ToUsers(domainUsers)
		assert.Equal(t, len(domainUsers), len(users))
		assert.Equal(t, domainUsers[0].ID, users[0].ID)
		assert.Equal(t, domainUsers[0].Name, users[0].Name)
		assert.Equal(t, domainUsers[0].Email, users[0].Email)
		assert.Equal(t, domainUsers[0].Phone, users[0].Phone)
		assert.Equal(t, domainUsers[0].Address, users[0].Address)
	})
}
