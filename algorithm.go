package main

type AlgorithmRunnerOptions struct {
	sourceNodeId      int
	destinationNodeId int
	renderer          GraphRenderer
}

type FindShortestPathAlgorithm struct {
	graph    Graph
	renderer GraphRenderer
}

func (pathAlgorithm *FindShortestPathAlgorithm) run(options AlgorithmRunnerOptions) []int {
	algorithm := DijkstraAlgorithm{
		graph:    pathAlgorithm.graph,
		renderer: pathAlgorithm.renderer,
	}

	return algorithm.run(options)
}
