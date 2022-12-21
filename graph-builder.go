package main

import (
	"fmt"
	"log"
	"strconv"
)

func build_nodes(nodeCount int, fileNodeLines []string) (nodes []Node) {
	nodes = make([]Node, nodeCount)

	for i, line := range fileNodeLines {
		var x, y, index = 0, 0, i + 1

		_, err := fmt.Sscanf(line, "%d %d", &x, &y)

		if err != nil {
			log.Fatal(err)
		}

		nodes = append(nodes, Node{x, y, index})
	}

	return
}

func build_graph(fileLines []string) {
	nodeCount, err := strconv.Atoi(fileLines[0])

	if err != nil {
		log.Fatalf("File: Can't recognize vertices, error: %s", err)
	}

	fileNodeLines := fileLines[1 : nodeCount+1]
	nodes := build_nodes(nodeCount, fileNodeLines)
}
