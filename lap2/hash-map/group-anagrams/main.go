package main

import (
	"fmt"
	"sort"
)

type sortRunes []rune

func (s sortRunes) Len() int           { return len(s) }
func (s sortRunes) Less(i, j int) bool { return s[i] < s[j] }
func (s sortRunes) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func groupAnagrams(strs []string) [][]string {
	anagrams := make(map[string][]string)
	for _, str := range strs {
		s := []rune(str)
		sort.Sort(sortRunes(s))
		sortedStr := string(s)

		if arr, ok := anagrams[sortedStr]; ok {
			arr = append(arr, str)
			anagrams[sortedStr] = arr
		} else {
			var arr []string
			arr = append(arr, str)
			anagrams[sortedStr] = arr
		}
	}
	result := make([][]string, len(anagrams))
	i := 0
	for _, strs := range anagrams {
		result[i] = strs
		i++
	}
	return result
}

func main() {
	type args struct {
		strs []string
	}
	type test struct {
		name string
		arg  args
		exp  [][]string
	}
	tests := []test{
		{
			name: "Example 1:",
			arg: args{
				strs: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			},
			exp: [][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}},
		},
		{
			name: "Example 2:",
			arg: args{
				strs: []string{""},
			},
			exp: [][]string{{""}},
		},
		{
			name: "Example 2:",
			arg: args{
				strs: []string{"a"},
			},
			exp: [][]string{{"a"}},
		},
	}
	for _, tt := range tests {
		fmt.Println("Case:", tt.name)
		fmt.Println(" Actual: ", groupAnagrams(tt.arg.strs))
		fmt.Println(" Exp: ", tt.exp)
	}
}
