package main

type AlgorithmRunnerOptions struct {
	sourceNodeId      int
	destinationNodeId int
}

type FindShortestPathAlgorithm struct {
	graph    Graph
}

func (pathAlgorithm *FindShortestPathAlgorithm) run(options AlgorithmRunnerOptions) []int {
	algorithm := DijkstraAlgorithm{
		graph:    pathAlgorithm.graph,
	}

	return algorithm.run(options)
}
