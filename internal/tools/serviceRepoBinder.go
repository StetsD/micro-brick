package tools

import (
	"github.com/stetsd/micro-brick/internal/domain/repositories"
	"github.com/stetsd/micro-brick/internal/domain/services"
	"log"
)

type ServiceCollection struct {
	UserService *services.ServiceUser
}

func Bind(serviceNames ...string) ServiceCollection {
	serviceCollection := ServiceCollection{}
	for _, service := range serviceNames {
		switch service {
		case "user":
			pgRepoUserStore, err := repositories.NewPgRepoUserStore()
			if err != nil {
				log.Fatalf("%v", err)
			}

			serviceCollection.UserService = &services.ServiceUser{
				UserStore: pgRepoUserStore,
			}
		}
	}

	return serviceCollection
}
