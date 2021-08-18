package linked_list

// 剑指 Offer 22. 链表中倒数第k个节点
// https://leetcode-cn.com/problems/lian-biao-zhong-dao-shu-di-kge-jie-dian-lcof/
func getKthFromEnd(head *ListNode, k int) *ListNode {

	list := make([]*ListNode,0)

	for tmp:=head;tmp!=nil;tmp = tmp.Next{
		list = append(list,tmp)
	}

	return list[len(list)-k]
}
