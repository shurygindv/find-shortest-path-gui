package graph

import (
	utils "shortestpath/app/pkg"
)

type GraphApi struct {
	Data Graph
}

func CreateApi(fileLines []string) GraphApi {
	return GraphApi{
		Data: BuildGraph(fileLines),
	}
}

func (graphApi *GraphApi) FindShortestPath(sourceNodeId int, destinationNodeId int) []int {
	algorithm := FindShortestPathAlgorithm{graphApi.Data}

	return algorithm.run(AlgorithmRunnerOptions{
		sourceNodeId,
		destinationNodeId,
	})
}

func (graphApi *GraphApi) GetNodesCount() int {
	return graphApi.Data.GetNodesCount()
}

func (graphApi *GraphApi) ConvertGraphTo2DArray() [][]int {
	matrix := utils.GenerateEmpty2DMatrix(graphApi.GetNodesCount())

	for _, edge := range graphApi.Data.Edges {
		weight := edge.CalculateWeight()

		// undirected graph
		matrix[edge.Tail.ID][edge.Head.ID] = weight
		matrix[edge.Head.ID][edge.Tail.ID] = weight
	}

	return matrix
}
