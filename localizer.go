package infrastructure

import (
	"github.com/Jamshid-Ismoilov/infrastructure/localize"
)

func Localize(service string) (server localize.ServiceGetter) {
	return localize.NewLocalizer(service)
}
