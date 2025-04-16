package main

import (
	"fmt"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello World")
	w.SetFullScreen(true)

	buttons := []*widget.Button{
		widget.NewButton("Youtube", func() {
			log.Println("click Youtube")
		}),
		widget.NewButton("Plex", func() {
			log.Println("click Plex")
		}),
		widget.NewButton("Plex", func() {
			log.Println("click Plex")
		}),
		widget.NewButton("Yandex music", func() {
			log.Println("click Yandex music")
		}),
		widget.NewButton("Xeoma", func() {
			log.Println("click Xeoma")
		}),
		widget.NewButton("Firefox", func() {
			log.Println("click firefox")
		}),
	}

	grid := container.New(layout.NewAdaptiveGridLayout(3), buttons[0], buttons[1], buttons[2], buttons[3], buttons[4], buttons[5])
	currentFocus := 0
	buttons[currentFocus].FocusGained()

	w.Canvas().SetOnTypedKey(func(ke *fyne.KeyEvent) {
		switch ke.Name {
		case fyne.KeyRight:
			currentFocus = (currentFocus + 1) % len(buttons)
			if currentFocus == 0 {
				buttons[len(buttons)-1].FocusLost()
				buttons[currentFocus].FocusGained()
			} else {
				buttons[((currentFocus - 1) % len(buttons))].FocusLost()
				buttons[currentFocus].FocusGained()
			}
		case fyne.KeyLeft:
			currentFocus = (currentFocus - 1 + len(buttons)) % len(buttons)
			buttons[((currentFocus + 1) % len(buttons))].FocusLost()
			buttons[currentFocus].FocusGained()
		case fyne.KeyUp:
			currentFocus = (currentFocus - 3) % len(buttons)
			if currentFocus < 0 {
				buttons[((currentFocus + 3) % len(buttons))].FocusLost()
				currentFocus = len(buttons) + currentFocus
			}
			buttons[((currentFocus + 3) % len(buttons))].FocusLost()
			buttons[currentFocus].FocusGained()
		case fyne.KeyDown:
			tmpCur := currentFocus
			currentFocus = (currentFocus + 3) % len(buttons)
			fmt.Println(currentFocus)
			if currentFocus <= 2 {
				fmt.Println("l: ", (len(buttons)-currentFocus)%len(buttons))
				fmt.Println("but: ", len(buttons), "curr: ", currentFocus)
				buttons[(tmpCur % len(buttons))].FocusLost()
				buttons[currentFocus].FocusGained()
			} else {
				buttons[((currentFocus - 3) % len(buttons))].FocusLost()
				buttons[currentFocus].FocusGained()
			}
		case fyne.KeySpace, fyne.KeyReturn:
			buttons[currentFocus].OnTapped()
		}
	})

	//vBox := container.New(layout.NewVBoxLayout(), grid)
	w.SetContent(grid)
	w.ShowAndRun()

}
