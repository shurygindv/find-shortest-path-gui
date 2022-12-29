package main

import (
	"log"

	"fyne.io/fyne/v2/widget"
)

func FindShortestPathButton(handleClick func()) *widget.Button {
	return widget.NewButton("Find shortest path", func() {
		log.Println("tapped")
		handleClick()
	})
}
