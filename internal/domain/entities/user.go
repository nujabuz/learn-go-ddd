package entities

import (
	"fmt"
	"time"
)

var (
	// MaxCharUsername for maximum character username.
	MaxCharUsername = 16

	// MinCharUsername for minimum character username.
	MinCharUsername = 4

	// MaxCharPassword for maximum character password.
	MaxCharPassword = 32

	// MinCharPassword for minimum character password.
	MinCharPassword = 8
)

// User entity.
type User struct {
	ID        int64 // autoincrement in database
	Username  string
	Password  string
	UpdatedAt time.Time
	CreatedAt time.Time
}

// NewUser to create user data based on the User object.
func NewUser(username string, password string) *User {
	return &User{
		Username:  username,
		Password:  password,
		UpdatedAt: time.Now(),
		CreatedAt: time.Now(),
	}
}

// UpdatedUsername to update the username in the User object.
func (u *User) UpdatedUsername(username string) error {
	u.Username = username
	u.UpdatedAt = time.Now()

	return u.validate()
}

// UpdatePassword to update the passwordin the User object.
func (u *User) UpdatePassword(password string) error {
	u.Password = password
	u.UpdatedAt = time.Now()

	return u.validate()
}

// validate to validate the value of the User object.
func (u *User) validate() error {
	// validate username
	if u.Username == "" {
		return fmt.Errorf("%w: username must not be empty", ErrValidation)
	}

	if len(u.Username) > MaxCharUsername {
		return fmt.Errorf("%w: username must not exceed %d characters", ErrValidation, MaxCharUsername)
	}

	if len(u.Username) < MinCharUsername {
		return fmt.Errorf("%w: username must be at least %d characters", ErrValidation, MinCharUsername)
	}

	if u.Password == "" {
		return fmt.Errorf("%w: password must not be empty", ErrValidation)
	}

	if len(u.Password) > MaxCharPassword {
		return fmt.Errorf("%w: password must not exceed %d characters", ErrValidation, MaxCharPassword)
	}

	if len(u.Password) < MinCharPassword {
		return fmt.Errorf("%w: password must be at least %d characters", ErrValidation, MinCharPassword)
	}

	return nil
}
