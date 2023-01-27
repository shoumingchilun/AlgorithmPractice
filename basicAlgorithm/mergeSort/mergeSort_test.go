package mergeSort

import (
	"fmt"
	"testing"
)

func TestMerge(t *testing.T) {
	fmt.Println(MergeSort([]int{3, 4, 2, 1, 5, 7, 8, 4, 7}))
}
