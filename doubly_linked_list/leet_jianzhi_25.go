package doubly_linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

// 剑指 Offer 25. 合并两个排序的链表
// https://leetcode-cn.com/problems/he-bing-liang-ge-pai-xu-de-lian-biao-lcof/
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

	head := new(ListNode)

	tmp := head
	for ; l1 != nil && l2 != nil; {

		if l1.Val > l2.Val {
			tmp.Next = l2
			l2 = l2.Next
		} else {
			tmp.Next = l1
			l1 = l1.Next
		}
		tmp = tmp.Next
	}
	if l2 == nil{
		 tmp.Next = l1
	}else {
		tmp.Next = l2
	}
	return head.Next
}
