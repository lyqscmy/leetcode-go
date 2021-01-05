package trie

type Trie struct {
	char     byte
	children [26]*Trie
	isLeaf   bool
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	node := this
	for i := 0; i < len(word); i++ {
		char := word[i]
		if node.children[char-'a'] == nil {
			node.children[char-'a'] = &Trie{char: char}
		}
		node = node.children[char-'a']
	}
	node.isLeaf = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	node := this
	for i := 0; i < len(word); i++ {
		char := word[i]
		if node.children[char-'a'] == nil {
			return false
		}
		node = node.children[char-'a']
	}
	return node.isLeaf
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	node := this
	for i := 0; i < len(prefix); i++ {
		char := prefix[i]
		if node.children[char-'a'] == nil {
			return false
		}
		node = node.children[char-'a']
	}
	return true
}
