package restoreIpAddresses

import (
	"fmt"
	"testing"
)

func Test_restoreIpAddresses(t *testing.T) {
	result := restoreIpAddresses("101023")
	for _, s := range result {
		fmt.Print(s, " ")
	}
}
