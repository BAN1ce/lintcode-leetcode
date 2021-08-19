package doubly_linked_list



// 剑指 Offer 52. 两个链表的第一个公共节点
// https://leetcode-cn.com/problems/liang-ge-lian-biao-de-di-yi-ge-gong-gong-jie-dian-lcof/
func getIntersectionNode(headA, headB *ListNode) *ListNode {

	for ; headB!=nil;{
		tmp := headA
		for ;tmp!=nil;tmp= tmp.Next{
			if tmp == headB {
				return tmp
			}
		}
		headB = headB.Next
	}

	return nil
}