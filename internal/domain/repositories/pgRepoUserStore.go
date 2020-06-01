package repositories

import (
	"context"
	"fmt"
	"github.com/stetsd/micro-brick/internal/app/schemas"
	"github.com/stetsd/micro-brick/internal/domain/models"
	"github.com/stetsd/micro-brick/internal/infrastructure/dbDriver"
	"github.com/stetsd/micro-brick/internal/infrastructure/logger"
)

type PgRepoUserStore struct {
	pgd *dbDriver.DbDriver
}

func NewPgRepoUserStore(pgDriver *dbDriver.DbDriver) (*PgRepoUserStore, error) {
	return &PgRepoUserStore{pgDriver}, nil
}

func (pgRUS *PgRepoUserStore) Login(_ context.Context, _ *models.User) error {
	fmt.Println("plug")
	return nil
}

func (pgRUS *PgRepoUserStore) Registration(data *schemas.RegistrationBody) error {
	// TODO: password not salted
	rows, err := pgRUS.pgd.Db.Query(`
		INSERT INTO "User" (name, email, password) 
		VALUES ($1, $2, $3)`,
		data.Name,
		data.Email,
		data.Password,
	)

	if err != nil {
		return err
	}

	defer func() {
		if err := rows.Close(); err != nil {
			logger.Log.Fatal(err.Error())
		}
	}()

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
