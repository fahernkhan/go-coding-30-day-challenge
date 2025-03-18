package main

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	testCases := []struct {
		nums   []int
		target int
		expect []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
		{[]int{1, 2, 3, 4, 5}, 9, []int{3, 4}},
		{[]int{1, 2, 3, 4, 5}, 3, []int{0, 1}},
		{[]int{-1, -2, -3, -4, -5}, -8, []int{2, 3}},
		{[]int{-1, 0, 1, 2, -1, -4}, 0, []int{0, 2}},
		{[]int{1}, 5, nil},
		{[]int{}, 5, nil},
		{[]int{1, 5, 2, 6}, 11, []int{1, 3}},
		{[]int{1, 5, 2, 6, 10, 12, 3}, 13, []int{0, 5}},
		{[]int{1, 5, 2, 6, 10, 12, 3}, 11, []int{2, 3}},
	}

	for _, tc := range testCases {
		result := twoSum(tc.nums, tc.target)
		if !reflect.DeepEqual(result, tc.expect) {
			t.Errorf("For nums: %v, target: %d, expected: %v, got: %v", tc.nums, tc.target, tc.expect, result)
		}
	}
}
