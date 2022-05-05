package main

import "fmt"

type DLLNode struct {
	key, value int
	next, prev *DLLNode
}

// constructor for DLL
func NewDLLNode(key, value int) *DLLNode {
	return &DLLNode{key: key, value: value}
}

type LRUCache struct {
	capacity int
	count    int
	head     *DLLNode
	tail     *DLLNode
	mp       map[int]*DLLNode
}

// constructor for LRUCache
func NewLRUCache(cap int) *LRUCache {
	// sentinel head and tail nodes to avoid handling corner cases
	head := NewDLLNode(0, 0)
	tail := NewDLLNode(0, 0)
	head.next = tail
	tail.prev = head
	return &LRUCache{
		capacity: cap,
		count:    0,
		head:     head,
		tail:     tail,
		mp:       map[int]*DLLNode{},
	}
}

func main() {

	cache := NewLRUCache(2)
	cache.Set(1, 2)
	fmt.Println(cache.Get(1))

	cache = NewLRUCache(2)
	cache.Set(1, 2)
	cache.Set(2, 3)
	cache.Set(1, 5)
	cache.Set(4, 5)
	cache.Set(6, 7)
	fmt.Println(cache.Get(4))
	cache.Set(1, 2)
	fmt.Println(cache.Get(3))

}

//O(1)
func (cache *LRUCache) Get(key int) int {
	// check if key exists in doubly linked list
	var res int = -1
	node := cache.mp[key]

	if node != nil {
		res = node.value
		cache.Delete(node)
		cache.AddToHead(node)
	}
	return res
}

//O(1)
func (cache *LRUCache) Set(key, value int) {
	node := cache.mp[key]

	if node != nil {
		node.value = value
		cache.Delete(node)
		cache.AddToHead(node)
	} else {
		node := NewDLLNode(key, value)
		cache.mp[key] = node

		if cache.count < cache.capacity {
			cache.count++
			cache.AddToHead(node)
		} else {
			// delete last node for DLL
			delete(cache.mp, cache.tail.prev.key)
			cache.Delete(cache.tail.prev)
			cache.AddToHead(node)
		}
	}
}

//O(1)
func (cache *LRUCache) Delete(node *DLLNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

//O(1)
func (cache *LRUCache) AddToHead(node *DLLNode) {
	node.next = cache.head.next
	node.prev = cache.head
	cache.head.next.prev = node
	cache.head.next = node
}
