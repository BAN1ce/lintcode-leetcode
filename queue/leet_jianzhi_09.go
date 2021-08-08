package queue

import "container/list"

type CQueue struct {
	l1 *list.List
	l2 *list.List
}

func Constructor() CQueue {

	return CQueue{l1: new(list.List), l2: new(list.List)}
}

func (this *CQueue) AppendTail(value int) {

	this.l1.PushBack(value)

}

func (this *CQueue) DeleteHead() int {

	if this.l2.Len() != 0 {

		if v, ok := this.l2.Back().Value.(int); ok {

			this.l2.Remove(this.l2.Back())
			return v
		} else {
			return -1
		}
	}

	if this.l1.Len() > 0 {
		for this.l1.Len() > 0 {
			this.l2.PushBack(this.l1.Remove(this.l1.Back()))
		}
		if v, ok := this.l2.Back().Value.(int); ok {
			this.l2.Remove(this.l2.Back())
			return v
		} else {
			return -1
		}

	}
	return -1
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
