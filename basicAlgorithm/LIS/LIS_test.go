package LIS

import (
	"fmt"
	"testing"
)

func TestLIS(t *testing.T) {
	minList, length := LIS([]int{5, 1, 3, 2, 4, 6, 5, 7})
	fmt.Println(minList, "长度：", length)
}
