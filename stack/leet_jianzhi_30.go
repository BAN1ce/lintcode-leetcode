package stack

import "container/list"

type MinStack struct {
	stack    *list.List
	minStack *list.List

}

/** initialize your data structure here. */
func Constructor() MinStack {

	return MinStack{
		stack: new(list.List),
		minStack:   new(list.List),
	}
}

func (this *MinStack) Push(x int) {

	if this.minStack.Len() == 0 {
		this.minStack.PushBack(x)
	} else {
		if x  < this.minStack.Back().Value.(int) {
			this.minStack.PushBack(x)
		}else {
			tmp := this.minStack.Back().Value.(int)
			this.minStack.PushBack(tmp)
		}
	}
	this.stack.PushBack(x)
}

func (this *MinStack) Pop() {
	this.stack.Remove(this.stack.Back())
	this.minStack.Remove(this.minStack.Back())
}

func (this *MinStack) Top() int {

	return this.stack.Back().Value.(int)
}

func (this *MinStack) Min() int {

	return  this.minStack.Back().Value.(int)
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */
