package ui

import (
	"image/color"
	"log"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"

	"shortestpath/app/internal/graph"
	"shortestpath/app/pkg"
)

type GraphRenderer struct {
	data graph.Graph
}

func getNodeCoordinates(node graph.Node) (float32, float32) {
	x := float32(node.X)
	y := float32(node.Y)

	return x, y
}

func drawText(name string, color color.Color) *canvas.Text {
	text := canvas.NewText(name, color)
	text.Alignment = fyne.TextAlignTrailing
	text.TextStyle = fyne.TextStyle{Bold: true}
	text.TextSize = 20

	return text
}

func (g *GraphRenderer) DrawAttachedNamesToNode(nodes []graph.Node) []fyne.CanvasObject {
	nodeNames := make([]fyne.CanvasObject, len(nodes))

	for i, node := range nodes {
		nodeName := strconv.Itoa(node.ID + 1)

		x, y := getNodeCoordinates(node)

		nodeText := drawText(nodeName, color.White)
		nodeText.Move(fyne.NewPos(x, y))

		nodeNames[i] = nodeText
	}

	return nodeNames
}

func (g *GraphRenderer) DrawLinesBetweenNodes(tail graph.Node, head graph.Node) *canvas.Line {
	line := canvas.NewLine(color.White)

	startNodeX, startNodeY := getNodeCoordinates(tail)
	endNodeX, endNodeY := getNodeCoordinates(head)

	line.Position1 = fyne.NewPos(startNodeX, startNodeY)
	line.Position2 = fyne.NewPos(endNodeX, endNodeY)

	return line
}

func (g *GraphRenderer) DrawLines(edges []graph.Edge) []fyne.CanvasObject {
	lines := make([]fyne.CanvasObject, len(edges))

	for i, edge := range edges {
		lines[i] = g.DrawLinesBetweenNodes(edge.Tail, edge.Head)
	}

	return lines
}

func (g *GraphRenderer) DrawEdgeWeights(edges []graph.Edge) []fyne.CanvasObject {
	weights := make([]fyne.CanvasObject, len(edges))

	for i, edge := range edges {
		weight := strconv.Itoa(edge.CalculateWeight())

		startNodeX, startNodeY := getNodeCoordinates(edge.Tail)
		endNodeX, endNodeY := getNodeCoordinates(edge.Head)

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

func (g *GraphRenderer) Draw() []fyne.CanvasObject {
	lines := g.DrawLines(g.data.Edges)
	nodeNames := g.DrawAttachedNamesToNode(g.data.Nodes)
	edgeWeights := g.DrawEdgeWeights(g.data.Edges)

	renderedGraph := append(lines, nodeNames...)
	renderedGraph = append(renderedGraph, edgeWeights...)

	return renderedGraph
}

func (g *GraphRenderer) DrawHighlightedPath(nodePathIds []int) []fyne.CanvasObject {
	orderedNodes := make([]graph.Node, len(nodePathIds))

	for i, nodeId := range nodePathIds {
		found, isEmpty := utils.Find2(g.data.Nodes, func(node graph.Node) bool {
			return nodeId == node.ID
		})

		if isEmpty {
			log.Fatalf("Can't create a final path, missing node")
		}

		orderedNodes[i] = *found
	}

	lines := make([]fyne.CanvasObject, 0)

	for i := range orderedNodes[1:] {
		line := g.DrawLinesBetweenNodes(orderedNodes[i+1], orderedNodes[i])

		line.StrokeWidth = 2
		line.StrokeColor = color.NRGBA{0xff, 0x00, 0x00, 0xff}

		lines = append(lines, line)
	}

	return append(g.Draw(), lines...)
}
