package combinationSum

import (
	"testing"
)

func Test_combinationSum(t *testing.T) {
	ret := combinationSum([]int{7, 3, 2}, 18)
	for _, nums := range ret {
		for _, num := range nums {
			print(num, " ")
		}
		println()
	}
}
