package main

import "math"

type DijkstraAlgorithm struct {
	graph    Graph
	renderer GraphRenderer
}

type DijkstraNode struct {
	index     int
	isVisited bool
	distance  float64
}

const INFINITY = math.MaxInt64
const VISITED_NODE = true
const UNVISITED_NODE_YET = false

func (dijkstraAlgorithm *DijkstraAlgorithm) run(data AlgorithmRunnerOptions) []int {
	hasVisitedNodeId := make(map[int]bool)
	distance := make(map[int]int)
	path := make([]int, 0)
	//
	graph := dijkstraAlgorithm.graph
	//	destinationNodeId := data.destinationNodeId
	sourceNodeId := data.sourceNodeId
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
			}
		}

		path = append(path, currentNodeId)
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

	return path
}
