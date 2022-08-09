package tree

import (
	"fmt"
	"testing"
)

func TestInsert(t *testing.T) {
	trie := Constructor()
	trie.Insert("hello")
	trie.Insert("world")
	fmt.Println(trie.StartsWith("he"))
	fmt.Println(trie.StartsWith("hw"))
	fmt.Println(trie.Search("world"))
}
