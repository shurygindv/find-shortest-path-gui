package main

import (
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	app := app.New()
	window := app.NewWindow("Main window")

	os.Setenv("FYNE_SCALE", "1.5")

	content := MainAppContent() // main-app-content.go

	window.CenterOnScreen()
	window.Resize(fyne.NewSize(500, 350))
	window.SetTitle("Path finder")
	window.SetContent(content)
	window.ShowAndRun()
}
