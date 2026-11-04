package entities_test

import (
	"testing"
	"time"

	"github.com/nujabuz/learn-go-ddd/internal/domain/entities"
	"github.com/stretchr/testify/assert"
)

func TestUser_UpdateUsername(t *testing.T) {
	user := entities.NewUser("nujabuz", "nujabuz123")

	// Give some time to ensure UpdateAt changes
	time.Sleep(1 * time.Millisecond)

	err := user.UpdatedUsername("nujabes")
	assert.Nil(t, err)
	assert.Equal(t, "nujabes", user.Username)
	assert.True(t, user.UpdatedAt.After(user.CreatedAt))
}

func TestUser_UpdateUsername_EmptyUsername(t *testing.T) {
	user := entities.NewUser("nujabuz", "nujabuz123")

	// Give some time to ensure UpdateAt changes
	time.Sleep(1 * time.Millisecond)

	err := user.UpdatedUsername("")
	assert.Error(t, err)
	assert.Equal(t, "", user.Username)
	assert.True(t, user.UpdatedAt.After(user.CreatedAt))
	assert.Contains(t, err.Error(), "username must not be empty")
}

func TestUser_UpdateUsername_ExceedMaximum(t *testing.T) {
	user := entities.NewUser("nujabuz", "nujabuz123")

	// Give some time to ensure UpdateAt changes
	time.Sleep(1 * time.Millisecond)

	err := user.UpdatedUsername("nujabuz.nujabuz.n")
	assert.Error(t, err)
	assert.Equal(t, "nujabuz.nujabuz.n", user.Username)
	assert.True(t, user.UpdatedAt.After(user.CreatedAt))
	assert.Contains(t, err.Error(), "username must not exceed")
}

func TestUser_UpdateUsername_LessThanMinimum(t *testing.T) {
	user := entities.NewUser("nujabuz", "nujabuz123")

	// Give some time to ensure UpdateAt changes
	time.Sleep(1 * time.Millisecond)

	err := user.UpdatedUsername("nuj")
	assert.Error(t, err)
	assert.Equal(t, "nuj", user.Username)
	assert.True(t, user.UpdatedAt.After(user.CreatedAt))
	assert.Contains(t, err.Error(), "username must be at least")
}

func TestUser_UpdatePassword(t *testing.T) {
	user := entities.NewUser("nujabuz", "nujabuz123")

	// Give some time to ensure UpdateAt changes
	time.Sleep(1 * time.Millisecond)

	err := user.UpdatePassword("nujabes321")
	assert.Nil(t, err)
	assert.Equal(t, "nujabes321", user.Password)
	assert.True(t, user.UpdatedAt.After(user.CreatedAt))
}

func TestUser_UpdatePassword_EmptyPassword(t *testing.T) {
	user := entities.NewUser("nujabuz", "nujabuz123")

	// Give some time to ensure UpdateAt changes
	time.Sleep(1 * time.Millisecond)

	err := user.UpdatePassword("")
	assert.Error(t, err)
	assert.Equal(t, "", user.Password)
	assert.True(t, user.UpdatedAt.After(user.CreatedAt))
	assert.Contains(t, err.Error(), "password must not be empty")
}

func TestUser_UpdatePassword_ExceedMaximum(t *testing.T) {
	user := entities.NewUser("nujabuz", "nujabuz123")

	// Give some time to ensure UpdateAt changes
	time.Sleep(1 * time.Millisecond)

	err := user.UpdatePassword("nujabuz.nujabuz.nujabuz.nujabuz.n")
	assert.Error(t, err)
	assert.Equal(t, "nujabuz.nujabuz.nujabuz.nujabuz.n", user.Password)
	assert.True(t, user.UpdatedAt.After(user.CreatedAt))
	assert.Contains(t, err.Error(), "password must not exceed")
}

func TestUser_UpdatePassword_LessThanMinimum(t *testing.T) {
	user := entities.NewUser("nujabuz", "nujabuz123")

	// Give some time to ensure UpdateAt changes
	time.Sleep(1 * time.Millisecond)

	err := user.UpdatePassword("nujabuz")
	assert.Error(t, err)
	assert.Equal(t, "nujabuz", user.Password)
	assert.True(t, user.UpdatedAt.After(user.CreatedAt))
	assert.Contains(t, err.Error(), "password must be at least")
}
