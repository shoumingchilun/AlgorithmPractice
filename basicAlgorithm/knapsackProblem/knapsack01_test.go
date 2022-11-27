package knapsackProblem

import (
	"fmt"
	"testing"
)

func TestKnapsack01(t *testing.T) {
	selected, value := knapsack01([]int{2, 2, 6, 5, 4}, []int{6, 3, 5, 4, 6}, 10)
	fmt.Println(selected)
	fmt.Println(value)
}
