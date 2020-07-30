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

func (h InHeap) Len() int           { return len(h) }
func (h InHeap) Less(i, j int) bool { return h[i].x+h[i].y < h[j].x+h[j].y }
func (h InHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

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

func kSamllestPairs(nums1 []int, nums2 []int, k int) [][]int {

	if len(nums1) == 0 || len(nums2) == 0 {
		return nil
	}

	h := InHeap{}
	heap.Init(&h)
	for i := 0; i < len(nums1) && i < k; i++ {
		heap.Push(&h, point{0, nums1[i], nums2[0]})
	}

	var result [][]int
	for count := 0; len(h) > 0; {
		p := heap.Pop(&h).(point)
		result = append(result, []int{p.x, p.y})
		count++
		if count == k {
			break
		}
		if p.index+1 < len(nums2) {
			heap.Push(&h, point{p.index + 1, p.x, nums2[p.index+1]})
		}
	}
	return result
}

func main() {

	nums1 := []int{1, 7, 11}
	nums2 := []int{2, 4, 6}
	k := 3
	out := kSamllestPairs(nums1, nums2, k)
	fmt.Println(out)
}
