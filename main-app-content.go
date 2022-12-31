package main

import (
	"fmt"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func FindShortestPathButton(handleClick func()) *widget.Button {
	return widget.NewButton("Find shortest path", func() {
		log.Println("tapped")
		handleClick()
	})
}

func MainAppContent() *fyne.Container {
	graph := BuildGraph(ImportFilePointsByPath("data.txt"))
	graphRenderer := GraphRenderer{graph}
	algorithm := FindShortestPathAlgorithm{
		graph,
		graphRenderer,
	}

	graphLayout := container.NewWithoutLayout(
		graphRenderer.Draw()...,
	)

	handleFindButtonClick := func() {
		algorithmOptions := AlgorithmRunnerOptions{
			sourceNodeId:      0, // 1
			destinationNodeId: 3, // 4
			renderer:          graphRenderer,
		}

		foundPath := algorithm.run(algorithmOptions)

		graphLayout.Objects = graphRenderer.DrawHighlightedPath(foundPath)
		graphLayout.Refresh()

		fmt.Print(foundPath)
	}

	content := container.New(
		layout.NewVBoxLayout(),
		graphLayout,
		layout.NewSpacer(),
		FindShortestPathButton(handleFindButtonClick),
	)

	return content
}
