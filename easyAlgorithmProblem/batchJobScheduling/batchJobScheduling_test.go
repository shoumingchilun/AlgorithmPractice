package batchJobScheduling

import (
	"fmt"
	"testing"
)

func TestBatchJobScheduling(t *testing.T) {
	bTime, order := BatchJobScheduling([]int{2, 3, 2}, []int{1, 1, 3})
	fmt.Println(bTime, order)
}
