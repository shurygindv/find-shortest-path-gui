package main

import (
	"image/color"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

type GraphRenderer struct {
	data Graph
}

func getNodeCoordinates(node Node) (float32, float32) {
	x := float32(node.x)
	y := float32(node.y)

	return x, y
}

func drawText(name string, color color.Color) *canvas.Text {
	text := canvas.NewText(name, color)
	text.Alignment = fyne.TextAlignTrailing
	text.TextStyle = fyne.TextStyle{Bold: true}
	text.TextSize = 20

	return text
}

func (g *GraphRenderer) DrawAttachedNamesToNode(nodes []Node) []fyne.CanvasObject {
	nodeNames := make([]fyne.CanvasObject, len(nodes))

	for i, node := range nodes {
		nodeName := strconv.Itoa(node.id + 1)

		x, y := getNodeCoordinates(node)

		nodeText := drawText(nodeName, color.White)
		nodeText.Move(fyne.NewPos(x, y))

		nodeNames[i] = nodeText
	}

	return nodeNames
}

func (g *GraphRenderer) DrawLines(edges []Edge) []fyne.CanvasObject {
	lines := make([]fyne.CanvasObject, len(edges))

	for i, edge := range edges {
		line := canvas.NewLine(color.White)

		startNodeX, startNodeY := getNodeCoordinates(edge.tail)
		endNodeX, endNodeY := getNodeCoordinates(edge.head)

		line.Position1 = fyne.NewPos(startNodeX, startNodeY)
		line.Position2 = fyne.NewPos(endNodeX, endNodeY)

		lines[i] = line
	}

	return lines
}

func (g *GraphRenderer) DrawEdgeWeights(edges []Edge) []fyne.CanvasObject {
	weights := make([]fyne.CanvasObject, len(edges))

	for i, edge := range edges {
		weight := strconv.Itoa(edge.CalculateWeight())

		startNodeX, startNodeY := getNodeCoordinates(edge.tail)
		endNodeX, endNodeY := getNodeCoordinates(edge.head)

		weightText := drawText(weight, color.Gray{0x99})
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
	lines := g.DrawLines(g.data.edges)
	nodeNames := g.DrawAttachedNamesToNode(g.data.nodes)
	edgeWeights := g.DrawEdgeWeights(g.data.edges)

	picture := append(lines, nodeNames...)
	picture = append(picture, edgeWeights...)

	return container.NewWithoutLayout(
		picture...,
	)
}
