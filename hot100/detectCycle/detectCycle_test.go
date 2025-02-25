package detectCycle

import (
	"fmt"
	"testing"
)

func Test_detectCycle(t *testing.T) {
	n1 := ListNode{Val: 1, Next: nil}
	n2 := ListNode{Val: 2, Next: &n1}
	n3 := ListNode{Val: 3, Next: &n2}
	n4 := ListNode{Val: 4, Next: &n3}
	n5 := ListNode{Val: 5, Next: &n4}
	n6 := ListNode{Val: 6, Next: &n5}
	n7 := ListNode{Val: 7, Next: &n6}
	n8 := ListNode{Val: 8, Next: &n7}

	n1.Next = &n4

	fmt.Println(detectCycle(&n8).Val)
	fmt.Println(detectCycle2(&n8).Val)
}
