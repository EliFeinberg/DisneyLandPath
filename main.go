package main

import (
	"fmt"
	"time"

	"github.com/EliFeinberg/DisneyLandPath/graph"
)

func main() {
	start := time.Now()
	TSP := graph.BuildGraph("./CSV/RideTimesSpring.csv", "./CSV/walktime.csv", "./CSV/small_rides.csv")
	path, rideTime := TSP.Traversal(60)
	for i, node := range path {
		fmt.Println(node.Name, "ride #", i+1)
	}
	fmt.Println("Total time:", rideTime)
	fmt.Println("Time elapsed:", time.Since(start))

	start = time.Now()
	path2, time2 := TSP.TraversalGo(60)
	fmt.Println("Go Routines Implementation")
	for i, node := range path2 {
		fmt.Println(node.Name, "ride #", i+1)
	}
	fmt.Println("Total time:", time2)
	fmt.Println("Time elapsed:", time.Since(start), "\n")

	diff := false
	for i := range path {
		if path[i] != path2[i] {
			fmt.Println("Paths differ at", i+1)
			diff = true
		}
	}
	if rideTime != time2 {
		fmt.Println("Times differ\n")
		diff = true
	}
	if diff {
		fmt.Println("Results differ")
	}
}
