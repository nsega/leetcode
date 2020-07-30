package main

import "fmt"

func main() {

	s := "abc"
	t := "ahbgdc"
	exp := true
	fmt.Println("Example 1:")
	fmt.Printf("  Input: s = %v, t = %v \n", s, t)
	fmt.Println("  Expected output: ", exp)
	fmt.Println("  Actual output: ", isSubsequence(s, t))

	s = "axc"
	t = "ahbgdc"
	exp = false
	fmt.Println("Example 2:")
	fmt.Printf("  Input: s = %v, t = %v \n", s, t)
	fmt.Println("  Expected output: ", exp)
	fmt.Println("  Actual output: ", isSubsequence(s, t))

	s = ""
	t = "ahbgdc"
	exp = true
	fmt.Println("Example 3:")
	fmt.Printf("  Input: s = %v, t = %v \n", s, t)
	fmt.Println("  Expected output: ", exp)
	fmt.Println("  Actual output: ", isSubsequence(s, t))
}

func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	idx := 0
	sch := []rune(s)
	for _, tch := range t {
		if sch[idx] == tch {
			idx++
		}
		if idx == len(s) {
			return true
		}
	}
	return false
}
