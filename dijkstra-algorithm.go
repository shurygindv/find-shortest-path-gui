package main

import (
	"math"
)

type DijkstraNode struct {
	index     int
	isVisited bool
	distance  float64
}

func createDijkstraNodesData(nodes []Node) map[int]DijkstraNode {
	var data = make(map[int]DijkstraNode)

	for _, node := range nodes {
		data[node.index] = DijkstraNode{node.index, false, math.Inf(1)}
	}

	return data
}

func has_unvisited_nodes(dijkstraNodeData map[int]DijkstraNode) bool {
	for _, node := range dijkstraNodeData {
		if !node.isVisited {
			return true
		}
	}

	return false
}

func find_shortest_path(sourceNode Node, graph Graph) []Node {
	dijkstraNodeData := createDijkstraNodesData(graph.nodes)

	sourceNodeData := dijkstraNodeData[sourceNode.index]

	sourceNodeData.distance = 0
	sourceNodeData.isVisited = true

	currentNode := sourceNodeData

	for has_unvisited_nodes(dijkstraNodeData) {
		adjacentLinks := filter(graph.links, func(link Link) bool {
			return link.tail.index == currentNode.index ||
				link.head.index == currentNode.index
		})

		for _, link := range adjacentLinks {
			var (nodeData DijkstraNode; prevNode Node; nextNode Node)

			if link.head.index == currentNode.index {
				nextNode = link.tail
				prevNode = link.head
			} else {
				nextNode = link.head
				prevNode = link.tail
			}

			nodeData = dijkstraNodeData[nextNode.index]
			distance := link.calculateWeight()

			nodeData.distance = distance + dijkstraNodeData[prevNode.index].distance
		}

		nodeWithMinDistance := dijkstraNodeData[0]

		for _, node := range dijkstraNodeData {
			if !node.isVisited && node.distance < nodeWithMinDistance.distance {
				nodeWithMinDistance = node
			}
		}

		nodeWithMinDistance.isVisited = true
		currentNode = nodeWithMinDistance
	}
	// TODO list
	return []Node{Node{}}
}
