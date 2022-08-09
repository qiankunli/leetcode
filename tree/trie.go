package tree

type Trie struct {
	Val      uint8
	isEnd    bool
	Children map[uint8]*Trie
}

func Constructor() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) {
	if len(word) == 0 {
		return
	}
	if len(word) == 1 {
		c := word[0]
		if len(this.Children) == 0 {
			this.Children = make(map[uint8]*Trie, 0)
			this.Children[c] = &Trie{Val: c, isEnd: true}
			return
		}
		if _, ok := this.Children[c]; ok {
			return
		}
		this.Children[c] = &Trie{Val: c, isEnd: true}
		return
	}
	c := word[0]
	var trie *Trie
	if len(this.Children) == 0 {
		this.Children = make(map[uint8]*Trie)
	}
	var ok bool
	if trie, ok = this.Children[c]; !ok {
		trie = &Trie{Val: c}
		this.Children[c] = trie
	}
	trie.Insert(word[1:])
}

// 迭代实现
func (this *Trie) Insert2(word string) {
	if len(word) == 0 {
		return
	}
	cur := this
	for _, ch := range word {
		if len(cur.Children) == 0 {
			cur.Children = make(map[uint8]*Trie)
		}
		var ok bool
		var trie *Trie
		if trie, ok = cur.Children[uint8(ch)]; !ok {
			trie = &Trie{Val: uint8(ch)}
			cur.Children[uint8(ch)] = trie
		}
		cur = trie
	}
	cur.isEnd = true
}

func (this *Trie) Search(word string) bool {
	if len(word) == 0 {
		return false
	}
	if len(word) == 1 {
		c := word[0]
		trie, ok := this.Children[c]
		return ok && trie.isEnd
	}
	c := word[0]
	trie, ok := this.Children[c]
	if !ok {
		return false
	}
	return trie.Search(word[1:])
}

func (this *Trie) StartsWith(prefix string) bool {
	if len(prefix) == 0 {
		return false
	}
	if len(prefix) == 1 {
		c := prefix[0]
		_, ok := this.Children[c]
		return ok
	}
	c := prefix[0]
	trie, ok := this.Children[c]
	if !ok {
		return false
	}
	return trie.StartsWith(prefix[1:])
}
