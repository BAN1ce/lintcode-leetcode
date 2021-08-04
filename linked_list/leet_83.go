package linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

// 删除排序链表中的重复元素
// https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list/
func deleteDuplicates(head *ListNode) *ListNode {

	tmp := head
	for ; tmp != nil; {

		innerTmp := tmp
		if tmp.Next != nil {
			for ; innerTmp.Next.Val == tmp.Val; {
				innerTmp = innerTmp.Next
				if innerTmp.Next == nil {
					break
				}
			}
			tmp.Next = innerTmp.Next
			tmp = innerTmp.Next
		} else {
			break
		}
	}
	return head
}
