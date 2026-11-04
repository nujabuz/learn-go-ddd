package entities

import (
	"fmt"
	"time"
)

var (
	// MaxCharSongTitle for maximum charater song title.
	MaxCharSongTitle = 100
	// MinCharSongTitle for minimum charater song title.
	MinCharSongTitle = 2
)

// Song type object.
type Song struct {
	ID        int64 // autoincrement in database.
	Title     string
	UpdatedAt time.Time
	CreatedAt time.Time
}

// NewSong to make new Song object.
func NewSong(title string) *Song {
	return &Song{
		Title:     title,
		UpdatedAt: time.Now(),
		CreatedAt: time.Now(),
	}
}

// UpdateTitle to updated title value in object Song.
func (s *Song) UpdateTitle(title string) error {
	s.Title = title
	s.UpdatedAt = time.Now()

	return s.validate()
}

// validate to validate value object Song.
func (s *Song) validate() error {
	if s.Title == "" {
		return fmt.Errorf("%w: title must not be empty", ErrValidation)
	}

	if len(s.Title) > MaxCharSongTitle {
		return fmt.Errorf("%w: title must not be exceed %d characters", ErrValidation, MaxCharSongTitle)
	}

	if len(s.Title) < MinCharSongTitle {
		return fmt.Errorf("%w: title must be at least %d characters", ErrValidation, MinCharSongTitle)
	}

	return nil
}
