package db

import (
	"testing"
	"time"

	"github.com/go-redis/redismock/v9"
	"github.com/stretchr/testify/assert"
)

func TestLocalCache_Set(t *testing.T) {
	db, mock := redismock.NewClientMock()

	cache := &LocalCache{
		cache: db,
	}

	t.Run("should set value in cache", func(t *testing.T) {
		mock.ExpectSet("any_key", "any_value", 10*time.Second).SetVal("")

		err := cache.Set("any_key", "any_value", 10)

		assert.Nil(t, err)
	})

	t.Run("should return error when something goes wrong", func(t *testing.T) {
		mock.ExpectSet("any_key", "any_value", 10*time.Second).SetErr(assert.AnError)

		err := cache.Set("any_key", "any_value", 10)

		assert.NotNil(t, err)
	})
}

func TestLocalCache_Get(t *testing.T) {
	db, mock := redismock.NewClientMock()

	cache := &LocalCache{
		cache: db,
	}

	t.Run("should get value from cache", func(t *testing.T) {
		mock.ExpectGet("any_key").SetVal("any_value")

		value, err := cache.Get("any_key")

		assert.Nil(t, err)
		assert.Equal(t, "any_value", value)
	})

	t.Run("should return error when something goes wrong", func(t *testing.T) {
		mock.ExpectGet("any_key").SetErr(assert.AnError)

		value, err := cache.Get("any_key")

		assert.NotNil(t, err)
		assert.Empty(t, value)
	})
}

func TestLocalCache_Delete(t *testing.T) {
	db, mock := redismock.NewClientMock()

	cache := &LocalCache{
		cache: db,
	}

	t.Run("should delete value from cache", func(t *testing.T) {
		mock.ExpectDel("any_key").SetVal(1)

		err := cache.Delete("any_key")

		assert.Nil(t, err)
	})

	t.Run("should return error when something goes wrong", func(t *testing.T) {
		mock.ExpectDel("any_key").SetErr(assert.AnError)

		err := cache.Delete("any_key")

		assert.NotNil(t, err)
	})
}
