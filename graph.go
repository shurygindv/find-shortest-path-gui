package main

type Node struct {
	id int // TODO: id is an index (file line)
	x  int
	y  int
}

type Weight struct {
	speed               int
	length              int
	loadingLevelPercent int
}

type Edge struct {
	head   Node
	tail   Node
	weight Weight
}

func (l *Edge) CalculateWeight() int {
	time := float64(l.weight.length / l.weight.speed)
	load := float64(l.weight.loadingLevelPercent) // / 100

	return int((time + 1) * load)
}

type Graph struct {
	edges []Edge
	nodes []Node
}

func (graph *Graph) getNodeIds() []int {
	return Map(graph.nodes, func(node Node) int {
		return node.id
	})
}

func (graph *Graph) convertTo2DArray() [][]int {
	matrix := GenerateEmpty2DMatrix(len(graph.nodes))

	for _, edge := range graph.edges {
		weight := edge.CalculateWeight()

		// undirected graph
		matrix[edge.tail.id][edge.head.id] = weight
		matrix[edge.head.id][edge.tail.id] = weight
	}

	return matrix
}
