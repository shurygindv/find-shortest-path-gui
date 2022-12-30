package main

import (
	"fmt"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

func main() {
	window := createWindow("Main window")

	graph := build_graph(import_file_points_by_path("data.txt"))
	graphRenderer := GraphRenderer{graph}
	algorithm := FindShortestPathAlgorithm{
		graph,
		graphRenderer,
	}

	graphLayout := container.New(
		layout.NewMaxLayout(),
		graphRenderer.Draw(),
	)

	handleFindButtonClick := func() {
		algorithmOptions := AlgorithmRunnerOptions{
			sourceNodeId:      0, // 1
			destinationNodeId: 3, // 4
			renderer:          graphRenderer,
		}

		path := algorithm.run(algorithmOptions)

		fmt.Println(path)
	}

	content := container.New(
		layout.NewVBoxLayout(),
		graphLayout,
		layout.NewSpacer(),
		layout.NewSpacer(),
		FindShortestPathButton(handleFindButtonClick),
	)

	window.SetContent(content)
	window.ShowAndRun()
}

func createWindow(name string) fyne.Window {
	app := app.New()
	window := app.NewWindow(name)

	os.Setenv("FYNE_SCALE", "1.5")

	window.CenterOnScreen()
	window.Resize(fyne.NewSize(500, 300))
	window.SetTitle("Path finder")

	return window
}
