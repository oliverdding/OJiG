package leetcode

/*
  看完题解才想到还可以递归，不知道为什么对于链表题目我喜欢迭代，对于数组我容易考虑递归。

  mergeTwoList =
  l1 -> mergeTwoList(l1.Next, l2) if l1.Val < l2.Val
  l2 -> mergeTwoList(l1, l2.Next) else

  边界也比较容易判断，nil的时候直接返回
*/
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoLists(l1, l2.Next)
		return l2
	}
}
