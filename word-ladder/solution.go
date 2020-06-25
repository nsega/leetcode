package main

import (
	"fmt"
)

func ladderLength(beginWord string, endWord string, wordList []string) int {

	wm := wordMap(wordList, beginWord)
	que := []string{beginWord}
	depth := 0
	for len(que) > 0 {
		depth++

		qlen := len(que)
		for i := 0; i < qlen; i++ {
			word := que[0]
			que = que[1:]

			candidates := candidates(word)
			for _, c := range candidates {
				if _, ok := wm[c]; ok {
					if c == endWord {
						return depth + 1
					}

					delete(wm, c)
					que = append(que, c)
				}
			}
		}
	}
	return 0
}

func wordMap(wordList []string, beginWord string) map[string]int {
	wm := make(map[string]int)
	for i, word := range wordList {
		if _, ok := wm[word]; !ok {
			if word != beginWord {
				wm[word] = i
			}
		}
	}
	return wm
}

func candidates(word string) []string {
	res := []string{}
	for i := 0; i < 26; i++ {
		for j := 0; j < len(word); j++ {
			if word[j] != byte(int('a')+i) {
				res = append(res, word[:j]+string(int('a')+i)+word[j+1:])
			}
		}
	}
	return res
}

func main() {
	beginWord := "hit"
	endWord := "cog"
	wordList := []string{"hot", "dot", "dog", "lot", "log", "cog"}
	fmt.Println(ladderLength(beginWord, endWord, wordList))
}
