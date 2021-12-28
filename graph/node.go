package graph

type Edge struct {
	time int
	node *Node
}

type Node struct {
	_idx       int
	_name      string
	_waitTimes []int
	_edges     []*Edge
}
