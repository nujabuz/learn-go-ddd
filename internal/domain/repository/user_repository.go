// Package repository of database
package repository

import (
	"context"

	"github.com/nujabuz/learn-go-ddd/internal/domain/entities"
)

type UserRepository interface {
	Create(ctx context.Context, user *entities.User) (*entities.User, error)
	Update(ctx context.Context, user *entities.User) (*entities.User, error)
	FindByID(ctx context.Context, id int64) (*entities.User, error)
	FindByUsername(ctx context.Context, username string) (*entities.User, error)
	Delete(ctx context.Context, id int64) error
}
