package leecode

func init() {
	type ListNode struct {
		Val  int
		Next *ListNode
	}
	addTwoNumbers := func(l1 *ListNode, l2 *ListNode) *ListNode {
		var res *ListNode
		for p, n := &res, 0; l1 != nil || l2 != nil || n != 0; {
			node := &ListNode{Val: n}
			*p, p = node, &node.Next
			if l1 != nil {
				node.Val += l1.Val
				l1 = l1.Next
			}
			if l2 != nil {
				node.Val += l2.Val
				l2 = l2.Next
			}
			if n = 0; node.Val > 9 {
				node.Val -= 10
				n = 1
			}
		}
		return res
	}
	Cmd(leecode{help: "两数相加",
		link: "https://leetcode.cn/problems/add-two-numbers/",
		test: `
					[2,4,3]         [5,6,4]   [7,0,8]
					[0]             [0]       [0]
					[9,9,9,9,9,9,9] [9,9,9,9] [8,9,9,9,0,0,0,1]
		`,
		hand: func(n1 []int, n2 []int) (res []int) {
			var l1, l2 *ListNode
			p := &l1
			for _, n := range n1 {
				node := &ListNode{Val: n}
				*p, p = node, &node.Next
			}
			p = &l2
			for _, n := range n2 {
				node := &ListNode{Val: n}
				*p, p = node, &node.Next
			}
			for l := addTwoNumbers(l1, l2); l != nil; {
				res, l = append(res, l.Val), l.Next
			}
			return
		},
	})
}
