package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserValidation(t *testing.T) {
	// test valid user
	validUser := &User{
		Username: "nujabuz",
		Password: "nujabuz123",
	}
	err := validUser.validate()
	assert.Nil(t, err)

	// test invalid user (empty username)
	invalidUser := &User{
		Username: "",
		Password: "nujabuz123",
	}
	err = invalidUser.validate()
	assert.Error(t, err)
}

func TestNewValidatedUser(t *testing.T) {
	// test valid user to be validated
	user := &User{
		Username: "nujabuz",
		Password: "nujabuz123",
	}
	validatedUser, err := NewValidatedUser(user)
	assert.NotNil(t, validatedUser)
	assert.Nil(t, err)
	assert.True(t, validatedUser.IsValid())

	// test invalid user to be validated (empty password)
	user = &User{
		Username: "nujabuz",
		Password: "",
	}
	invalidUser, err := NewValidatedUser(user)
	assert.NotNil(t, err)
	assert.Nil(t, invalidUser)
}
