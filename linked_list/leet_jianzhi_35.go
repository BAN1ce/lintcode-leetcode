package linked_list

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// 剑指 Offer 35. 复杂链表的复制
// https://leetcode-cn.com/problems/fu-za-lian-biao-de-fu-zhi-lcof/
func copyRandomList(head *Node) *Node {

	m := make(map[*Node]int)
	tmp := head
	var i int
	for tmp != nil {
		m[tmp] = i
		i++
		tmp = tmp.Next
	}

	i = 0
	var result *Node

	var last *Node
	newNodes := make([]*Node,0)
	for head != nil {

		tmp := &Node{
			Val:    head.Val,
			Next:   head.Next,
			Random: head.Random,
		}
		if i == 0 {
			result = tmp
		} else {
			last.Next = tmp
		}
		i++
		last = tmp
		head = head.Next
		newNodes = append(newNodes,tmp)
	}

	tmp = result
	for tmp != nil{
		if tmp.Random != nil {
			index := m[tmp.Random]
			tmp.Random = newNodes[index]
		}
		tmp = tmp.Next
	}
	return result
}
