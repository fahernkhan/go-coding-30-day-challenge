package main

import "fmt"

func twoSum(nums []int, target int) []int {
	for i, value := range nums {
		for j := i + 1; j < len(nums); j++ {
			if value+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func main() {
	testCases := []struct {
		nums   []int
		target int
		expect []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
		{[]int{1, 2, 3, 4, 5}, 10, nil},
	}

	for _, testCase := range testCases {
		result := twoSum(testCase.nums, testCase.target)
		fmt.Printf("nums: %v, target: %d, expected: %v, got: %v\n", testCase.nums, testCase.target, testCase.expect, result)
	}
}
