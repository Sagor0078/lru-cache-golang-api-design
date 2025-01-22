package main

type Node struct {
	key  int
	val  int
	prev *Node
	next *Node
}

type LRUCache struct {
	capacity int
	size     int
	cache    map[int]*Node
	head     *Node
	tail     *Node
}

func newNode(key, val int) *Node {
	return &Node{key: key, val: val}
}

func Constructor(capacity int) LRUCache {
	head := newNode(-1, -1)
	tail := newNode(-1, -1)

	head.next = tail
	tail.prev = head

	return LRUCache{
		capacity: capacity,
		size:     0,
		cache:    make(map[int]*Node),
		head:     head,
		tail:     tail,
	}
}

// (Node = (x, y))
func (this *LRUCache) removeNode(node *Node) {
	prev := node.prev
	next := node.next
	prev.next = next
	next.prev = prev
}

func (this *LRUCache) addNode(node *Node) {
	next := this.head.next
	this.head.next = node
	node.prev = this.head
	node.next = next
	next.prev = node
}

func (this *LRUCache) Get(key int) int {
	if node, exists := this.cache[key]; exists {
		this.removeNode(node)
		this.addNode(node)
		return node.val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, exists := this.cache[key]; exists {
		this.removeNode(node)
		node.val = value
		this.addNode(node)
		this.cache[key] = node
	} else {
		if this.size == this.capacity {
			lru := this.tail.prev
			this.removeNode(lru)
			delete(this.cache, lru.key)
			this.size--
		}
		newNode := newNode(key, value)
		this.addNode(newNode)
		this.cache[key] = newNode
		this.size++
	}
}
