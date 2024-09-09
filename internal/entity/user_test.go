package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("John", "test@email.com", "test")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.Equal(t, "John", user.Name)
	assert.Equal(t, "test@email.com", user.Email)
	assert.NotEmpty(t, user.Password)
}

func TestValidatePassword(t *testing.T) {
	user, _ := NewUser("John", "test@email.com", "test")
	assert.NotNil(t, user.Password)
	assert.True(t, user.ValidatePassword("test"))
	assert.False(t, user.ValidatePassword("wrong"))
	assert.NotEqual(t, "test", user.Password)
}
