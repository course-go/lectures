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
	app := app.New()
	window := app.NewWindow("My cool app")

	clickBinding := binding.NewInt()
	nameBinding := binding.NewString()
	nameInput := widget.NewEntryWithData(nameBinding)
	submitButton := widget.NewButtonWithIcon("Submit", theme.ConfirmIcon(), func() {
		current, _ := clickBinding.Get()
		_ = clickBinding.Set(current + 1)
	})
	clicks := widget.NewLabelWithData(binding.IntToString(clickBinding))
	nameContainer := container.New(layout.NewHBoxLayout(), nameInput, submitButton, clicks)
	check := widget.NewCheck("Go is the best!", func(value bool) {
		log.Println("This must be always checked!")
	})
	radio := widget.NewRadioGroup([]string{"Gophers", "Even more Gophers"}, func(value string) {
		log.Println("Radio set to", value)
	})
	choicesContainer := container.New(layout.NewHBoxLayout(), check, radio)

	content := container.New(layout.NewVBoxLayout(), nameContainer, choicesContainer)
	window.SetContent(content)
	window.ShowAndRun()
}

// END OMIT
