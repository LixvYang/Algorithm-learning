// 实现trie树
package leetcode

type Trie struct {
	child [26] *Trie
	isEnd bool
}

//初始化Trie
func Constructor() Trie {
	return Trie{}
}

//插入操作
func (this *Trie) Insert(word string)  {
	node := this
	for _, ch := range word {
			ch -= 'a'
			if node.child[ch] == nil {
					node.child[ch] = &Trie{}
			}
			node = node.child[ch]
	}
	node.isEnd = true
}

//查找是否存在单词or前缀（node.isEnd==true为单词）
func (this *Trie) SearchPrefix(prefix string) *Trie {
	node := this
	for _, ch := range prefix {
			ch -= 'a'
			if node.child[ch] == nil {
					return nil
			}
			node = node.child[ch]
	}
	return node
}

//查找单词
func (this *Trie) Search(word string) bool {
	node := this.SearchPrefix(word)
	if node != nil && node.isEnd == true {
			return true
	}
	return false
}

//查找树中是否存在前缀
func (this *Trie) StartsWith(prefix string) bool {
	if this.SearchPrefix(prefix) != nil {
			return true
	}
	return false
}
