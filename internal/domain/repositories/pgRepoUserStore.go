package repositories

import (
	"context"
	"fmt"
	"github.com/stetsd/micro-brick/internal/domain/models"
)

// TODO: zapilit pg driver

type PgRepoUserStore struct {
	// plug
}

func NewPgRepoUserStore() (*PgRepoUserStore, error) {
	return &PgRepoUserStore{}, nil
}

func (pgRUS *PgRepoUserStore) Login(ctx context.Context, user *models.User) error {
	fmt.Println("plug")
	return nil
}

func (pgRUS *PgRepoUserStore) Registration(ctx context.Context, user *models.User) error {
	fmt.Println("plug")
	return nil
}

func (pgRUS *PgRepoUserStore) Logout(ctx context.Context, id int) error {
	fmt.Println("plug")
	return nil
}

func (pgRUS *PgRepoUserStore) Put(ctx context.Context, user *models.User) (*models.User, error) {
	fmt.Println("plug")
	return &models.User{}, nil
}

func (pgRUS *PgRepoUserStore) Delete(ctx context.Context, id int) error {
	fmt.Println("plug")
	return nil
}

func (pgRUS *PgRepoUserStore) GetById(ctx context.Context, id int) (*models.User, error) {
	fmt.Println("plug")
	return &models.User{}, nil
}
