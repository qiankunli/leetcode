package design

import (
	"fmt"
	"testing"
)

func TestLRUCache_Get(t *testing.T) {
	//cache := Constructor(2)
	//cache.Put(1, 1)
	//cache.Put(2, 2)
	//cache.Put(3, 3)
	//fmt.Println(cache.Get(1))

	//cache := Constructor(1)
	//cache.Put(2, 1)
	//fmt.Println(cache.Get(2))

	//cache := Constructor(2)
	//cache.Get(2)
	//cache.Put(2, 6)
	//cache.Get(1)
	//cache.Put(1, 5)
	//cache.Put(1, 2)
	//cache.Get(1)
	//cache.Get(2)

	//cache := Constructor(2)
	//cache.Put(2, 1)
	//cache.Put(2, 2)
	//fmt.Println(cache.Get(2))
	//cache.Put(1, 1)
	//cache.Put(4, 1)
	//fmt.Println(cache.Get(2))

	cache := Constructor(2)
	cache.Put(2, 1)
	cache.Put(1, 1)
	cache.Put(2, 3)
	cache.Put(4, 1)
	fmt.Println(cache.Get(1))
	fmt.Println(cache.Get(2))
}
