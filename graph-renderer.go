package main

import (
	"fmt"
	"image/color"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

type GraphRenderer struct {
	data Graph
}

func get_node_coordinates(node Node) (float32, float32) {
	x := float32(node.x)
	y := float32(node.y)

	return x, y
}

func draw_text(name string, color color.Color) *canvas.Text {
	text := canvas.NewText(name, color)
	text.Alignment = fyne.TextAlignTrailing
	text.TextStyle = fyne.TextStyle{Bold: true}
	text.TextSize = 20

	return text
}

func (g *GraphRenderer) DrawAttachedNamesToNode(nodes []Node) []fyne.CanvasObject {
	nodeNames := make([]fyne.CanvasObject, len(nodes))

	for i, node := range nodes {
		nodeName := strconv.Itoa(node.index)

		x, y := get_node_coordinates(node)

		nodeText := draw_text(nodeName, color.White)
		nodeText.Move(fyne.NewPos(x, y))

		nodeNames[i] = nodeText
	}

	return nodeNames
}

func (g *GraphRenderer) DrawLines(links []Link) []fyne.CanvasObject {
	lines := make([]fyne.CanvasObject, len(links))

	for i, link := range links {
		line := canvas.NewLine(color.White)

		startNodeX, startNodeY := get_node_coordinates(link.tail)
		endNodeX, endNodeY := get_node_coordinates(link.head)

		line.Position1 = fyne.NewPos(startNodeX, startNodeY)
		line.Position2 = fyne.NewPos(endNodeX, endNodeY)

		lines[i] = line
	}

	return lines
}

func (g *GraphRenderer) DrawLinkWeights(links []Link) []fyne.CanvasObject {
	weights := make([]fyne.CanvasObject, len(links))

	for i, link := range links {
		weight := fmt.Sprintf("%.0f", link.calculateWeight())

		startNodeX, startNodeY := get_node_coordinates(link.tail)
		endNodeX, endNodeY := get_node_coordinates(link.head)

		weightText := draw_text(weight, color.Gray{0x99})
		weightText.TextSize = 12

		weightText.Move(fyne.NewPos(
			(startNodeX+endNodeX)/2,
			(startNodeY+endNodeY)/2,
		))

		weights[i] = weightText
	}

	return weights
}

func (g *GraphRenderer) Draw() *fyne.Container {
	lines := g.DrawLines(g.data.links)
	nodeNames := g.DrawAttachedNamesToNode(g.data.nodes)
	linkWeights := g.DrawLinkWeights(g.data.links)

	picture := append(lines, nodeNames...)
	picture = append(picture, linkWeights...)

	return container.NewWithoutLayout(
		picture...,
	)
}
