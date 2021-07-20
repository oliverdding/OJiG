package leetcode

/*
  既然可以递归，那么也可以迭代。
  维护一个结果列表，从旧列表头一个个拆节点放入结果列表头，实现倒序。
*/
func reverseList(head *ListNode) *ListNode {
	var (
		resultHead = ListNode{
			0,
			nil,
		} // fake head
		tmp  = head
		next *ListNode
	)
	for tmp != nil {
		// mark old list
		next = tmp.Next
		// extract `tmp` to resultHead -> tmp -> ...
		tmp.Next = resultHead.Next
		resultHead.Next = tmp
		// next round
		tmp = next
	}
	return resultHead.Next
}
