package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
  由于链表长度未知，想要一遍扫描得知倒数第n个，唯一的办法是双指针。
  链表题有一个小技巧，伪头简化边界条件。
*/
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	fakeHead := ListNode{
		0,
		head,
	}
	var (
		slow = &fakeHead
		fast = slow
	)
	for ; n != 0; n-- {
		fast = fast.Next
	}
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}
	slow.Next = slow.Next.Next
	return fakeHead.Next
}
