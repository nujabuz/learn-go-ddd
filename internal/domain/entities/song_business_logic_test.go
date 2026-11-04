package entities_test

import (
	"testing"
	"time"

	"github.com/nujabuz/learn-go-ddd/internal/domain/entities"
	"github.com/stretchr/testify/assert"
)

func TestNewSong_UpdateTitle(t *testing.T) {
	// test success update title
	song := entities.NewSong("Lady Brown")

	time.Sleep(1 * time.Millisecond)

	err := song.UpdateTitle("Luv Sic")
	assert.Nil(t, err)
	assert.Equal(t, "Luv Sic", song.Title)
	assert.True(t, song.UpdatedAt.After(song.CreatedAt))
}

func TestNewSong_UpdateTitle_EmptyTitle(t *testing.T) {
	// test success update title
	song := entities.NewSong("Lady Brown")

	time.Sleep(1 * time.Millisecond)

	err := song.UpdateTitle("")
	assert.NotNil(t, err)
	assert.Equal(t, "", song.Title)
	assert.True(t, song.UpdatedAt.After(song.CreatedAt))
	assert.Contains(t, err.Error(), "title must not be empty")
}

func TestNewSong_UpdateTitle_ExceedMaximum(t *testing.T) {
	// test success update title
	song := entities.NewSong("Lady Brown")

	time.Sleep(1 * time.Millisecond)

	err := song.UpdateTitle("aabcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789abcdefghijklmnopqrstuvwxyz01")
	assert.NotNil(t, err)
	assert.Equal(t, "aabcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789abcdefghijklmnopqrstuvwxyz01", song.Title)
	assert.True(t, song.UpdatedAt.After(song.CreatedAt))
	assert.Contains(t, err.Error(), "title must not be exceed")
}

func TestNewSong_UpdateTitle_LessThanMinimum(t *testing.T) {
	// test success update title
	song := entities.NewSong("Lady Brown")

	time.Sleep(1 * time.Millisecond)

	err := song.UpdateTitle("n")
	assert.NotNil(t, err)
	assert.Equal(t, "n", song.Title)
	assert.True(t, song.UpdatedAt.After(song.CreatedAt))
	assert.Contains(t, err.Error(), "title must be at least")
}
