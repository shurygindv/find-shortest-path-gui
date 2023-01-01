package ui

import (
	"fmt"

	"shortestpath/app/internal/graph"
	"shortestpath/app/pkg"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

func MainAppLayout(graphApi graph.GraphApi) *fyne.Container {
	graphRenderer := GraphRenderer{
		data: graphApi.Data,
	}

	graphLayout := container.NewWithoutLayout(
		graphRenderer.Draw()...,
	)

	// Selects
	options, optionsMap := utils.GenerateNodeOptionsForSelect(
		graphApi.GetNodesCount(),
	)

	sourceSelect, sourceSelectAPI := SelectContainer(SelectComboBoxProps{
		label:   "Source node",
		options: options,
	})

	destinationSelect, destinationSelectAPI := SelectContainer(SelectComboBoxProps{
		label:   "Destination node",
		options: options,
	})

	handleFindButtonClick := func() {
		sourceNodeId := optionsMap[sourceSelectAPI.Selected]
		destinationNodeId := optionsMap[destinationSelectAPI.Selected]

		foundPath := graphApi.FindShortestPath(sourceNodeId, destinationNodeId)

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
