package main

import (
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
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
	algorithm := FindShortestPathAlgorithm{
		graph,
		graphRenderer,
	}

	graphLayout := container.New(
		layout.NewMaxLayout(),
		graphRenderer.Draw(),
	)

	content := container.New(
		layout.NewVBoxLayout(),
		graphLayout,
		layout.NewSpacer(),
		layout.NewSpacer(),
		FindShortestPathButton(func() {
			sourceNode := Node{index: 1, x: 128, y: 97}

			algorithmOptions := AlgorithmRunnerOptions{
				sourceNode,
				graphRenderer,
			}

			algorithm.run(algorithmOptions)
		}),
	)

	window.SetContent(content)
	window.ShowAndRun()
}
