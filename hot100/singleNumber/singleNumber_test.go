package singleNumber

import (
	"fmt"
	"testing"
)

func Test_singleNumber(t *testing.T) {
	nums := []int{1, 3, 3, 5, 5, 4, 4}
	fmt.Println(singleNumber(nums))
}
