package singleSourceShortestPaths

import (
	"fmt"
	"math"
	"testing"
	"workSpace/basicDataStructure"
)

func TestDijkstra(t *testing.T) {
	var am = basicDataStructure.CreateAdjacencyMatrix(5)
	am.Distances[0][1] = 10
	am.Distances[0][2] = math.MaxFloat64
	am.Distances[0][3] = 30
	am.Distances[0][4] = 100
	am.Distances[1][0] = math.MaxFloat64
	am.Distances[1][2] = 50
	am.Distances[1][3] = math.MaxFloat64
	am.Distances[1][4] = math.MaxFloat64
	am.Distances[2][0] = math.MaxFloat64
	am.Distances[2][1] = math.MaxFloat64
	am.Distances[2][3] = math.MaxFloat64
	am.Distances[2][4] = 10
	am.Distances[3][0] = math.MaxFloat64
	am.Distances[3][1] = math.MaxFloat64
	am.Distances[3][2] = 20
	am.Distances[3][4] = 60
	am.Distances[4][0] = math.MaxFloat64
	am.Distances[4][1] = math.MaxFloat64
	am.Distances[4][2] = math.MaxFloat64
	am.Distances[4][3] = math.MaxFloat64
	fmt.Println(Dijkstra(am, 0))
}
