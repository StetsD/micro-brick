package services

import (
	"github.com/stetsd/micro-brick/internal/domain/repositoryInterfaces"
)

type ServiceUser struct {
	UserStore repositoryInterfaces.UserStore
}

const ServiceUserName = "ServiceUser"

func (su *ServiceUser) Login() error {
	// TODO: validation module, logging
	// TODO: DB include
	// TODO: id, email hardcode

	//user := &models.User{1, name, ""}

	//return &models.User{}
	return nil
}

func (su *ServiceUser) Registration() error {
	// TODO: validation module, logging
	// TODO: DB include
	// TODO: id, email hardcode

	//user := &models.User{1, name, email}

	//return &models.User{}
	return nil
}
