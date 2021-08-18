package linked_list

// 剑指 Offer 18. 删除链表的节点
// https://leetcode-cn.com/problems/shan-chu-lian-biao-de-jie-dian-lcof/
func deleteNode(head *ListNode, val int) *ListNode {

	var last *ListNode
	for tmp := head; tmp != nil; {
		if tmp.Val == val {
			if tmp == head {
				return tmp.Next
			}
			last.Next = tmp.Next
			return head
		}
		last = tmp
	}

	return head
}
