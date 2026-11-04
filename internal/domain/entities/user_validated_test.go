package entities_test

import (
	"testing"

	"github.com/nujabuz/learn-go-ddd/internal/domain/entities"
	"github.com/stretchr/testify/assert"
)

func TestNewValidatedUser(t *testing.T) {
	// test valid user to be validated
	user := &entities.User{
		Username: "nujabuz",
		Password: "nujabuz123",
	}
	validatedUser, err := entities.NewValidatedUser(user)
	assert.NotNil(t, validatedUser)
	assert.Nil(t, err)
	assert.True(t, validatedUser.IsValid())

	// test invalid user to be validated (empty password)
	user = &entities.User{
		Username: "nujabuz",
		Password: "",
	}
	invalidUser, err := entities.NewValidatedUser(user)
	assert.NotNil(t, err)
	assert.Nil(t, invalidUser)
}
