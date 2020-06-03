package main

import (
	"container/heap"
	"fmt"
)

type InHeap []int

func (h InHeap) Len() int           { return len(h) }
func (h InHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h InHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *InHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *InHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type KthLargest struct {
	k    int
	nums *InHeap
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
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
	if len(*this.nums) < this.k {
		heap.Push(this.nums, val)
	} else if val > (*this.nums)[0] {
		heap.Push(this.nums, val)
		heap.Pop(this.nums)
	}
	return (*this.nums)[0]
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
