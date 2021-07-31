package main

import (
	"container/heap"
	"fmt"
)

type InHeap []int

func (h InHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h InHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h InHeap) Len() int           { return len(h) }

func (h *InHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *InHeap) Pop() interface{} {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}

// KthLargest struct will be instantiated and called as such:
// obj := Constructor(k, nums);
// param_1 := obj.Add(val);
type KthLargest struct {
	k int
	h *InHeap
}

func Constructor(k int, nums []int) KthLargest {
	h := &InHeap{}
	heap.Init(h)
	for i := 0; i < len(nums); i++ {
		heap.Push(h, nums[i])
	}
	for len(*h) > k {
		heap.Pop(h)
	}
	return KthLargest{k, h}
}

func (this *KthLargest) Add(val int) int {
	if len(*this.h) < this.k {
		heap.Push(this.h, val)
	} else if val > (*this.h)[0] {
		heap.Push(this.h, val)
		heap.Pop(this.h)
	}
	return (*this.h)[0]
}

func main() {
	k := 3
	arr := []int{4, 5, 8, 2}
	kthLargest := Constructor(k, arr)
	fmt.Println(kthLargest.Add(3))  // returns 4
	fmt.Println(kthLargest.Add(5))  // returns 5
	fmt.Println(kthLargest.Add(10)) // returns 5
	fmt.Println(kthLargest.Add(9))  // returns 8
	fmt.Println(kthLargest.Add(4))  // returns 8
}
