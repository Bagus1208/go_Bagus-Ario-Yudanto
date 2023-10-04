package usecase

import (
	"clean_architecture/dto"
	"clean_architecture/repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	var mdl = repository.MockUserRepository{Status: "invalid"}
	var useCase = NewUserUseCase(&mdl)
	t.Run("Mapping error", func(t *testing.T) {
		var user = dto.UserDTO{}
		result, err := useCase.Create(user)

		assert.Empty(t, result)
		assert.NotNil(t, err)
	})

	t.Run("Process error", func(t *testing.T) {
		var user = dto.UserDTO{
			Email:    "bagus@gmail.com",
			Password: "bagus123",
		}

		result, err := useCase.Create(user)

		assert.Empty(t, result)
		assert.NotNil(t, err)
	})

	t.Run("Success", func(t *testing.T) {
		var mdl = repository.MockUserRepository{Status: "valid"}
		var useCase = NewUserUseCase(&mdl)

		var user = dto.UserDTO{
			Email:    "bagus@gmail.com",
			Password: "bagus123",
		}

		result, err := useCase.Create(user)

		assert.NotEmpty(t, result)
		assert.Nil(t, err)
		assert.Equal(t, "bagus@gmail.com", result.Email)
		assert.Equal(t, "bagus123", result.Password)
	})
}

func TestGetAllUsers(t *testing.T) {
	var mdl = repository.MockUserRepository{Status: "valid"}
	var useCase = NewUserUseCase(&mdl)

	result, err := useCase.All()

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "bagus@gmail.com", result[0].Email)
	assert.Equal(t, "rio@gmail.com", result[1].Email)
}

func TestLogin(t *testing.T) {
	var mdl = repository.MockUserRepository{Status: "valid"}
	var useCase = NewUserUseCase(&mdl)

	result := useCase.Login("bagus@gmail.com", "bagus123")

	assert.NotNil(t, result)
	assert.Equal(t, "bagus@gmail.com", result.Email)
	assert.Equal(t, "bagus123", result.Password)
}
