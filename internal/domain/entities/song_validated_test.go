package entities_test

import (
	"testing"

	"github.com/nujabuz/learn-go-ddd/internal/domain/entities"
	"github.com/stretchr/testify/assert"
)

func TestNewValidationSong(t *testing.T) {
	// test valid song to be validated song
	song := &entities.Song{Title: "Lady Brown"}
	validSong, err := entities.NewValidatedSong(song)
	assert.Nil(t, err)
	assert.NotNil(t, validSong)
	assert.Equal(t, "Lady Brown", validSong.Title)

	// test invalid song to be validated song
	song = &entities.Song{Title: ""}
	invalidSong, err := entities.NewValidatedSong(song)
	assert.Nil(t, invalidSong)
	assert.NotNil(t, err)
}
