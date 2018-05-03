package main

type LRUCache struct {
	cap, size   int
	data        map[int]int
	nodes       map[int]*node
	bottom, top *node
}

type node struct {
	key       int
	pre, next *node
}

func Constructor(capacity int) LRUCache {
	return LRUCache{capacity, 0, make(map[int]int), make(map[int]*node), nil, nil}
}

func (this *LRUCache) Get(key int) int {
	v, exist := this.data[key]
	if !exist {
		return -1
	}
	// update stack
	updateStackTop(this, key)
	return v
}

func (this *LRUCache) Put(key int, value int) {
	_, exist := this.data[key]
	this.data[key] = value
	// update stack
	if exist {
		// update stack
		updateStackTop(this, key)
	} else {
		newNode := &node{key, this.top, nil}
		if this.size > 0 {
			this.top.next = newNode
		} else {
			this.bottom = newNode
		}
		this.top = newNode
		this.nodes[key] = newNode
		if this.size < this.cap {
			this.size += 1
		} else {
			oldest := this.bottom.key
			delete(this.data, oldest)
			delete(this.nodes, oldest)
			this.bottom.next.pre = nil
			this.bottom = this.bottom.next
		}
	}
}

func updateStackTop(this *LRUCache, key int) {
	node := this.nodes[key]
	if node != this.top {
		pre, next := node.pre, node.next
		next.pre = pre
		if pre != nil {
			pre.next = next
		} else {
			this.bottom = next
		}
		node.pre = this.top
		this.top.next = node
		node.next = nil
		this.top = node
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

func main() {
}
