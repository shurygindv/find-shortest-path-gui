package main

import (
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

func main() {
	app := app.New()
	window := app.NewWindow("Main window")

	os.Setenv("FYNE_SCALE", "1.5")

	window.CenterOnScreen()
	window.Resize(fyne.NewSize(500, 300))
	window.SetTitle("Path finder")

	graph := build_graph(import_file_points_by_path("data.txt"))
	graphRenderer := GraphRenderer{graph}

	content := container.NewWithoutLayout(
		container.NewVBox(),
		graphRenderer.Draw(),
	)

	window.SetContent(content)
	window.ShowAndRun()
}
