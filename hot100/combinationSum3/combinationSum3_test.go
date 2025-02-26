package combinationSum3

import (
	"fmt"
	"testing"
)

func Test_combinationSum3(t *testing.T) {
	result := combinationSum3(3, 9)
	for _, nums := range result {
		for _, num := range nums {
			fmt.Print(num, " ")
		}
		fmt.Println()
	}
}
