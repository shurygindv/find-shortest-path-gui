package main

import (
	"image/color"
	"os"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func drawGraph(g Graph) *fyne.Container {
	var text []fyne.CanvasObject
	lines := make([]fyne.CanvasObject, len(g.links))

	for i, link := range g.links {
		line := canvas.NewLine(color.White)

		startNodeName := strconv.Itoa(link.head.index)
		startNodeX := float32(link.tail.x)
		startNodeY := float32(link.tail.y)

		startNodeText := canvas.NewText(startNodeName, color.White)
		startNodeText.Alignment = fyne.TextAlignTrailing
		startNodeText.TextStyle = fyne.TextStyle{Bold: true}
		startNodeText.TextSize = 20
		//
		endNodeName := strconv.Itoa(link.tail.index)
		endNodeX := float32(link.head.x)
		endNodeY := float32(link.head.y)

		endNodeText := canvas.NewText(endNodeName, color.White)
		endNodeText.Alignment = fyne.TextAlignTrailing
		endNodeText.TextStyle = fyne.TextStyle{Bold: true}
		endNodeText.TextSize = 20

		//
		line.Position1 = fyne.NewPos(startNodeX, startNodeY)
		line.Position2 = fyne.NewPos(endNodeX, endNodeY)

		startNodeText.Move(fyne.NewPos(startNodeX, startNodeY))
		endNodeText.Move(fyne.NewPos(endNodeX+10, endNodeY+10))

		lines[i] = line
		text = append(text, startNodeText, endNodeText)
	}

	var picture = append(lines, text...)

	return container.NewWithoutLayout(
		picture...,
	)

}

func main() {
	app := app.New()
	window := app.NewWindow("Main window")

	os.Setenv("FYNE_SCALE", "1.5")

	window.CenterOnScreen()
	window.Resize(fyne.NewSize(500, 300))
	window.SetTitle("Path finder")

	graph := build_graph(import_file_points_by_path("data.txt"))
	content := drawGraph(graph)

	window.SetContent(content)
	window.ShowAndRun()
}
