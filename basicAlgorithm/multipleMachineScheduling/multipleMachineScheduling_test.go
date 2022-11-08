package multipleMachineScheduling

import "testing"

func TestMultipleMachineScheduling(t *testing.T) {
	MultipleMachineScheduling([]int{1, 2, 3, 4, 5}, 5)
	println()
	MultipleMachineScheduling([]int{1, 2, 3, 4, 5}, 1)
	println()
	MultipleMachineScheduling([]int{1, 2, 3, 4, 5}, 3)
}
