package infrastructure

import (
	"github.com/Jamshid-Ismoilov/infrastructure/localize"
)

func Localize(service, word string) localize.Getter {
	return localize.NewLocalizer(service, word)
}
