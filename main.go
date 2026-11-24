package main

import "fmt"

// PlusOne increments a big number stored as decimal digits (most significant
// first). Walk from the end: a non-9 digit absorbs the carry and ends it;
// all-9s grow the number by one digit.
func PlusOne(digits []int) []int {
	out := append([]int{}, digits...)
	for i := len(out) - 1; i >= 0; i-- {
		if out[i] < 9 {
			out[i]++
			return out
		}
		out[i] = 0 // 9 rolls over; carry continues leftward
	}
	return append([]int{1}, out...) // every digit was 9
}

func main() {
	fmt.Println(PlusOne([]int{1, 2, 9}))
	fmt.Println(PlusOne([]int{9, 9, 9}))
}
