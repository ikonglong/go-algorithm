package heap

import "context"

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */

// https://leetcode.cn/problems/implement-trie-prefix-tree/

type TrieNode struct {
	isEnding bool // 当前节点表示的字符是否是一个单词的结束
	// data     byte // byte = Character，该字段没有使用，注释掉
	children [26]*TrieNode
}

type Trie struct {
	root *TrieNode
	ctx  context.Context
}

func NewTrieTree() Trie { // leetcode: Constructor
	return Trie{
		root: &TrieNode{},
	}
}

func (this *Trie) Insert(word string) {
	p := this.root
	for i := 0; i < len(word); i++ {
		c := word[i]
		if p.children[c-'a'] == nil {
			p.children[c-'a'] = &TrieNode{}
		}
		p = p.children[c-'a']
	}
	p.isEnding = true
}

func (this *Trie) Search(word string) bool {
	p := this.root
	for _, c := range word {
		if p.children[c-'a'] == nil {
			return false
		}
		p = p.children[c-'a']
	}
	return p.isEnding
	// bug
	// return true
}

func (this *Trie) StartsWith(prefix string) bool {
	p := this.root
	for _, c := range prefix {
		if p.children[c-'a'] == nil {
			return false
		}
		p = p.children[c-'a']
	}
	return true
}
