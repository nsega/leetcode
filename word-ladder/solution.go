package main

import (
	"fmt"
)

func ladderLength(beginWord string, endWord string, wordList []string) int {
	m := make(map[string]bool)
	for _, v := range wordList {
		m[v] = true
	}

	var list []string
	list = append(list, beginWord)

	res := 1
	for len(list) != 0 {
		l := len(list)
		for i := 0; i < l; i++ {
			word := list[0]
			list = list[1:]
			if word == endWord {
				return res
			}
			list = findWord(word, m, list)
		}
		res++
	}
	return 0
}

func findWord(word string, m map[string]bool, l []string) []string {
	m[word] = false
	for ch := range word {
		for j := 0; j < 26; j++ {
			tmp := []rune(word)
			tmp[ch] = rune('a' + j)
			str := string(tmp)
			if _, ok := m[str]; ok {
				if len(l) != 0 && str != l[len(l)-1] {
					l = append(l, str)
				}
				if len(l) == 0 {
					l = append(l, str)
				}
			}
		}
	}
	return l
}

func main() {
	beginWord := "hit"
	endWord := "cog"
	wordList := []string{"hot", "dot", "dog", "lot", "log", "cog"}
	fmt.Println(ladderLength(beginWord, endWord, wordList))
}
