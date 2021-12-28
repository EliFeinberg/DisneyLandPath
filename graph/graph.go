package graph

import (
	"strconv"

	"github.com/EliFeinberg/DisneyLandPath/utils"
)

var (
	_graphNodes []*Node
)

func buildGraph(rideCSV string, walkCSV string, idCSV string) {
	// parse CSV files
	rideRows := utils.ParseCSV(rideCSV)
	rideRows = rideRows[1:] // remove header row
	walkRows := utils.ParseCSV(walkCSV)
	walkRows = walkRows[1:] // remove header row
	idRows := utils.ParseCSV(idCSV)
	idRows = idRows[1:] // remove header row

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

		Node := &Node{
			_idx:       i,
			_name:      ride[2],
			_waitTimes: make([]int, len(rideRows)),
			_edges:     make([]*Edge, len(rideRows)),
		}
		_graphNodes = append(_graphNodes, Node)

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

			_graphNodes[rideA]._edges = append(_graphNodes[rideA]._edges, &Edge{
				time: walkTime,
				node: _graphNodes[rideB],
			})
			_graphNodes[rideB]._edges = append(_graphNodes[rideB]._edges, &Edge{
				time: walkTime,
				node: _graphNodes[rideA],
			})
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
			waitTime := utils.calculateIdx(rideTime[1])

			if waitTime < 16 {
				_graphNodes[ride]._waitTimes[waitTime] = waitTime
			}
		}

	}

}
