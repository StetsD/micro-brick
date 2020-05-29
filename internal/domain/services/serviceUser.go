package services

import (
	"context"
	"github.com/stetsd/micro-brick/internal/domain/models"
	"github.com/stetsd/micro-brick/internal/domain/repositoryInterfaces"
)

type ServiceUser struct {
	UserStore repositoryInterfaces.UserStore
}

func (su *ServiceUser) Login(ctx context.Context, name string, password string) *models.User {
	// TODO: validation module, logging
	// TODO: DB include
	// TODO: id, email hardcode

	user := &models.User{1, name, ""}

	return user
}

func (su *ServiceUser) Registration(ctx context.Context, name, email, password string) *models.User {
	// TODO: validation module, logging
	// TODO: DB include
	// TODO: id, email hardcode

	user := &models.User{1, name, email}

	return user
}
