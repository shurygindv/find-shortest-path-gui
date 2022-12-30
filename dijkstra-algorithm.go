package main

import (
	"container/list"
	"math"
)

type DijkstraAlgorithm struct {
	graph    Graph
	renderer GraphRenderer
}

const INFINITY = math.MaxInt64
const VISITED_NODE = true
const UNVISITED_NODE_YET = false

func (dijkstraAlgorithm *DijkstraAlgorithm) run(data AlgorithmRunnerOptions) []int {
	graph := dijkstraAlgorithm.graph
	destinationNodeId := data.destinationNodeId
	sourceNodeId := data.sourceNodeId
	//
	hasVisitedNodeId := make(map[int]bool)
	distance := make(map[int]int)
	shortestPathToAnyNode := make([]int, len(graph.nodes))
	//
	//
	graphMatrix := graph.convertTo2DArray()
	//
	for _, nodeId := range graph.getNodeIds() {
		distance[nodeId] = INFINITY
	}

	distance[sourceNodeId] = 0

	currentNodeId := data.sourceNodeId

	for !hasVisitedNodeId[currentNodeId] {
		adjacentNodes := graphMatrix[currentNodeId]

		for nodeId := range adjacentNodes {
			adjacentNodeWeight := adjacentNodes[nodeId]

			if hasVisitedNodeId[nodeId] || adjacentNodeWeight == 0 {
				continue
			}

			if (distance[currentNodeId] + adjacentNodeWeight) < distance[nodeId] {
				distance[nodeId] = distance[currentNodeId] + adjacentNodeWeight
				shortestPathToAnyNode[nodeId] = currentNodeId
			}
		}

		hasVisitedNodeId[currentNodeId] = VISITED_NODE
		minDistance := INFINITY

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

	//
	shortestPathList := list.New()

	shortestPathList.PushFront(destinationNodeId)
	currentNodeId = destinationNodeId

	for shortestPathList.Front().Value != sourceNodeId {
		nextNodeId := shortestPathToAnyNode[currentNodeId]
		currentNodeId = nextNodeId

		shortestPathList.PushFront(currentNodeId)
	}

	result := make([]int, 0)

	for element := shortestPathList.Front(); element != nil; element = element.Next() {
		result = append(result, element.Value.(int))
	}

	return result
}
