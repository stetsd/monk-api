package tools

import (
	"github.com/stetsd/monk-api/internal/domain/repositories"
	"github.com/stetsd/monk-api/internal/domain/services"
	"github.com/stetsd/monk-api/internal/infrastructure/dbDriver"
)

type ServiceCollection struct {
	ServiceUser services.ServiceUser
}

func Bind(driver *dbDriver.DbDriver, serviceNames ...string) ServiceCollection {
	serviceCollection := ServiceCollection{}
	for _, service := range serviceNames {
		switch service {
		case services.ServiceUserName:
			pgRepoUserStore := repositories.NewPgRepoUserStore(driver)
			serviceCollection.ServiceUser = services.ServiceUser{
				UserStore: pgRepoUserStore,
			}
		}
	}

	return serviceCollection
}
