package reverseList

/*
https://leetcode.cn/problems/reverse-linked-list/?envType=study-plan-v2&envId=top-100-liked
206. 反转链表
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p1, p2, p3 := head, head.Next, head.Next.Next
	p1.Next = nil

	for p2 != nil {
		p2.Next = p1
		if p3 != nil {
			p1, p2, p3 = p2, p3, p3.Next
		} else {
			p1, p2 = p2, p3
		}
	}
	return p1
}
