package design

import (
	"leetcode/linkedlist"
)

// 460
type LFUCache struct {
	capacity int
	len      int
	minFeq   int
	hash     map[int]*LFUCacheValue
	// key = freq
	freq map[int]*linkedlist.DoubleLinkedList
}
type LFUCacheValue struct {
	value int
	freq  int
	node  *linkedlist.Node
}

func LFUCacheConstructor(capacity int) LFUCache {
	return LFUCache{
		capacity: capacity,
		hash:     make(map[int]*LFUCacheValue, capacity),
		freq:     make(map[int]*linkedlist.DoubleLinkedList, capacity),
	}
}

func (this *LFUCache) Get(key int) int {
	v, ok := this.hash[key]
	if ok {
		this.addFreq(v.freq, v.node)
		v.freq++
		return v.value
	}
	return -1
}

func (this *LFUCache) Put(key int, value int) {
	if this.capacity == 0 {
		return
	}
	v, ok := this.hash[key]
	if ok {
		this.addFreq(v.freq, v.node)
		v.freq++
		v.value = value
		return
	}
	if this.len >= this.capacity {
		// 找到频率最小的元素
		node := this.freq[this.minFeq].Tail()
		this.deleteFreq(this.minFeq, node)
		delete(this.hash, node.Val)
		this.len--
	}
	v = &LFUCacheValue{
		value: value,
		freq:  1,
		node: &linkedlist.Node{
			Val: key,
		},
	}
	this.hash[key] = v
	if _, ok := this.freq[1]; !ok {
		this.freq[1] = linkedlist.NewDoubleLinkedList()
	}
	this.freq[1].AppendToFirst(v.node)
	this.minFeq = 1 //  插入的为1， minFeq 最小也是1
	this.len++
}
func (this *LFUCache) addFreq(freq int, node *linkedlist.Node) {
	this.freq[freq].Delete(node)
	if this.freq[freq].IsEmpty() {
		delete(this.freq, freq)
		if this.minFeq == freq {
			this.minFeq++
		}
	}
	if _, ok := this.freq[freq+1]; !ok {
		this.freq[freq+1] = linkedlist.NewDoubleLinkedList()
	}
	this.freq[freq+1].AppendToFirst(node)
}
func (this *LFUCache) deleteFreq(freq int, node *linkedlist.Node) {
	this.freq[freq].Delete(node)
	if this.freq[freq].IsEmpty() {
		delete(this.freq, freq)
	}
}
