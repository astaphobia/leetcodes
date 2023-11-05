package array

import (
	"fmt"
	"sort"
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

// rotate to rotate an array of n elements to the right by k steps.
// handled with 3 steps:
//  1. reverse the whole array
//  1. reverse the first - k elements
//  2. reverse the k -n elements
func rotate(nums []int, k int) {
	var reverse = func(ns []int, start, end int) {
		for i := start; i < end; i++ {
			ns[i], ns[end] = ns[end], ns[i]
			end--
		}
		return
	}

	// make sure k is valid
	k %= len(nums)

	end := len(nums) - 1

	reverse(nums, 0, end)
	fmt.Println(nums)
	reverse(nums, 0, k-1)
	fmt.Println(nums)
	reverse(nums, k, end)
	fmt.Println(nums)
	return
}

// containsDuplicate is to check if an array contains duplicates.
// https://leetcode.com/explore/interview/card/top-interview-questions-easy/92/array/578/
func containsDuplicate(nums []int) bool {
	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			return true
		}
	}
	return false
}
