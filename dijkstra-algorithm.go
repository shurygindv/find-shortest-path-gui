package main

type DijkstraAlgorithm struct {
	graph    Graph
	renderer GraphRenderer
}

type DijkstraNode struct {
	index     int
	isVisited bool
	distance  float64
}

func (dijkstraAlgorithm *DijkstraAlgorithm) run(options AlgorithmRunnerOptions) []Node {
	// TODO rethink
	dijkstraAlgorithm.graph.convertTo2DArray()

	return []Node{Node{}}
}
