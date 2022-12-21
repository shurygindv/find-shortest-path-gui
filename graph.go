package main

type Node struct {
	index int
	x int
	y int
}

type Weight struct {
  speed int 
  distance int
  loadingLevelPercent int
}

type Link struct {
	head Node
	tail Node 
	weight Weight
}

type Graph struct {
	links []Link
}