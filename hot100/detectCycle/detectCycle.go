package detectCycle

/*
https://leetcode.cn/problems/linked-list-cycle-ii?envType=study-plan-v2&envId=top-100-liked
142. 环形链表 II
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	existingNode := make(map[*ListNode]struct{})
	p1 := head
	for p1 != nil {
		_, exist := existingNode[p1]
		if exist {
			return p1
		} else {
			existingNode[p1] = struct{}{}
		}
		p1 = p1.Next
	}
	return nil
}
func detectCycle2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	p1, p2 := head, head
	for p1 != nil && p2 != nil {
		p1 = p1.Next
		p2 = p2.Next
		if p2 != nil {
			p2 = p2.Next
		} else {
			return nil
		}
		if p1 == p2 {
			p3 := head
			for p2 != p3 {
				p2 = p2.Next
				p3 = p3.Next
			}
			return p2
		}
	}
	return nil
}
