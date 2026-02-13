/*
LeetCode Problem #208: Implement Trie (Prefix Tree)
Difficulty: Medium

A trie (pronounced as 'try') or prefix tree is a tree data structure used to efficiently store and retrieve keys in a dataset of strings. Implement the Trie class.
*/

package LC208

type TrieNode struct {
	childeren  [26]*TrieNode
	isTerminal bool
}
type Trie struct {
	root *TrieNode
}

func Constructor() Trie {

	return Trie{root: &TrieNode{}}
}

func (this *Trie) Insert(word string) {
	node := this.root
	for _, v := range word {
		ch := v - 'a'
		if nil == node.childeren[ch] {
			node.childeren[ch] = &TrieNode{}
		}
		node = node.childeren[ch]
	}
	node.isTerminal = true
}

func (this *Trie) Search(word string) bool {

	node := this.root

	for _, v := range word {
		ch := v - 'a'
		if nil == node.childeren[ch] {
			return false
		}
		node = node.childeren[ch]
	}
	return node.isTerminal

}

func (this *Trie) StartsWith(prefix string) bool {
	node := this.root

	for _, v := range prefix {
		ch := v - 'a'
		if node.childeren[ch] == nil {
			return false
		}
		node = node.childeren[ch]
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
