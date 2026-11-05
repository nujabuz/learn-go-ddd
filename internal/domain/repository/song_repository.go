package repository

import (
	"context"

	"github.com/nujabuz/learn-go-ddd/internal/domain/entities"
)

type Song interface {
	Create(c context.Context, song *entities.ValidatedSong) (*entities.Song, error)
	Update(c context.Context, song *entities.ValidatedSong) (*entities.Song, error)
	FindByID(c context.Context, id int64) (*entities.Song, error)
	Delete(c context.Context, id int64) error
}
