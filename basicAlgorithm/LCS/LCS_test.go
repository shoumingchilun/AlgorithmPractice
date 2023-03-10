package LCS

import (
	"fmt"
	"testing"
)

func TestLCS(t *testing.T) {
	fmt.Println(LCS([]string{"A", "B", "C", "B", "D", "A", "B"}, []string{"B", "D", "C", "A", "B", "A"}))
}
