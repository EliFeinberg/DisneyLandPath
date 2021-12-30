package test

import (
	"testing"

	"github.com/EliFeinberg/DisneyLandPath/graph"
	"github.com/stretchr/testify/assert"
)

func getExpectedPath() []string {
	return []string{"Astro Orbitor", "Big Thunder Mountain Railroad", "Space Mountain"}
}

func testGraphTraversalSmall(t *testing.T) {
	TSP := graph.BuildGraph("./CSV/TEST_RideTimes.csv", "./CSV/TEST_WalkTimes.csv", "./CSV/TEST_rides.csv")
	path, rideTime := TSP.Traversal(60)
	t.Log("Total time:", rideTime)
	for i, node := range path {
		assert.Equal(t, getExpectedPath()[i], node.Name)
	}
	assert.Equal(t, rideTime, 32)
}

func testGraphTraversalGoSmall(t *testing.T) {
	TSP := graph.BuildGraph("./CSV/TEST_RideTimes.csv", "./CSV/TEST_WalkTimes.csv", "./CSV/TEST_rides.csv")
	path, rideTime := TSP.TraversalGo(60)
	t.Log("Total time:", rideTime)

	for i, node := range path {
		assert.Equal(t, getExpectedPath()[i], node.Name)
	}
	assert.Equal(t, rideTime, 32)
}

func TestGraph(t *testing.T) {
	t.Log("Testing Graph Traversal")
	t.Run("Testing Regular Traversal", testGraphTraversalSmall)
	t.Run("Testing Go Routines Traversal", testGraphTraversalGoSmall)
}
