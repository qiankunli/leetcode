package design

import "leetcode/linkedlist"

// 146
type LRUCache struct {
	capacity int
	len      int
	hash     map[int]*CacheValue
	list     *linkedlist.DoubleLinkedList
}
type CacheValue struct {
	value int
	node  *linkedlist.Node
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		hash:     make(map[int]*CacheValue, capacity),
		list:     linkedlist.NewDoubleLinkedList(),
	}
}

func (this *LRUCache) Get(key int) int {
	v, ok := this.hash[key]
	if !ok {
		return -1
	}
	this.list.MoveToFirst(v.node)
	return v.value
}

func (this *LRUCache) Put(key int, value int) {
	if v, ok := this.hash[key]; ok {
		v.value = value
		this.list.MoveToFirst(v.node)
		return
	}
	if this.len >= this.capacity {
		tail := this.list.Tail()
		delete(this.hash, tail.Val)
		this.list.Delete(tail)
		this.len--
	}
	node := &linkedlist.Node{
		Val: key, // node 实际存的是key，以便于删除尾元素时定位到map的key，不存value 是避免key覆盖时需要操作list
	}
	v := &CacheValue{
		value: value,
		node:  node,
	}
	this.hash[key] = v
	this.list.AppendToFirst(node)
	this.len++
}
