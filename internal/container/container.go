package container

import (
	"github.com/brunobandev/reppyapi/infrastructure/repository"
	"github.com/brunobandev/reppyapi/internal/service"
)

type Container struct {
	Services     Services
	Repositories Repositories
}

type Services struct {
	AccommodationService service.AccommodationService
	BillService          service.BillService
	PropertyService      service.PropertyService
	UserService          service.UserService
}

type Repositories struct {
	PropertyRepository repository.PropertyRepository
}

func Inject() Container {
	as := service.NewAccommodationService()
	bs := service.NewBillService()
	ps := service.NewPropertyService()
	us := service.NewUserService()

	services := Services{
		AccommodationService: as,
		BillService:          bs,
		PropertyService:      ps,
		UserService:          us,
	}

	pr := repository.NewPropertyRepository()

	repositories := Repositories{
		PropertyRepository: pr,
	}

	return Container{
		Services:     services,
		Repositories: repositories,
	}
}
