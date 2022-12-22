package main

import (
	"fmt"
	"log"
	"strconv"
)

func find[T any](array []T, fn func(T) bool) T {
	for _, v := range array {
		if fn(v) {
			return v
		}
	}

	return *new(T)
}

func build_nodes(nodeCount int, fileNodeLines []string) (nodes []Node) {
	nodes = make([]Node, nodeCount)

	for i, line := range fileNodeLines {
		x, y, index := 0, 0, i+1

		_, err := fmt.Sscanf(line, "%d %d", &x, &y)

		if err != nil {
			log.Fatal(err)
		}

		nodes[i] = Node{index, x, y}
	}

	return
}

func build_links(linkCount int, fileLinkLines []string, nodes []Node) (links []Link) {
	links = make([]Link, linkCount)

	for i, line := range fileLinkLines {
		var (
			indexA              int
			indexB              int
			speed               int
			loadingLevelPercent int
			distance            int
		)

		_, err := fmt.Sscanf(line, "%d %d %d %d %d", &indexA, &indexB, &speed, &loadingLevelPercent, &distance)

		if err != nil {
			log.Fatal(err)
		}

		tail := find(nodes, func(node Node) bool {
			return node.index == indexA
		})

		head := find(nodes, func(node Node) bool {
			return node.index == indexB
		})

		weight := Weight{speed, distance, loadingLevelPercent}

		links[i] = Link{head, tail, weight}
	}

	return
}

func build_graph(fileLines []string) Graph {
	nodeCount, err := strconv.Atoi(fileLines[0])

	if err != nil {
		log.Fatalf("File: Can't recognize nodes, error: %s", err)
	}

	fileNodeLines := fileLines[1 : nodeCount+1]
	nodes := build_nodes(nodeCount, fileNodeLines)

	linkCount, err2 := strconv.Atoi(fileLines[nodeCount+1])

	if err2 != nil {
		log.Fatalf("File: Can't recognize links between nodes, error: %s", err2)
	}

	fileLinksLines := fileLines[nodeCount+2 : nodeCount+2+linkCount]
	links := build_links(linkCount, fileLinksLines, nodes)

	return Graph{links}
}
