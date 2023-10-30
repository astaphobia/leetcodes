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

// maxProfit is test to find the maximum profit from buying and selling stock in given array of prices.
// https://leetcode.com/explore/interview/card/top-interview-questions-easy/92/array/564/
func maxProfit(prices []int) int {

	var maxProfit int
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			maxProfit += prices[i] - prices[i-1]
		}
	}

	return maxProfit
}
