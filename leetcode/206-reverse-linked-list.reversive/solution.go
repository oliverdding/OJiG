package leetcode

/*
  看了官方题解的迭代：
  我的方案是另起一个链表维护反转后的链表。
  官方方案中，反转的不是链表的节点，而是链表的指向！
  ---
  我：
  1 -> 2 -> 3 -> 4 -> nil
    ==>
  4 -> 3 -> 2 -> 1 -> nil
  ---
  官方：
  1 -> 2 -> 3 -> 4 -> nil
    ==>
  nil <- 1 <- 2 <- 3 <- 4
  ---
  注意区别很大
*/
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var (
		pre *ListNode = nil
		cur           = head
		nex           = head.Next
	)
	for nex != nil {
		cur.Next = pre
		pre = cur
		cur = nex
		nex = nex.Next
	}
	cur.Next = pre // 最后还要接一下
	return cur
}
