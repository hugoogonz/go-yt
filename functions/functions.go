package functions

// Return Negative
// In this simple assignment you are given a number and have to make it negative. But maybe the number is already negative?

// Examples
// MakeNegative(1)    // return -1
// MakeNegative(-5)   // return -5
// MakeNegative(0)    // return 0

// Link: https://www.codewars.com/kata/return-negative

func MakeNegative(x int) int {
	if x > 0 {
		return -x
	}
	return x
}
