package linked_list

// 剑指 Offer 24. 反转链表
// https://leetcode-cn.com/problems/fan-zhuan-lian-biao-lcof/
func reverseList(head *ListNode) *ListNode {

	reverse := reversePrint(head)

	tmp := head
	i := 0
	for tmp != nil {
		tmp.Val = reverse[i]
		tmp = tmp.Next
		i++
	}
	return head
}
