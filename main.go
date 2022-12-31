package main

import (
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	app := app.New()
	window := app.NewWindow("Main window")

	os.Setenv("FYNE_SCALE", "1.3")

	content := MainAppLayout() // main-app-layout.go

	window.CenterOnScreen()
	window.Resize(fyne.NewSize(550, 370))
	window.SetTitle("Path finder")
	window.SetContent(content)
	window.ShowAndRun()
}
