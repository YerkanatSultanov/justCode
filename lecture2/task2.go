package lecture2

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	a := list1
	b := list2
	ans := &ListNode{}
	res := ans

	if a == nil {
		return b
	} else if b == nil {
		return a
	}

	for a != nil && b != nil {
		if a.Val <= b.Val {
			ans.Next = a
			a = a.Next
		} else {
			ans.Next = b
			b = b.Next
		}
		ans = ans.Next
	}

	if a != nil {
		ans.Next = a
	} else if b != nil {
		ans.Next = b
	}

	return res.Next
}
