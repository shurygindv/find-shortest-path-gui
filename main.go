package main

import (
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"

	"shortestpath/app/internal/graph"
	"shortestpath/app/internal/ui"
)

func main() {
	os.Setenv("FYNE_SCALE", "1.3")

	app := app.New()
	window := app.NewWindow("Main window")

	graphApi := graph.CreateApi(ImportFilePointsByPath("data.txt"))

	content := ui.MainAppLayout(graphApi) // internal/ui/main-app-layout.go

	window.CenterOnScreen()
	window.Resize(fyne.NewSize(550, 370))
	window.SetTitle("Path finder")
	window.SetContent(content)
	window.ShowAndRun()
}
