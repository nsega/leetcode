package main

import (
	"fmt"
	"sort"
)

type Pair struct {
	Number int
	Count  int
}

type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Count < p[j].Count }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func topKFrequent(nums []int, k int) []int {

	fmap := make(map[int]int)
	for _, num := range nums {
		fmap[num]++
	}

	pl := make(PairList, len(fmap))
	i := 0
	for k, v := range fmap {
		pl[i] = Pair{k, v}
		i++
	}
	sort.Sort(sort.Reverse(pl))

	result := make([]int, k)
	i = 0
	for _, v := range pl {
		if i >= k {
			break
		}
		result[i] = v.Number
		i++
	}
	return result
}

func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	k := 2
	output := topKFrequent(nums, k)
	fmt.Println(output)
}
