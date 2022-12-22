package main

type Node struct {
	index int
	x     int
	y     int
	//
	isVisited bool
	label     float64
}

type Weight struct {
	speed               int
	distance            int
	loadingLevelPercent int
}

type Link struct {
	head   Node
	tail   Node
	weight Weight
}

func (l *Link) calculateWeight() float64 {
	time := float64(l.weight.distance / l.weight.speed)
	load := float64(l.weight.loadingLevelPercent / 100)

	return time * load
}

type Graph struct {
	links []Link
	nodes []Node
}

type Target struct {
	sourceNodeIndex      int
	destinationNodeIndex int
}
