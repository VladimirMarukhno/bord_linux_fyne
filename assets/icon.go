package assets

import (
	"log"

	"fyne.io/fyne/v2"
)

func GetIcon(s string) fyne.Resource {
	icon, err := fyne.LoadResourceFromPath("assets/pictures/" + s)
	if err != nil {
		log.Println("ошибка добовления иконки"+s, err)
	}

	return icon
}
