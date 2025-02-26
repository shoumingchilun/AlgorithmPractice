package partitioningPalindromic

import (
	"fmt"
	"testing"
)

func Test_partitioningPalindromic(t *testing.T) {
	result := partitioningPalindromic("aab")
	for _, strings := range result {
		for _, s := range strings {
			fmt.Print(s, " ")
		}
		fmt.Println()
	}
}
