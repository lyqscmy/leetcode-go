package solution

import "container/list"

type LRUCache struct {
	capacity int
	size     int
	evict    *list.List
	items    map[int]*list.Element
}

type Entry struct {
	key int
	val int
}

func NewLRUCache(capacity int) LRUCache {
	return LRUCache{capacity: capacity, evict: list.New(), items: make(map[int]*list.Element, capacity)}
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.items[key]
	if !ok {
		return -1
	}
	this.evict.MoveToFront(node)
	return node.Value.(*Entry).val
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.items[key]; ok {
		e := node.Value.(*Entry)
		e.val = value
		this.evict.MoveToFront(node)
		return
	}

	if this.size+1 > this.capacity {
		value := this.evict.Remove(this.evict.Back())
		delete(this.items, value.(*Entry).key)
		this.size--
	}
	node := this.evict.PushFront(&Entry{key, value})
	this.items[key] = node
	this.size++
}
