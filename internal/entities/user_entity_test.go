package entities_test

import (
	"testing"

	"github.com/IcaroSilvaFK/goexpert_first_api/internal/entities"
	"github.com/stretchr/testify/assert"
)

func TestShouldCreateAnUser(t *testing.T) {

	user, err := entities.NewUserEntity("jhonDoe@example.com", "Jhon Doe", "123456")

	assert.Nil(t, err)
	assert.NotEmpty(t, user)
	assert.NotEmpty(t, user.ID)
	assert.Equal(t, "jhonDoe@example.com", user.Email)
	assert.NotEmpty(t, user.Password)
	assert.Equal(t, "Jhon Doe", user.Name)
}

func TestUser_Validate_Password(t *testing.T) {

	user, err := entities.NewUserEntity("jhonDoe@example.com", "Jhon Doe", "123456")

	assert.Nil(t, err)
	assert.Nil(t, user.ValidatePassword("123456"))
	assert.Error(t, user.ValidatePassword("test"))
	assert.NotEqual(t, "123456", user.Password)
}
