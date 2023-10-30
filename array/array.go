package array

import (
	"fmt"
)

// removeDuplicates is test to remove duplicates in a sorted array in-place algorithm
// https://leetcode.com/explore/interview/card/top-interview-questions-easy/92/array/727/
// https://en.wikipedia.org/wiki/In-place_algorithm
func removeDuplicates(nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	for i := 0; i < len(nums); i++ {
		if i == len(nums)-1 {
			break
		}

		if nums[i] == nums[i+1] {
			nums = append(nums[:i], nums[i+1:]...)
			fmt.Printf("i: %d: nums: %+v\n", i, nums)
			i--
		}

	}
	return len(nums)
}
