package main

// Given an array of integers, return indices of the two numbers such that they add up to a specific target.
// You may assume that each input would have exactly one solution, and you may not use the same element twice.

//solution: Runtime: 36 ms, faster than 30.36% of Go online submissions for Two Sum.
//	Memory Usage: 2.9 MB, less than 100.00% of Go online submissions for Two Sum.

// next steps: refresh hashmap

func twosum(nums []int, target int) []int {
	lnums := len(nums)
	for i1 := 0; i1 < lnums; i1++ {
		for i2 := i1 + 1; i2 < lnums; i2++ {
			if nums[i1]+nums[i2] == target {
				return []int{i1, i2}
			}
		}
	}
	return nil
}
