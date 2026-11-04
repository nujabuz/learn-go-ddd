package entities_test

import (
	"testing"

	"github.com/nujabuz/learn-go-ddd/internal/domain/entities"
	"github.com/stretchr/testify/assert"
)

func TestNewSong(t *testing.T) {
	song := entities.NewSong("Lady Brown")
	assert.Equal(t, "Lady Brown", song.Title)
}
