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

	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(strs))

}
