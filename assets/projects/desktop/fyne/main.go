package main

import (
	"log"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// START OMIT

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("My giga cool app")

	var clicks uint
	nameBinding := binding.NewString()
	name := widget.NewEntryWithData(nameBinding)
	button := widget.NewButtonWithIcon("Click", theme.CheckButtonIcon(), func() {
		clicks++
	})
	nameContainer := container.New(layout.NewHBoxLayout(), name, button)

	check := widget.NewCheck("Go is the best!", func(value bool) {
		log.Println("This must be always checked!")
	})
	radio := widget.NewRadioGroup([]string{"Gophers", "Even more Gophers"}, func(value string) {
		log.Println("Radio set to", value)
	})
	choicesContainer := container.New(layout.NewHBoxLayout(), check, radio)

	content := container.New(layout.NewVBoxLayout(), nameContainer, choicesContainer)
	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}

// END OMIT
