package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := [][]int{
		{0,30},{5,10},{15,20},
	}
	exp := 2
	fmt.Println("Example 1:")
	fmt.Printf("  Input: n = %d\n", nums)
	fmt.Println("  Expected output: ", exp)
	fmt.Println("  Actual output: ", minMeetingRooms(nums))

	nums = [][]int{
		{7,10},{2,4},
	}
	exp = 1
	fmt.Println("Example 1:")
	fmt.Printf("  Input: n = %d\n", nums)
	fmt.Println("  Expected output: ", exp)
	fmt.Println("  Actual output: ", minMeetingRooms(nums))
}


func minMeetingRooms(intervals [][]int) int { 
    
    if len(intervals) == 0 {
        return 0
    }
    
    starts := make([]int,len(intervals))
    ends := make([]int,len(intervals))  
    for i, v := range intervals {
        starts[i] = v[0]
        ends[i] = v[1]
    }
    
    sort.Ints(starts)
    sort.Ints(ends)
    cnt := 0
    endIterator := 0
    for i := range starts {
        if starts[i] < ends[endIterator] {
            cnt++
        } else {
            endIterator++
        }               
    }
    return cnt    
}
