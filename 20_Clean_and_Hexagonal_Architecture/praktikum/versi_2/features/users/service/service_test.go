package service_test

import (
	"clean_architecture/features/users/dto"
	"clean_architecture/features/users/mock"
	"clean_architecture/features/users/service"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	var repo = mock.Repository{Status: "invalid"}
	var serv = service.New(&repo)

	t.Run("Mapping error", func(t *testing.T) {
		var user = dto.UserInput{}
		result, err := serv.Create(user)

		assert.Empty(t, result)
		assert.NotNil(t, err)
	})

	t.Run("Process error", func(t *testing.T) {
		var user = dto.UserInput{
			Email:    "bagus@gmail.com",
			Password: "bagus123",
		}

		result, err := serv.Create(user)

		assert.Empty(t, result)
		assert.NotNil(t, err)
	})

	t.Run("Success", func(t *testing.T) {
		var repo = mock.Repository{Status: "valid"}
		var serv = service.New(&repo)

		var user = dto.UserInput{
			Email:    "bagus@gmail.com",
			Password: "bagus123",
		}

		result, err := serv.Create(user)

		assert.NotEmpty(t, result)
		assert.Nil(t, err)
		assert.Equal(t, "bagus@gmail.com", result.Email)
		assert.Equal(t, "bagus123", result.Password)
	})
}

func TestGetAllUsers(t *testing.T) {
	var repo = mock.Repository{Status: "valid"}
	var serv = service.New(&repo)

	result, err := serv.GetAll()

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "bagus@gmail.com", result[0].Email)
	assert.Equal(t, "rio@gmail.com", result[1].Email)
}

func TestLogin(t *testing.T) {
	var repo = mock.Repository{Status: "valid"}
	var serv = service.New(&repo)

	result := serv.Login("bagus@gmail.com", "bagus123")

	assert.NotNil(t, result)
	assert.Equal(t, "bagus@gmail.com", result.Email)
	assert.Equal(t, "bagus123", result.Password)
}
