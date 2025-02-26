package letterCombinations

import (
	"fmt"
	"testing"
)

func Test_letterCombinations(t *testing.T) {
	result := letterCombinations2("23")
	for _, s := range result {
		fmt.Print(s, " ")
	}
}
