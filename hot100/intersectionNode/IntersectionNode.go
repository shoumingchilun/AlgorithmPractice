package intersectionNode

/*
https://leetcode.cn/problems/intersection-of-two-linked-lists/description/?envType=study-plan-v2&envId=top-100-liked
160. 相交链表
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	usedNodes := make(map[*ListNode]struct{})
	p1 := headA
	for p1 != nil {
		usedNodes[p1] = struct{}{}
		p1 = (*p1).Next
	}
	p2 := headB
	for p2 != nil {
		_, exists := usedNodes[p2]
		if exists {
			return p2
		}
		p2 = (*p2).Next
	}
	return nil
}

// getIntersectionNode2最优
func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	p1 := headA
	p2 := headB
	for {
		if p1 == p2 {
			return p1
		}
		if p1 == nil {
			p1 = headB
		} else {
			p1 = p1.Next
		}
		if p2 == nil {
			p2 = headA
		} else {
			p2 = p2.Next
		}
	}
}

func getIntersectionNode3(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	p1 := headA
	for p1 != nil {
		p1.Val = -p1.Val
		p1 = p1.Next
	}
	var ret *ListNode
	p2 := headB
	for p2 != nil {
		if p2.Val < 0 {
			ret = p2
			break
		}
		p2 = p2.Next
	}
	p1 = headA
	for p1 != nil {
		p1.Val = -p1.Val
		p1 = p1.Next
	}
	return ret
}
