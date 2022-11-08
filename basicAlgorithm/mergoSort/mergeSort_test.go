package mergoSort

import (
	"fmt"
	"testing"
)

func TestMerge(t *testing.T) {
	fmt.Println(MergoSort([]int{3, 4, 2, 1, 5, 7, 8, 4, 7}))
}
