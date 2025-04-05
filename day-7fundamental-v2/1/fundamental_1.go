package main

import "fmt"

func exampleBigO(n int) {
	// O(1) - Operasi konstan
	fmt.Println(n)

	// O(n) - Loop linear
	for i := 0; i < n; i++ {
		fmt.Println(i)
	}

	// O(n^2) - Nested loop
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Println(i, j)
		}
	}

	// O(log n) - Binary Search
	for i := n; i > 0; i /= 2 {
		fmt.Println(i)
	}
}

func twoSumSorted(arr []int, target int) []int {
	left, right := 0, len(arr)-1
	for left < right {
		sum := arr[left] + arr[right]
		if sum == target {
			return []int{left, right}
		}
		if sum < target {
			left++
		} else {
			right--
		}
	}
	return nil
}

func longestUniqueSubstring(s string) int {
	charIndex := make(map[byte]int)
	maxLength := 0
	start := 0

	for end := 0; end < len(s); end++ {
		if idx, exists := charIndex[s[end]]; exists && idx >= start {
			start = idx + 1
		}
		charIndex[s[end]] = end
		if end-start+1 > maxLength {
			maxLength = end - start + 1
		}
	}
	return maxLength
}

func binarySearch(arr []int, target int) int {
	low, high := 0, len(arr)-1

	for low <= high {
		mid := low + (high-low)/2 // Hindari overflow
		if arr[mid] == target {
			return mid
		}
		if arr[mid] < target {
			low = mid + 1 // Putuskan setengah kiri
		} else {
			high = mid - 1 // Putuskan setengah kanan
		}
	}
	return -1
}
