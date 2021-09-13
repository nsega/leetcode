package main

import "fmt"

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
			for _, candidate := range candidates(word) {
				if _, ok := wm[candidate]; ok {
					if candidate == endWord {
						return depth + 1
					}
					delete(wm, candidate)
					que = append(que, candidate)
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
	var res []string
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

	type args struct {
		beginWord string
		endWord   string
		wordList  []string
	}
	type test struct {
		name string
		arg  args
		exp  int
	}
	tests := []test{
		{
			name: "Example 1:",
			arg: args{
				beginWord: "hit",
				endWord:   "cog",
				wordList: []string{
					"hot", "dot", "dog", "lot", "log", "cog",
				},
			},
			exp: 5,
		},
		{
			name: "Example 2:",
			arg: args{
				beginWord: "hit",
				endWord:   "cog",
				wordList: []string{
					"hot", "dot", "dog", "lot", "log",
				},
			},
			exp: 0,
		},
	}
	for _, tt := range tests {
		fmt.Println("Case:", tt.name)
		fmt.Println(" Actual: ", ladderLength(tt.arg.beginWord, tt.arg.endWord, tt.arg.wordList))
		fmt.Println(" Exp: ", tt.exp)
	}
}
