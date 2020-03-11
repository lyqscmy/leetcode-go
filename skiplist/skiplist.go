package leetcode

// Should be enough for 2^64 elements
const MAXLEVEL = 64

type skiplistLevel struct {
	next *skiplistNode
	span int
}

type skiplistNode struct {
	key   int
	pre   *skiplistNode
	level []skiplistLevel
}

type skiplist struct {
	header, tail *skiplistNode
	length       int
	level        int
}

func NewSkiplist() *skiplist {
	sl := &skiplist{level: 1, length: 0}
	sl.header = NewSkiplistNode(MAXLEVEL, -1)
	for i := 0; i < MAXLEVEL; i++ {
		sl.header.level[i].next = nil
		sl.header.level[i].span = 0
	}
	sl.header.pre = nil
	sl.tail = nil
	return sl
}

func NewSkiplistNode(level int, key int) *skiplistNode {
	return &skiplistNode{key: key, level: make([]skiplistLevel, level)}
}

func (sl *skiplist) get(key int) *skiplistNode {
	x := sl.header
	for i := sl.level - 1; i >= 0; i-- {
		for next := x.level[i].next; next != nil && next.key < key; {
			x = x.level[i].next
		}
		if x.key == key {
			return x
		}
	}
	return x
}
