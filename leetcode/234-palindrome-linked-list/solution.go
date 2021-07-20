package leetcode

/*
  除了栈、反转后半部分，还可以通过递归方式解决。
  这段代码展示了如何实现链表的*倒序遍历*。
*/
func isPalindrome(head *ListNode) bool {
	var traverse func(*ListNode) bool
	traverse = func(node *ListNode) bool {
		if node == nil {
			return true
		}
		if !traverse(node.Next) {
			return false
		}
		if head.Val != node.Val {
			return false
		}
		head = head.Next
		return true
	}
	return traverse(head)
}
