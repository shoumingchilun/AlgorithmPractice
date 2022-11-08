package roundRobinSchedule

import (
	"fmt"
	"testing"
)

func TestRoundRobinSchedule(t *testing.T) {
	fmt.Println(RoundRobinSchedule(4))
}

func TestSimpleRoundRobinSchedule(t *testing.T) {
	n := 4
	var a = SimpleRoundRobinSchedule(n)
	for i := 0; i < 1<<n; i++ {
		fmt.Println(a[i])
	}
}
