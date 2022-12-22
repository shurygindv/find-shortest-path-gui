package main

func find_link_with_min_weights(links []Link) Link {
	linkWithMinWeight := links[0]
	minWeight := linkWithMinWeight.head.label

	for _, link := range links {
		if link.head.label < minWeight && link.head.isVisited == false {
			minWeight = link.head.label
			linkWithMinWeight = link
		}
	}

	return linkWithMinWeight
}

func find_shortest_path(target Target, graph Graph) []Node {
	sourceNode := find(graph.nodes, func(node Node) bool {
		return node.index == target.sourceNodeIndex
	})

	sourceNode.label = 0

	for sourceNode.isVisited == false {
		neighboringLinks := filter(graph.links, func(link Link) bool {
			return link.tail.index == sourceNode.index
		})

		for _, link := range neighboringLinks {
			linkWeight := link.calculateWeight()

			if linkWeight < link.head.label {
				link.head.label = linkWeight
			}
		}

		nearestLink := find_link_with_min_weights(neighboringLinks)

		sourceNode.isVisited = true

		sourceNode = nearestLink.head
	}

	// TODO list
	return []Node{Node{}}
}
