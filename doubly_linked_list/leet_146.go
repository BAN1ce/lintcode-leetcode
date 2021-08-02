package doubly_linked_list

import (
	"container/list"
)

// 146. LRU 缓存机制
// https://leetcode-cn.com/problems/lru-cache/

type LRUCache struct {
	list     *list.List
	capacity int
	m        map[int]*list.Element
}
type n struct {
	key   int
	value int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{list: list.New(), capacity: capacity, m: make(map[int]*list.Element)}
}

func (this *LRUCache) Get(key int) int {

	if v, ok := this.m[key]; ok {
		if tmp, ol := v.Value.(*n); ol {
			this.list.MoveAfter(v, this.list.Back())
			return tmp.value
		}
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {

	if v, ok := this.m[key]; ok {
		if tmp, ol := v.Value.(*n); ol {
			tmp.value = value
			this.list.MoveAfter(v, this.list.Back())
		}
		return
	}

	if this.list.Len() == this.capacity {
		p := this.list.Front()
		if p != nil {
			this.list.Remove(p)
			if tmp, ok := p.Value.(*n); ok {
				delete(this.m, tmp.key)
			}
		}
	}
	last := this.list.PushBack(&n{
		key:   key,
		value: value,
	})
	this.m[key] = last

}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
