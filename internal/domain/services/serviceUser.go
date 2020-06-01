package services

import (
	"github.com/stetsd/micro-brick/internal/app/schemas"
	"github.com/stetsd/micro-brick/internal/domain/repositoryInterfaces"
)

type ServiceUser struct {
	UserStore repositoryInterfaces.UserStore
}

const ServiceUserName = "ServiceUser"

func (su *ServiceUser) Login() error {

	return nil
}

func (su *ServiceUser) Registration(data *schemas.RegistrationBody) error {

	err := su.UserStore.Registration(data)

	if err != nil {
		return err
	}

	//return errors.ErrorEmailExists

	return nil
}
