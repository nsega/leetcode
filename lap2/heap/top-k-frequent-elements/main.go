package main

import (
	"fmt"
	"sort"
)

type Pair struct {
	num   int
	count int
}

type Pairs []Pair

func (p Pairs) Len() int           { return len(p) }
func (p Pairs) Less(i, j int) bool { return p[i].count < p[j].count }
func (p Pairs) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func topKFrequent(nums []int, k int) []int {
	fmap := make(map[int]int)
	for _, num := range nums {
		fmap[num]++
	}

	pairs := make(Pairs, len(fmap))
	i := 0
	for k, v := range fmap {
		pairs[i] = Pair{k, v}
		i++
	}
	sort.Sort(sort.Reverse(pairs))

	result := make([]int, k)
	i = 0
	for _, v := range pairs {
		if i >= k {
			break
		}
		result[i] = v.num
		i++
	}
	return result
}

func main() {
	type args struct {
		nums []int
		k    int
	}
	type test struct {
		name string
		arg  args
		exp  []int
	}
	tests := []test{
		{
			name: "Example 1",
			arg: args{
				nums: []int{1, 1, 1, 2, 2, 3},
				k:    2,
			},
			exp: []int{1, 2},
		},
		{
			name: "Example 2",
			arg: args{
				nums: []int{1},
				k:    1,
			},
			exp: []int{1},
		},
	}
	for _, tt := range tests {
		fmt.Println("Case: ", tt.name)
		fmt.Println(" Actual: ", topKFrequent(tt.arg.nums, tt.arg.k))
		fmt.Println(" Exp: ", tt.exp)
	}
}
