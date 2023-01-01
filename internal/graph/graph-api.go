package graph

import (
	utils "shortestpath/app/pkg"
)

type GraphApi struct {
	Data Graph
}

func InitializeGraphFromFileStrings(fileLines []string) GraphApi {
	graph := BuildGraph(fileLines)

	return GraphApi{
		Data: graph,
	}
}

func (graphApi *GraphApi) FindShortestPath(sourceNodeId int, destinationNodeId int) []int {
	algorithm := FindShortestPathAlgorithm{*graphApi}

	return algorithm.run(AlgorithmRunnerOptions{
		sourceNodeId,
		destinationNodeId,
	})
}

func (graphApi *GraphApi) GetNodesCount() int {
	return len(graphApi.Data.Nodes)
}

func (graphApi *GraphApi) GetNodeIds() []int {
	return utils.Map(graphApi.Data.Nodes, func(node Node) int {
		return node.ID
	})
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
