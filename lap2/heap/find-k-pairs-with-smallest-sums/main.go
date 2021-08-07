package main

import (
	"container/heap"
	"fmt"
)

type point struct {
	index int
	x, y  int
}

type InHeap []point

func (h InHeap) Len() int { return len(h) }
func (h InHeap) Less(i, j int) bool {
	return h[i].x+h[i].y < h[j].x+h[j].y
}
func (h InHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *InHeap) Push(x interface{}) {
	*h = append(*h, x.(point))
}

func (h *InHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return nil
	}
	h := InHeap{}
	heap.Init(&h)
	for i := 0; i < len(nums1) && i < k; i++ {
		heap.Push(&h, point{0, nums1[i], nums2[0]})
	}

	var res [][]int
	for count := 0; len(h) > 0; {
		p := heap.Pop(&h).(point)
		res = append(res, []int{p.x, p.y})
		count++
		if count == k {
			break
		}
		if p.index+1 < len(nums2) {
			heap.Push(&h, point{p.index + 1, p.x, nums2[p.index+1]})
		}
	}
	return res
}

func main() {
	type args struct {
		nums1 []int
		nums2 []int
		k     int
	}
	type test struct {
		name string
		arg  args
		exp  [][]int
	}
	tests := []test{
		{
			name: "Example 1",
			arg: args{
				nums1: []int{1, 7, 11},
				nums2: []int{2, 4, 6},
				k:     3,
			},
			exp: [][]int{{1, 2}, {1, 4}, {1, 6}},
		},
		{
			name: "Example 2",
			arg: args{
				nums1: []int{1, 1, 2},
				nums2: []int{1, 2, 3},
				k:     2,
			},
			exp: [][]int{{1, 1}, {1, 1}},
		},
		{
			name: "Example 3",
			arg: args{
				nums1: []int{1, 2},
				nums2: []int{3},
				k:     3,
			},
			exp: [][]int{{1, 3}, {2, 3}},
		},
	}
	for _, tt := range tests {
		fmt.Println("Case:", tt.name)
		fmt.Println(" Actual: ", kSmallestPairs(tt.arg.nums1, tt.arg.nums2, tt.arg.k))
		fmt.Println(" Exp: ", tt.exp)
	}
}
