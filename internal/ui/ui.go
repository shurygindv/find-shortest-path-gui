package ui

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func FindShortestPathButton(handleClick func()) *widget.Button {
	return widget.NewButton("Find a shortest path", func() {
		log.Println("tapped")
		handleClick()
	})
}

type SelectComboBoxProps struct {
	options []string
	label   string
}

func SelectContainer(props SelectComboBoxProps) (*fyne.Container, *widget.Select) {
	selectWidgetAPI := widget.NewSelect(props.options, func(value string) {
		log.Println("Select set to", value)
	})

	container := container.NewGridWithRows(
		2,
		widget.NewLabel(props.label),
		selectWidgetAPI,
	)

	return container, selectWidgetAPI
}
