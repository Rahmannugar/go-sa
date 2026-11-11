package main

import "fmt"

type TrieNode struct {
	children map[rune]*TrieNode
	isWord   bool
}

type Trie struct {
	root *TrieNode
}

// NewTrie creates an empty trie.
func NewTrie() *Trie {
	return &Trie{
		root: &TrieNode{
			children: make(map[rune]*TrieNode),
		},
	}
}

// Add stores a word, creating nodes for characters that are not already present.
// Time complexity: O(m), Space complexity: O(m)
func (t *Trie) Add(word string) {
	current := t.root

	for _, char := range word {
		if _, exists := current.children[char]; !exists {
			current.children[char] = &TrieNode{
				children: make(map[rune]*TrieNode),
			}
		}

		current = current.children[char]
	}

	current.isWord = true
}

// Exists reports whether the exact word is stored in the trie.
// Time complexity: O(m), Space complexity: O(1)
func (t *Trie) Exists(word string) bool {
	current := t.root

	for _, char := range word {
		child, exists := current.children[char]
		if !exists {
			return false
		}

		current = child
	}

	return current.isWord
}

// searchLevel collects stored words below the current node.
func (t *Trie) searchLevel(
	current *TrieNode,
	prefix string,
	words *[]string,
) {
	if current.isWord {
		*words = append(*words, prefix)
	}

	for char, child := range current.children {
		t.searchLevel(child, prefix+string(char), words)
	}
}

// WordsWithPrefix returns stored words that begin with the given prefix.
func (t *Trie) WordsWithPrefix(prefix string) []string {
	current := t.root

	for _, char := range prefix {
		child, exists := current.children[char]
		if !exists {
			return nil
		}

		current = child
	}

	words := []string{}
	t.searchLevel(current, prefix, &words)

	return words
}

// FindMatches returns stored words found within the document.
func (t *Trie) FindMatches(document string) []string {
	matches := []string{}
	chars := []rune(document)

	for i := 0; i < len(chars); i++ {
		current := t.root

		for j := i; j < len(chars); j++ {
			child, exists := current.children[chars[j]]
			if !exists {
				break
			}

			current = child

			if current.isWord {
				matches = append(matches, string(chars[i:j+1]))
			}
		}
	}

	return matches
}

func main() {
	trie := NewTrie()

	trie.Add("developer")
	trie.Add("development")
	trie.Add("devops")
	trie.Add("designer")
	trie.Add("design")

	fmt.Println("Exists:")
	fmt.Println(trie.Exists("developer"))
	fmt.Println(trie.Exists("develop"))

	fmt.Println("\nWords with prefix:")
	fmt.Println(trie.WordsWithPrefix("dev"))
	fmt.Println(trie.WordsWithPrefix("des"))

	fmt.Println("\nSubstring matches:")

	matchTrie := NewTrie()
	matchTrie.Add("cat")
	matchTrie.Add("car")
	matchTrie.Add("dog")

	document := "The car chased the dog."
	fmt.Println(matchTrie.FindMatches(document))
}
