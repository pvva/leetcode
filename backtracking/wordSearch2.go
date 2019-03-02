package main

import "fmt"

func main() {
	//fmt.Println(findWords(
	//	[][]byte{
	//		{'o', 'a', 'a', 'n'},
	//		{'e', 't', 'a', 'e'},
	//		{'i', 'h', 'k', 'r'},
	//		{'i', 'f', 'l', 'v'},
	//	},
	//	[]string{"oath", "pea", "eat", "rain"},
	//))
	//fmt.Println(findWords(
	//	[][]byte{
	//		{'a', 'b'},
	//		{'c', 'd'},
	//	},
	//	[]string{"ab", "cb", "ad", "bd", "ac", "ca", "da", "bc", "db", "adcb", "dabc", "abb", "acb"},
	//))
	fmt.Println(findWords(
		[][]byte{
			{'a', 'b'},
			{'a', 'a'},
		},
		[]string{"aba", "baa", "bab", "aaab", "aaa", "aaaa", "aaba"},
	))
}

/*
Given a 2D board and a list of words from the dictionary, find all words in the board.

Each word must be constructed from letters of sequentially adjacent cell, where "adjacent" cells are
those horizontally or vertically neighboring. The same letter cell may not be used more than once in a word.

Input:
words = ["oath","pea","eat","rain"] and board =
[
  ['o','a','a','n'],
  ['e','t','a','e'],
  ['i','h','k','r'],
  ['i','f','l','v']
]

Output: ["eat","oath"]
*/
func findWords(board [][]byte, words []string) []string {
	l := len(board)
	if l == 0 {
		return []string{}
	}
	l *= len(board[0])

	trie := newTrie()
	for _, w := range words {
		if len(w) <= l {
			trie.AddString(w)
		}
	}

	result := []string{}
	for x := 0; x < len(board); x++ {
		for y := 0; y < len(board[x]); y++ {
			for char, t := range trie.next {
				if char != board[x][y] {
					continue
				}
				fwDfs(board, x, y, t, &result)
			}
		}
	}

	return result
}

var dirs = [][]int{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}

func fwDfs(board [][]byte, x, y int, trie *Trie, result *[]string) {
	p := board[x][y]
	if p == '-' {
		return
	}

	if trie.word != "" {
		*result = append(*result, trie.word)
		trie.word = ""
	}

	for _, dir := range dirs {
		nx := x + dir[0]
		ny := y + dir[1]

		if nx < 0 || nx >= len(board) || ny < 0 || ny >= len(board[nx]) {
			continue
		}

		r := board[nx][ny]
		if t, ex := trie.next[r]; ex {
			board[x][y] = '-'
			fwDfs(board, nx, ny, t, result)
			board[x][y] = p
		}
	}
}

type Trie struct {
	word string
	next map[byte]*Trie
}

func newTrie() *Trie {
	return &Trie{
		next: make(map[byte]*Trie),
	}
}

func (trie *Trie) AddString(s string) {
	ct := trie
	for _, r := range []byte(s) {
		nt, ex := ct.next[r]
		if !ex {
			nt = newTrie()
			ct.next[r] = nt
		}
		ct = nt
	}
	ct.word = s
}

func (trie *Trie) getNext(char byte) *Trie {
	if t, ex := trie.next[char]; ex {
		return t
	}

	return nil
}
