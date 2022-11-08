package quickSort

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	fmt.Println(QuickSort([]int{3, 4, 2, 1, 5, 7, 8, 4, 7}))
}
