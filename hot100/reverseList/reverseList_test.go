package reverseList

import (
	"fmt"
	"testing"
)

func Test_reverseList(t *testing.T) {
	n1 := ListNode{Val: 1}
	n2 := ListNode{Val: 2, Next: &n1}
	n3 := ListNode{Val: 3, Next: &n2}
	ret := reverseList(&n3)
	fmt.Printf("%d  %d  %d", ret.Val, ret.Next.Val, ret.Next.Next.Val)
}
