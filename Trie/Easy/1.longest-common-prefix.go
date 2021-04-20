// author:- https://github.com/piyushkumar96/Data-Structure-In-Go
/*
    Topic:- Longest Common Prefix
	Write a function to find the longest common prefix string amongst an array of strings.
    If there is no common prefix, return an empty string "".

	Input: strs = ["flower","flow","flight"]
	Output: "fl"

    Input: strs = ["dog","racecar","car"]
    Output: ""
*/

package main

import (
	"fmt"
)

const (
	NO_OF_CHARS = 26
)

type TrieNode struct {
	children  [NO_OF_CHARS]*TrieNode
	isWordEnd bool
}

type Trie struct {
	root *TrieNode
}

func initTrie() *Trie {
	return &Trie{
		root: &TrieNode{},
	}
}

func (t *Trie) insertWordInTrie(word string) {
	wordLen := len(word)
	curr := t.root

	for i := 0; i < wordLen; i++ {
		index := word[i] - 'a'
		if curr.children[index] == nil {
			curr.children[index] = &TrieNode{}
		}

		curr = curr.children[index]
	}

	curr.isWordEnd = true

}

func createTrie(words []string) *Trie {
	trie := initTrie()
	for i := 0; i < len(words); i++ {
		trie.insertWordInTrie(words[i])
	}

	return trie
}

func (t *Trie) countChildren(node *TrieNode, index *int) int {
	count := 0
	for i := 0; i < NO_OF_CHARS; i++ {
		if node.children[i] != nil {
			count++
			*index = i
		}
	}

	return count
}

func (t *Trie) walkTrie() string {
	str := ""
	var index int
	curr := t.root
	for t.countChildren(curr, &index) == 1 && curr.isWordEnd == false {
		curr = curr.children[index]
		str = str + string(rune(index+'a'))
	}

	return str
}

func longestCommonPrefix(words []string) string {
	t := createTrie(words)
	return t.walkTrie()
}

func main() {
	var n int
	fmt.Println("Enter the number of words:- \t")
	fmt.Scan(&n)

	words := make([]string, n)
	fmt.Println("Enter the words:- ")
	for i := 0; i < n; i++ {
		fmt.Scan(&words[i])
	}

	comPrefix := longestCommonPrefix(words)
	fmt.Printf("The longest common prefix is %s", comPrefix)
}
