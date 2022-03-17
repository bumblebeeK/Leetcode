package solution

import "fmt"

/*
 * @lc app=leetcode.cn id=208 lang=golang
 *
 * [208] 实现 Trie (前缀树)
 */

// @lc code=start
type Trie struct {
	trie  [][]int
	count []int
	index int
}

func Constructor() Trie {
	var t = Trie{
		trie:  make([][]int, 100110),
		count: make([]int, 100110),
	}
	for i, _ := range t.trie {
		t.trie[i] = make([]int, 26)
	}

	return t

}

func (this *Trie) Insert(word string) {
	p := 0
	for i := 0; i < len(word); i++ {
		el := word[i] - 'a'
		if this.trie[p][el] == 0 {
			this.index += 1
			this.trie[p][el] = this.index
		}
		p = this.trie[p][el]
	}
	this.count[p] += 1
}

func (this *Trie) Search(word string) bool {
	p := 0
	for i := 0; i < len(word); i++ {
		el := word[i] - 'a'
		if this.trie[p][el] == 0 {
			return false
		}
		p = this.trie[p][el]
	}
	return this.count[p] != 0
}

func (this *Trie) StartsWith(prefix string) bool {
	p := 0
	for i := 0; i < len(prefix); i++ {
		el := prefix[i] - 'a'
		if this.trie[p][el] == 0 {
			return false
		}
		p = this.trie[p][el]
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
// @lc code=end
