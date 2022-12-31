package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

func MainAppLayout() *fyne.Container {
	// Graph layout
	graph := BuildGraph(ImportFilePointsByPath("data.txt"))
	graphRenderer := GraphRenderer{graph}
	algorithm := FindShortestPathAlgorithm{
		graph,
	}

	graphLayout := container.NewWithoutLayout(
		graphRenderer.Draw()...,
	)

	// Selects layout
	options, optionsMap := GenerateNodeOptionsForSelect(len(graph.nodes))

	sourceSelect, sourceSelectAPI := SelectContainer(SelectComboBoxProps{
		label:   "Source node",
		options: options,
	})

	destinationSelect, destinationSelectAPI := SelectContainer(SelectComboBoxProps{
		label:   "Destination node",
		options: options,
	})

	handleFindButtonClick := func() {
		algorithmOptions := AlgorithmRunnerOptions{
			sourceNodeId:      optionsMap[sourceSelectAPI.Selected],
			destinationNodeId: optionsMap[destinationSelectAPI.Selected],
		}

		foundPath := algorithm.run(algorithmOptions)

		graphLayout.Objects = graphRenderer.DrawHighlightedPath(foundPath)
		graphLayout.Refresh()

		fmt.Print(foundPath)
	}

	selectNodesContainer := container.NewGridWithColumns(
		3,
		sourceSelect,
		destinationSelect,
		FindShortestPathButton(handleFindButtonClick),
	)

	// Result
	content := container.New(
		layout.NewVBoxLayout(),
		selectNodesContainer,
		graphLayout,
		layout.NewSpacer(),
	)

	return content
}
