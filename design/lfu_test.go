package design

import (
	"fmt"
	"testing"
)

func TestLFUCache_Get(t *testing.T) {
	cache := LFUCacheConstructor(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	fmt.Println(cache.Get(1))
	cache.Put(3, 3)
	fmt.Println(cache.Get(2))
	fmt.Println(cache.Get(3))
	cache.Put(4, 4)
	fmt.Println(cache.Get(1))
	fmt.Println(cache.Get(3))
	fmt.Println(cache.Get(4))

	//cache := LFUCacheConstructor(0)
	//cache.Put(0, 0)
	//fmt.Println(cache.Get(0))

}
