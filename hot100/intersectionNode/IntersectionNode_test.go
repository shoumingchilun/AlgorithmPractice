package intersectionNode

import (
	"fmt"
	"testing"
)

func Test_getIntersectionNode(t *testing.T) {
	n0 := ListNode{Val: 0, Next: nil}
	n1 := ListNode{Val: 1, Next: nil}
	n2 := ListNode{Val: 2, Next: nil}
	n3 := ListNode{Val: 3, Next: nil}
	n4 := ListNode{Val: 4, Next: nil}
	n5 := ListNode{Val: 5, Next: nil}
	n6 := ListNode{Val: 6, Next: nil}
	n7 := ListNode{Val: 7, Next: nil}

	n0.Next = &n1
	n1.Next = &n2
	n2.Next = &n3
	n3.Next = &n4

	n5.Next = &n6
	n6.Next = &n7
	n7.Next = &n3

	fmt.Print(getIntersectionNode(&n0, &n5).Val)
	fmt.Print(getIntersectionNode2(&n0, &n5).Val)
	fmt.Print(getIntersectionNode3(&n0, &n5).Val)
}
