package graph

type Edge struct {
	time int
	node *Node
}

type Node struct {
	Index     int
	Name      string
	WaitTimes []int
	Edges     []*Edge
	linked    bool
}

func New(idx int, name string, size int) *Node {
	return &Node{
		Index:     idx,
		Name:      name,
		WaitTimes: make([]int, 32),
		Edges:     make([]*Edge, size),
		linked:    false,
	}
}

func (n *Node) rmAdj(idx int) {
	n.Edges = append(n.Edges[:idx], n.Edges[idx+1:]...)
	if idx < n.Index {
		n.Index--
	}
}

func (n *Node) addWaitTime(idx int, time int) {
	n.WaitTimes[idx] = time
	n.linked = true
}

func (n *Node) addWalkTime(idx int, time int, node *Node) {
	n.Edges[idx] = &Edge{
		time: time,
		node: node,
	}
}
