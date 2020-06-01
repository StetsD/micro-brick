package tools

import (
	"github.com/stetsd/micro-brick/internal/domain/repositories"
	"github.com/stetsd/micro-brick/internal/domain/services"
	"github.com/stetsd/micro-brick/internal/infrastructure/dbDriver"
	"log"
)

type ServiceCollection struct {
	ServiceUser services.ServiceUser
}

func Bind(driver *dbDriver.DbDriver, serviceNames ...string) ServiceCollection {
	serviceCollection := ServiceCollection{}
	for _, service := range serviceNames {
		switch service {
		case services.ServiceUserName:
			pgRepoUserStore, err := repositories.NewPgRepoUserStore(driver)
			if err != nil {
				log.Fatalf("%v", err)
			}

			serviceCollection.ServiceUser = services.ServiceUser{
				UserStore: pgRepoUserStore,
			}
		}
	}

	return serviceCollection
}
