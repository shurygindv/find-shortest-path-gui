package graph

type Node struct {
	ID int // TODO: id is an index (file line)
	X  int
	Y  int
}

type Weight struct {
	speed               int
	length              int
	loadingLevelPercent int
}

type Edge struct {
	Head   Node
	Tail   Node
	weight Weight
}

func (l *Edge) CalculateWeight() int {
	time := float64(l.weight.length / l.weight.speed)
	load := float64(l.weight.loadingLevelPercent) // / 100

	return int((time + 1) * load)
}

type Graph struct {
	Edges []Edge
	Nodes []Node
}
