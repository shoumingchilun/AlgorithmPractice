package optimalSearchTree

import (
	"fmt"
	"testing"
)

func TestOptimalSearchTree(t *testing.T) { //期望为2.04
	i := OptimalSearchTree([]string{"A", "B", "C", "D", "E"}, []float64{0.04, 0.1, 0.02, 0.3, 0.02, 0.1, 0.05, 0.2, 0.06, 0.1, 0.01})
	fmt.Println(i)
}
