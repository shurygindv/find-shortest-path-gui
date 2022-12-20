package main

import (
	"fmt"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func main() {
	app := app.New()
	window := app.NewWindow("Main window")

	window.CenterOnScreen()
	window.Resize(fyne.NewSize(700, 500))
	window.SetTitle("Path finder")

	content := widget.NewButton("Import file", func() {
		dialog.ShowFileOpen(func(uri fyne.URIReadCloser, err error) {
			if err != nil {
				log.Fatal(err)
				return
			}

			var points = import_file_points_stream(uri)

			fmt.Print(len(points))
		}, window)

	})

	window.SetContent(content)
	window.ShowAndRun()
}
