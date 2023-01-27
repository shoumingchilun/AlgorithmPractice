package heapSort

import (
	"fmt"
	"testing"
)

func TestHeapSort(t *testing.T) {
	fmt.Println(HeapSort([]int{3, 4, 2, 1, 5, 7, 8, 4, 7}))
}
