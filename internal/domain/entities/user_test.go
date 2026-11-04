package entities_test

import (
	"testing"

	"github.com/nujabuz/learn-go-ddd/internal/domain/entities"
)

func TestNewUser_Success(t *testing.T) {
	user := entities.NewUser("nujabuz", "nujabuz123")

	if user.Username != "nujabuz" {
		t.Errorf("expected username to be 'nujabuz', but got '%s'", user.Username)
	}

	if user.Password != "nujabuz123" {
		t.Errorf("expected password to be 'nujabuz123', but got '%s'", user.Password)
	}
}
