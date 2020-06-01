package services

import (
	"github.com/stetsd/micro-brick/internal/domain/repositoryInterfaces"
)

type ServiceUser struct {
	UserStore repositoryInterfaces.UserStore
}

const ServiceUserName = "ServiceUser"

func (su *ServiceUser) Login() error {
	// TODO: DB include

	return nil
}

func (su *ServiceUser) Registration() error {
	// TODO: DB include
	//return errors.ErrorEmailExists

	return nil
}
