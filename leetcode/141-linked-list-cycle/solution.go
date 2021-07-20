package leetcode

/*
  第一种方案就是记录下所有走过的路径，此时，时间o(n)，空间o(n)。
  为了达到需求，用o(1)空间，我考虑使用快慢指针。
*/
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	var (
		fast = head.Next
		slow = head
	)
	for fast != nil {
		if fast == slow {
			return true
		}
		slow = slow.Next
		if fast.Next == nil {
			break
		}
		fast = fast.Next.Next
	}
	return false
}
