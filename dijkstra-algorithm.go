package main

import (
	"container/list"
	"math"
)

type DijkstraAlgorithm struct {
	graph    Graph
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

func (dijkstraAlgorithm *DijkstraAlgorithm) createInfiniteDistances() map[int]int {
	distance := make(map[int]int)

	for _, nodeId := range dijkstraAlgorithm.graph.getNodeIds() {
		distance[nodeId] = INFINITE
	}

	return distance
}

// TODO: may rethink
func (dijkstraAlgorithm *DijkstraAlgorithm) run(data AlgorithmRunnerOptions) []int {
	graph := dijkstraAlgorithm.graph
	//
	hasVisitedNodeId := make(map[int]bool)
	shortestPathToAnyNode := make([]int, len(graph.nodes))
	//
	graphMatrix := graph.convertTo2DArray()
	distance := dijkstraAlgorithm.createInfiniteDistances()
	//
	distance[data.sourceNodeId] = 0

	currentNodeId := data.sourceNodeId

	for !hasVisitedNodeId[currentNodeId] {
		adjacentNodes := graphMatrix[currentNodeId]

		for nodeId := range adjacentNodes {
			adjacentNodeWeight := adjacentNodes[nodeId]

			if hasVisitedNodeId[nodeId] || adjacentNodeWeight == 0 {
				continue
			}

			prevPlusNextDistance := distance[currentNodeId] + adjacentNodeWeight

			if prevPlusNextDistance < distance[nodeId] {
				distance[nodeId] = prevPlusNextDistance

				shortestPathToAnyNode[nodeId] = currentNodeId
			}
		}

		hasVisitedNodeId[currentNodeId] = true
		minDistance := INFINITE

		for nodeId := range adjacentNodes {
			if hasVisitedNodeId[nodeId] || adjacentNodes[nodeId] == 0 {
				continue
			}

			if distance[nodeId] < minDistance {
				minDistance = distance[nodeId]
				currentNodeId = nodeId
			}
		}
	}

	return dijkstraAlgorithm.extractShortestPathFor(shortestPathToAnyNode, ExtractPathParams{
		sourceNodeId:      data.sourceNodeId,
		destinationNodeId: data.destinationNodeId,
	})
}
