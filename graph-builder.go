package main

import (
	"fmt"
	"log"
	"strconv"
)

func buildNodes(nodeCount int, fileNodeLines []string) (nodes []Node) {
	nodes = make([]Node, nodeCount)

	for i, line := range fileNodeLines {
		x, y, id := 0, 0, i

		_, err := fmt.Sscanf(line, "%d %d", &x, &y)

		if err != nil {
			log.Fatal(err)
		}

		nodes[i] = Node{id, x, y}
	}

	return
}

func buildEdges(edgesCount int, fileEdgeLines []string, nodes []Node) (edges []Edge) {
	edges = make([]Edge, edgesCount)

	for i, line := range fileEdgeLines {
		var (
			nodeIdA             int
			nodeIdB             int
			speed               int
			loadingLevelPercent int
			distance            int
		)

		_, err := fmt.Sscanf(line, "%d %d %d %d %d", &nodeIdA, &nodeIdB, &speed, &loadingLevelPercent, &distance)

		if err != nil {
			log.Fatal(err)
		}

		tail := Find(nodes, func(node Node) bool {
			return node.id == nodeIdA-1
		})

		head := Find(nodes, func(node Node) bool {
			return node.id == nodeIdB-1
		})

		weight := Weight{speed, distance, loadingLevelPercent}

		edges[i] = Edge{*head, *tail, weight}
	}

	return
}

func BuildGraph(fileLines []string) Graph {
	nodeCount, err := strconv.Atoi(fileLines[0])

	if err != nil {
		log.Fatalf("File: Can't recognize nodes, error: %s", err)
	}

	fileNodeLines := fileLines[1 : nodeCount+1]
	nodes := buildNodes(nodeCount, fileNodeLines)

	edgesCount, err2 := strconv.Atoi(fileLines[nodeCount+1])

	if err2 != nil {
		log.Fatalf("File: Can't recognize edges between nodes, error: %s", err2)
	}

	fileEdgesLines := fileLines[nodeCount+2 : nodeCount+2+edgesCount]
	edges := buildEdges(edgesCount, fileEdgesLines, nodes)

	return Graph{edges, nodes}
}
