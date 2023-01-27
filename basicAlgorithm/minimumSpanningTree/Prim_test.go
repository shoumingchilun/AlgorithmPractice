package minimumSpanningTree

import (
	"fmt"
	"math"
	"testing"
	"workSpace/basicDataStructure"
)

func TestPrim(t *testing.T) {
	var am = basicDataStructure.CreateAdjacencyMatrix(6)
	am.Distances[0][1] = 6
	am.Distances[0][2] = 1
	am.Distances[0][3] = 5
	am.Distances[0][4] = math.MaxFloat64
	am.Distances[0][5] = math.MaxFloat64

	am.Distances[1][0] = 6
	am.Distances[1][2] = 5
	am.Distances[1][3] = math.MaxFloat64
	am.Distances[1][4] = 3
	am.Distances[1][5] = math.MaxFloat64

	am.Distances[2][0] = 1
	am.Distances[2][1] = 5
	am.Distances[2][3] = 5
	am.Distances[2][4] = 6
	am.Distances[2][5] = 4

	am.Distances[3][0] = 5
	am.Distances[3][1] = math.MaxFloat64
	am.Distances[3][2] = 5
	am.Distances[3][4] = math.MaxFloat64
	am.Distances[3][5] = 2

	am.Distances[4][0] = math.MaxFloat64
	am.Distances[4][1] = 3
	am.Distances[4][2] = 6
	am.Distances[4][3] = math.MaxFloat64
	am.Distances[4][5] = 6

	am.Distances[5][0] = math.MaxFloat64
	am.Distances[5][1] = math.MaxFloat64
	am.Distances[5][2] = 4
	am.Distances[5][3] = 2
	am.Distances[5][4] = 6
	fmt.Println(Prim(am))
}
