package solution

/*
 * @lc app=leetcode.cn id=146 lang=golang
 *
 * [146] LRU 缓存机制
 */

// @lc code=start
type LRUCache struct {
	cache      map[int]*LinkedNode
	capacity   int
	size       int
	head, tail *LinkedNode
}

type LinkedNode struct {
	key, value int
	prev, next *LinkedNode
}

func initLinkedNode(key, value int) *LinkedNode {
	return &LinkedNode{key: key, value: value}
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		cache:    map[int]*LinkedNode{},
		capacity: capacity,
		head:     initLinkedNode(0, 0),
		tail:     initLinkedNode(0, 0),
	}
	l.head.next = l.tail
	l.tail.prev = l.head
	return l

}

func (this *LRUCache) Get(key int) int {
	if _, ok := this.cache[key]; !ok {
		return -1
	}
	node := this.cache[key]
	this.moveToHead(node)
	return node.value
}

func (this *LRUCache) removeTail() *LinkedNode {
	node := this.tail.prev
	this.removeNode(node)
	return node
}

func (this *LRUCache) removeNode(node *LinkedNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (this *LRUCache) addToHead(node *LinkedNode) {
	node.prev = this.head
	node.next = this.head.next
	this.head.next.prev = node
	this.head.next = node
}

func (this *LRUCache) moveToHead(node *LinkedNode) {
	this.removeNode(node)
	this.addToHead(node)
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.cache[key]; !ok {
		node := initLinkedNode(key, value)
		this.cache[key] = node
		this.addToHead(node)
		this.size++
		if this.size > this.capacity {
			removed := this.removeTail()
			delete(this.cache, removed.key)
			this.size--
		}
	} else {
		node := this.cache[key]
		this.moveToHead(node)
		this.cache[key].value = value
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end
