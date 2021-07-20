package leetcode

/*
  看了官方题解的递归：
  我的递归虽然也是o(n)，但是没必要双返回值。因为原head其实指向了反转后链表的链表尾
*/
func reverseList3(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	reversed := reverseList3(head.Next) // reversed指向反转后链表的链表头，而head指向链表尾
	head.Next.Next = head
	head.Next = nil
	return reversed
}
