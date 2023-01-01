package main

import (
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"

	"shortestpath/app/internal/graph"
	"shortestpath/app/internal/ui"
)

func main() {
	app := app.New()
	window := app.NewWindow("Main window")

	os.Setenv("FYNE_SCALE", "1.3")

	fileLines := ImportFilePointsByPath("data.txt")
	graphApi := graph.InitializeGraphFromFileStrings(fileLines)

	//
	content := ui.MainAppLayout(graphApi)

	window.CenterOnScreen()
	window.Resize(fyne.NewSize(550, 370))
	window.SetTitle("Path finder")
	window.SetContent(content)
	window.ShowAndRun()
}
