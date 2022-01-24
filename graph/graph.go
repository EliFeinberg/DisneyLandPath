package graph

import (
	"math/rand"
	"strconv"
	"sync"
	"time"

	"github.com/EliFeinberg/DisneyLandPath/utils"
)

const (
	MAX_INT = int(^uint(0) >> 1)
)

type Graph struct {
	Nodes []*Node
}

func BuildGraph(rideCSV string, walkCSV string, idCSV string) Graph {
	_graphNodes := make([]*Node, 0)
	// parse CSV files
	rideRows := utils.ParseCSV(rideCSV)
	rideRows = rideRows[1:] // remove header row
	walkRows := utils.ParseCSV(walkCSV)
	walkRows = walkRows[1:] // remove header row
	idRows := utils.ParseCSV(idCSV)
	idRows = idRows[1:] // remove header row

	// fmt.Println("CSV files parsed")
	// build graph

	// setup graph nodes slice
	// Maps ride id to map index for use during traversal
	idMap := make(map[int]int)
	i := 0
	for _, ride := range idRows {
		rideID, err := strconv.Atoi(ride[0])
		if err != nil {
			panic(err)
		}
		_graphNodes = append(_graphNodes, New(i, ride[2], len(idRows)))
		_graphNodes[i].addWalkTime(i, -1, _graphNodes[i])
		idMap[rideID] = i
		i++
	}
	// updates each node with walking data
	for _, walk := range walkRows {
		rideIDA, err := strconv.Atoi(walk[0])
		if err != nil {
			panic(err)
		}
		rideIDB, err := strconv.Atoi(walk[1])
		if err != nil {
			panic(err)
		}
		rideA, existA := idMap[rideIDA]
		rideB, existB := idMap[rideIDB]
		if existA && existB {
			walkTime, err := strconv.Atoi(walk[2])
			if err != nil {
				panic(err)
			}

			_graphNodes[rideA].addWalkTime(rideB, walkTime, _graphNodes[rideB])
			_graphNodes[rideB].addWalkTime(rideA, walkTime, _graphNodes[rideA])
		}
	}
	// adds ride wait times to each node
	for _, rideTime := range rideRows {
		val, err := strconv.Atoi(rideTime[0])
		if err != nil {
			panic(err)
		}
		ride, exist := idMap[val]
		if exist {
			waitTime := utils.CalculateIdx(rideTime[1])
			Time, err := strconv.Atoi(rideTime[2])
			if err != nil {
				panic(err)
			}

			if waitTime < 16 {
				continue
			}
			_graphNodes[ride].addWaitTime(waitTime-16, Time)

		}

	}

	for i := 0; i < len(_graphNodes); i++ {
		if !_graphNodes[i].linked {
			for j := 0; j < len(_graphNodes); j++ {
				_graphNodes[j].rmAdj(i)
			}
			_graphNodes = append(_graphNodes[:i], _graphNodes[i+1:]...)
			i--
		}
	}
	//Testing
	// for _, node := range _graphNodes {
	// 	for i := 0; i < 32; i++ {
	// 		fmt.Println("Wait time at", node.Name, "at", i, ": ", node.WaitTimes[i])
	// 	}
	// }

	return Graph{
		Nodes: _graphNodes,
	}
}

type queueItem struct {
	node *Node
	time int
	mask int
}

// Dijkstra's algorithm traversal of graph
func (g *Graph) Traversal(startTime int) ([]*Node, int) {
	costMatrix := make([][]int, len(g.Nodes))
	for i := range costMatrix {
		costMatrix[i] = make([]int, 1<<uint(len(g.Nodes)))
		for j := range costMatrix[i] {
			costMatrix[i][j] = MAX_INT
		}
	}
	prevMatrix := make([][]*Node, len(g.Nodes))
	for i := range prevMatrix {
		prevMatrix[i] = make([]*Node, 1<<uint(len(g.Nodes)))
		for j := range prevMatrix[i] {
			prevMatrix[i][j] = nil
		}
	}
	q := make([]queueItem, 0)

	// Shuffle nodes to determine if starting point matters
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(q), func(i, j int) { q[i], q[j] = q[j], q[i] })

	for i := 0; i < len(g.Nodes); i++ {
		costMatrix[i][1<<uint(i)] = 0
		prevMatrix[i][1<<uint(i)] = g.Nodes[i]
		q = append(q, queueItem{g.Nodes[i], startTime, 1 << uint(i)})
	}

	for len(q) > 0 {
		item := q[0]
		q = q[1:]
		currIdx, currTime, currMask := item.node.Index, item.time, item.mask
		// fmt.Println("Current node:", currIdx, "Current time:", currTime, "Current mask:", currMask)
		if currTime/30 > 31 {
			continue
		}
		for _, edge := range item.node.Edges {
			if edge.node == nil || edge.node.Index == currIdx {
				continue
			}
			edgeIdx := edge.node.Index
			edgeTime := edge.time + item.node.WaitTimes[currTime/30]
			edgeMask := currMask | (1 << uint(edgeIdx))
			// fmt.Println("Edge node:", edgeIdx, "Edge time:", edgeTime, "Edge mask:", edgeMask)
			if costMatrix[edgeIdx][edgeMask] > costMatrix[currIdx][currMask]+edgeTime {
				costMatrix[edgeIdx][edgeMask] = costMatrix[currIdx][currMask] + edgeTime
				prevMatrix[edgeIdx][edgeMask] = item.node
				Time := costMatrix[currIdx][currMask] + edgeTime + currTime
				q = append(q, queueItem{edge.node, Time, edgeMask})
			}
		}
	}

	// for i := 0; i < len(costMatrix); i++ {
	// 	fmt.Println(costMatrix[i][(1<<uint(len(g.Nodes)))-1])
	// }

	path := make([]*Node, 0)
	endIdx, totTime, currMask := -1, MAX_INT, uint((1<<uint(len(g.Nodes)))-1)
	for j := 1; j < 1<<len(g.Nodes); j++ {
		for i := 0; i < len(g.Nodes); i++ {
			if totTime > costMatrix[i][1<<len(g.Nodes)-j] {
				totTime = costMatrix[i][1<<len(g.Nodes)-j]
				endIdx = i
				currMask = uint((1 << len(g.Nodes)) - j)
			}
		}
		if endIdx != -1 {
			break
		}
	}

	curr := g.Nodes[endIdx]
	path = append(path, curr)
	for curr != nil {
		indx := curr.Index
		curr = prevMatrix[indx][currMask]
		if curr == nil || curr.Index == indx {
			break
		}
		currMask = currMask & (currMask - (1 << uint(indx)))
		path = append(path, curr)
	}
	return reverse(path), totTime
}

func (g *Graph) TraversalGo(startTime int) ([]*Node, int) {
	var wg sync.WaitGroup
	costMatrix := make([][]int, len(g.Nodes))
	for i := range costMatrix {
		costMatrix[i] = make([]int, 1<<uint(len(g.Nodes)))
		for j := range costMatrix[i] {
			costMatrix[i][j] = MAX_INT
		}
	}
	prevMatrix := make([][]*Node, len(g.Nodes))
	for i := range prevMatrix {
		prevMatrix[i] = make([]*Node, 1<<uint(len(g.Nodes)))
		for j := range prevMatrix[i] {
			prevMatrix[i][j] = nil
		}
	}
	q := make([]queueItem, 0)
	for i := 0; i < len(g.Nodes); i++ {

		costMatrix[i][1<<uint(i)] = 0
		prevMatrix[i][1<<uint(i)] = g.Nodes[i]
		// fmt.Printf("Queue memory loc %p\n", &q)
		q = append(q, queueItem{g.Nodes[i], startTime, 1 << uint(i)})

	}

	for i := 0; i < len(g.Nodes); i++ {
		wg.Add(1)
		// fmt.Println("Starting at node", g.Nodes[i].Name, " work #", i)
		go workerDijkstra(q, costMatrix, prevMatrix, &wg)

	}

	wg.Wait()

	// for i := 0; i < len(costMatrix); i++ {
	// 	fmt.Println(costMatrix[i][(1<<uint(len(g.Nodes)))-1])
	// }

	path := make([]*Node, 0)
	endIdx, totTime, currMask := -1, MAX_INT, uint((1<<uint(len(g.Nodes)))-1)
	for j := 1; j < 1<<len(g.Nodes); j++ {
		for i := 0; i < len(g.Nodes); i++ {
			if totTime > costMatrix[i][1<<len(g.Nodes)-j] {
				totTime = costMatrix[i][1<<len(g.Nodes)-j]
				endIdx = i
				currMask = uint((1 << len(g.Nodes)) - j)
			}
		}
		if endIdx != -1 {
			break
		}
	}

	curr := g.Nodes[endIdx]
	path = append(path, curr)
	for curr != nil {
		indx := curr.Index
		curr = prevMatrix[indx][currMask]
		if curr == nil || curr.Index == indx {
			break
		}
		currMask = currMask & (currMask - (1 << uint(indx)))
		path = append(path, curr)
	}
	return reverse(path), totTime
}

func workerDijkstra(q []queueItem, costMatrix [][]int, prevMatrix [][]*Node, wg *sync.WaitGroup) {
	defer wg.Done()
	for len(q) > 0 {
		item := q[0]
		q = q[1:]
		currIdx, currTime, currMask := item.node.Index, item.time, item.mask
		// fmt.Println("Current node:", currIdx, "Current time:", currTime, "Current mask:", currMask)

		if currTime/30 > 31 {
			continue
		}
		for _, edge := range item.node.Edges {
			if edge.node == nil || edge.node.Index == currIdx {
				continue
			}
			edgeIdx := edge.node.Index
			edgeTime := edge.time + item.node.WaitTimes[currTime/30]
			edgeMask := currMask | (1 << uint(edgeIdx))
			// fmt.Println("Edge node:", edgeIdx, "Edge time:", edgeTime, "Edge mask:", edgeMask)
			if costMatrix[edgeIdx][edgeMask] > costMatrix[currIdx][currMask]+edgeTime {
				costMatrix[edgeIdx][edgeMask] = costMatrix[currIdx][currMask] + edgeTime
				prevMatrix[edgeIdx][edgeMask] = item.node
				Time := costMatrix[currIdx][currMask] + edgeTime + currTime
				q = append(q, queueItem{edge.node, Time, edgeMask})
			}
		}
	}
	// fmt.Println("Done")
}

// reverse a slice of nodes
func reverse(list []*Node) []*Node {
	a := make([]*Node, len(list))
	copy(a, list)
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	return a
}
