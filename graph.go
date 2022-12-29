package main

type Node struct {
	id int
	x  int
	y  int
}

type Weight struct {
	speed               int
	length              int
	loadingLevelPercent int
}

type Link struct {
	head   Node
	tail   Node
	weight Weight
}

func (l *Link) calculateWeight() int {
	time := float64(l.weight.length / l.weight.speed)
	load := float64(l.weight.loadingLevelPercent) // / 100

	return int((time + 1) * load)
}

type Graph struct {
	links []Link
	nodes []Node
}

func (graph *Graph) convertTo2DArray() [][]int {
	matrix := GenerateEmpty2DMatrix(len(graph.nodes))

	for _, link := range graph.links {
		weight := link.calculateWeight()

		// undirected graph
		matrix[link.tail.id][link.head.id] = weight
		matrix[link.head.id][link.tail.id] = weight
	}

	return matrix
}
