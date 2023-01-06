package graph

import (
	"container/list"
	"math"
)

type DijkstraAlgorithm struct {
	graph Graph
}

const INFINITE = math.MaxInt64

type ExtractPathParams struct {
	sourceNodeId      int
	destinationNodeId int
}

func (dijkstraAlgorithm *DijkstraAlgorithm) extractShortestPathFor(
	shortestPathForAnyNodes []int,
	params ExtractPathParams,
) []int {
	sourceNodeId := params.sourceNodeId
	currentNodeId := params.destinationNodeId
	shortestPathList := list.New()

	shortestPathList.PushFront(currentNodeId)

	for shortestPathList.Front().Value != sourceNodeId {
		nextNodeId := shortestPathForAnyNodes[currentNodeId]
		currentNodeId = nextNodeId

		shortestPathList.PushFront(currentNodeId)
	}

	result := make([]int, 0)

	for element := shortestPathList.Front(); element != nil; element = element.Next() {
		result = append(result, element.Value.(int))
	}

	return result
}

func (algorithm *DijkstraAlgorithm) createInfiniteDistances() map[int]int {
	distance := make(map[int]int)

	for _, nodeId := range algorithm.graph.GetNodeIds() {
		distance[nodeId] = INFINITE
	}

	return distance
}

// TODO: may rethink
func (algorithm *DijkstraAlgorithm) run(data AlgorithmRunnerOptions) []int {
	graph := algorithm.graph
	//
	graphMatrix := graph.ConvertTo2DArray()
	distance := algorithm.createInfiniteDistances()
	isReachedNodeId := make(map[int]bool)
	shortestPathToAnyNode := make([]int, graph.GetNodesCount())
	//
	distance[data.sourceNodeId] = 0

	currentNodeId := data.sourceNodeId

	for !isReachedNodeId[currentNodeId] {
		neighbors := graphMatrix[currentNodeId]

		for nodeId := range neighbors {
			neighborNodeWeight := neighbors[nodeId]

			if isReachedNodeId[nodeId] || neighborNodeWeight == 0 {
				continue
			}

			prevPlusNextDistance := distance[currentNodeId] + neighborNodeWeight

			if prevPlusNextDistance < distance[nodeId] {
				distance[nodeId] = prevPlusNextDistance

				shortestPathToAnyNode[nodeId] = currentNodeId
			}
		}

		isReachedNodeId[currentNodeId] = true
		minDistance := INFINITE

		for nodeId := range neighbors {
			if isReachedNodeId[nodeId] || neighbors[nodeId] == 0 {
				continue
			}

			if distance[nodeId] < minDistance {
				minDistance = distance[nodeId]
				currentNodeId = nodeId
			}
		}
	}

	return algorithm.extractShortestPathFor(shortestPathToAnyNode, ExtractPathParams{
		sourceNodeId:      data.sourceNodeId,
		destinationNodeId: data.destinationNodeId,
	})
}
