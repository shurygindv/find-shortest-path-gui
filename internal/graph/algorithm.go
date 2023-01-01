package graph

type AlgorithmRunnerOptions struct {
	sourceNodeId      int
	destinationNodeId int
}

type FindShortestPathAlgorithm struct {
	graphApi GraphApi
}

func (pathAlgorithm *FindShortestPathAlgorithm) run(options AlgorithmRunnerOptions) []int {
	algorithm := DijkstraAlgorithm{
		graphApi: pathAlgorithm.graphApi,
	}

	return algorithm.run(options)
}
