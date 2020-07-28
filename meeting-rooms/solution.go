package main

import (
	"fmt"
	"sort"
)

type Array [][]int
func (a Array) Len() int { return len(a) }
func (a Array) Less(i,j int) bool {
	for x := range a[i] {
		if a[i][x] == a[j][x] {
			continue
		}
		return a[i][x] < a[j][x]
	}
	return false
}
func (a Array) Swap(i,j int) { a[i], a[j] = a[j], a[i] }

func main() {
	nums := [][]int{
		{0,30},{5,10},{15,20},
	}
	exp := false
	fmt.Println("Example 1:")
	fmt.Printf("  Input: n = %v\n", nums)
	fmt.Println("  Expected output: ", exp)
	fmt.Println("  Actual output: ", canAttendMeetings(nums))

	nums = [][]int{
		{7,10},{2,4},
	}
	exp = true
	fmt.Println("Example 1:")
	fmt.Printf("  Input: n = %v\n", nums)
	fmt.Println("  Expected output: ", exp)
	fmt.Println("  Actual output: ", canAttendMeetings(nums))
}

func canAttendMeetings(intervals [][]int) bool {
	arr := Array(intervals)
	sort.Sort(&arr)
	for i:=0; i < len(arr)-1; i++ {
		if arr[i][1] <= arr[i+1][0] {
			continue
		}
		return false
	}
	return true
}

