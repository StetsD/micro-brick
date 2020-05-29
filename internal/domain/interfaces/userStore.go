package interfaces

import (
	"context"
	"github.com/stetsd/micro-brick/internal/domain/models"
)

type UserStore interface {
	Login(ctx context.Context, user *models.User) error
	Registration(ctx context.Context, user *models.User) error
	Logout(ctx context.Context, id int) error
	Put(ctx context.Context, user *models.User) (*models.User, error)
	Delete(ctx context.Context, id int) error
	GetById(ctx context.Context, id int) (*models.User, error)
}
