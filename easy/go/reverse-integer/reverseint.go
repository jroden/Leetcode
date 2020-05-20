package reverseint

import (
	"math"
)

// https://leetcode.com/problems/reverse-integer

// Runtime: 4 ms, faster than 38.03% of Go online submissions for Reverse Integer.
// Memory Usage: 2.3 MB, less than 6.67% of Go online submissions for Reverse Integer.
// TODO: in v2 iteration solve with one loop & reduce memory by removing int slice

func reverse(x int) int {
	nums := []float64{}
	// extract each place into nums array
	for i := 10; i > 0; i = i * 10 {
		nums = append(nums, math.Floor(float64((x%i)/(i/10))))
		if x/i == 0 {
			break
		}
	}
	var mplier float64
	rval := 0
	// loop through nums array and append conversion result to rval
	for i, val := range nums {
		mplier = math.Pow(10, float64(len(nums)-(i+1)))
		rval = rval + int(mplier*val)
	}
	// handle overflow
	if rval > int((math.Pow(2, 31)-1)) || rval < int(-(math.Pow(2, 31))) {
		return 0
	} else if x < 0 && rval > 0 {
		return -rval
	}
	return rval
}
