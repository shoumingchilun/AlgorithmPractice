package chorusFormation

import (
	"fmt"
	"testing"
)

func TestChorusFormation(t *testing.T) {
	index, number := chorusFormation([]int{186, 186, 150, 200, 160, 130, 197, 220})
	fmt.Println(index)
	fmt.Println(number)
}
