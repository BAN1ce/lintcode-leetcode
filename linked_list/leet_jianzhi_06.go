package linked_list


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 剑指 Offer 06. 从尾到头打印链表
// https://leetcode-cn.com/problems/cong-wei-dao-tou-da-yin-lian-biao-lcof/
func reversePrint(head *ListNode) []int {


	result := make([]int,0)
	for head != nil{
		tmp := make([]int,0)
		tmp = append(tmp, head.Val)
		result = append(tmp,result...)
		head = head.Next
	}

	return result

}
