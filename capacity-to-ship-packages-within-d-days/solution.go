package main

import (
	"fmt"
	"math"
)

func main() {
	weights := []int{1,2,3,4,5,6,7,8,9,10}
	D := 5
	fmt.Println("Example 1:")
	fmt.Println("Expected outuput: 15")
	fmt.Println("Actual outuput: ", shipWithinDays(weights, D))

	weights = []int{3,2,2,4,1,4}
	D = 3
	fmt.Println("Example 2:")
	fmt.Println("Expected outuput: 6")
	fmt.Println("Actual outuput: ", shipWithinDays(weights, D))

	weights = []int{1,2,3,1,1}
	D = 4
	fmt.Println("Example 3:")
	fmt.Println("Expected outuput: 3")
	fmt.Println("Actual outuput: ", shipWithinDays(weights, D))
}

func shipWithinDays(weights []int, D int) int {
    low, high := 1, math.MaxInt32
    for low < high {
        mid := low + (high-low)/2
        if validateFitInD(weights, D, mid) { 
            high = mid
        } else {
            low = mid + 1
        }        
    }
    return low    
}

func validateFitInD(weights []int, D int, capacity int) bool {
    cnt := 0
    cur := 0
    for _, v := range weights {
        if v > capacity {
            return false
        }
        if cur + v > capacity {
            cnt++
            cur = 0
        }
        cur += v
    }
    if cur > 0 {
        cnt++
    }
    return cnt <= D
}