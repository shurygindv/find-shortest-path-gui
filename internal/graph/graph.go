package graph

import utils "shortestpath/app/pkg"

type Node struct {
	ID int // TODO: id is an index (file line)
	X  int
	Y  int
}

type Weight struct {
	bandwidth           int
	length              int
	loadingLevelPercent int
}

type Edge struct {
	Head   Node
	Tail   Node
	weight Weight
}

func (l *Edge) CalculateWeight() int {
	time := float64(l.weight.length / l.weight.bandwidth)
	load := float64(l.weight.loadingLevelPercent) // / 100

	return int((time + 1) * load)
}

type Graph struct {
	Edges []Edge
	Nodes []Node
}

func (graph *Graph) GetNodesCount() int {
	return len(graph.Nodes)
}

func (graph *Graph) GetNodeIds() []int {
	return utils.Map(graph.Nodes, func(node Node) int {
		return node.ID
	})
}

func (graph *Graph) ConvertTo2DArray() [][]int {
	matrix := utils.GenerateEmpty2DMatrix(graph.GetNodesCount())

	for _, edge := range graph.Edges {
		weight := edge.CalculateWeight()

		// undirected graph
		matrix[edge.Tail.ID][edge.Head.ID] = weight
		matrix[edge.Head.ID][edge.Tail.ID] = weight
	}

	return matrix
}

