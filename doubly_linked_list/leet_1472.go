package doubly_linked_list

type node struct {
	pre   *node
	next  *node
	value string
}

// 1472. 设计浏览器历史记录
// https://leetcode-cn.com/problems/design-browser-history/
type BrowserHistory struct {
	currentNode *node
}

func Constructor1(homepage string) BrowserHistory {

	return BrowserHistory{
		currentNode: &node{
			pre:   nil,
			next:  nil,
			value: homepage,
		},
	}
}

func (this *BrowserHistory) Visit(url string) {

	this.currentNode.next = &node{
		pre:   this.currentNode,
		next:  nil,
		value: url,
	}
	this.currentNode = this.currentNode.next
}

func (this *BrowserHistory) Back(steps int) string {

	for i := 0; i < steps; i++ {
		tmp := this.currentNode.pre
		if tmp != nil {
			tmp = this.currentNode.pre
			this.currentNode = tmp
		} else {
			return this.currentNode.value
		}
	}
	return this.currentNode.value
}

func (this *BrowserHistory) Forward(steps int) string {

	for i := 0; i < steps; i++ {
		tmp := this.currentNode.next
		if tmp != nil {
			tmp = this.currentNode.next
			this.currentNode = tmp
		} else {
			return this.currentNode.value
		}
	}
	return this.currentNode.value
}

/**
 * Your BrowserHistory object will be instantiated and called as such:
 * obj := Constructor(homepage);
 * obj.Visit(url);
 * param_2 := obj.Back(steps);
 * param_3 := obj.Forward(steps);
 */
