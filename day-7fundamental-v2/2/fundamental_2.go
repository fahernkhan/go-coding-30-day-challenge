package main

import (
	"fmt"
)

// ========== Fundamental Utama ==========
// 1. Two Pointers Technique
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

// 2. Sliding Window
func longestUniqueSubstring(s string) int {
	charIndex := make(map[byte]int)
	maxLength := 0
	start := 0

	for end := 0; end < len(s); end++ {
		if idx, exists := charIndex[s[end]]; exists && idx >= start {
			start = idx + 1
		}
		charIndex[s[end]] = end
		currentLength := end - start + 1
		if currentLength > maxLength {
			maxLength = currentLength
		}
	}
	return maxLength
}

// 3. Binary Search
func binarySearch(arr []int, target int) int {
	low, high := 0, len(arr)-1

	for low <= high {
		mid := low + (high-low)/2 // Prevent overflow
		if arr[mid] == target {
			return mid
		}
		if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

// 4. Recursion
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

// 5. Divide and Conquer (Merge Sort)
func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	result := make([]int, 0)
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result
}

func main() {
	// Two Pointers Example
	sortedArr := []int{2, 5, 7, 11, 15}
	target := 18
	twoSumResult := twoSumSorted(sortedArr, target)
	fmt.Printf("Two Pointers Example:\nArray: %v\nTarget: %d\nResult: %v\n\n",
		sortedArr, target, twoSumResult)

	// Sliding Window Example
	inputString := "abcabcbb"
	maxSubstring := longestUniqueSubstring(inputString)
	fmt.Printf("Sliding Window Example:\nInput: %s\nLongest Unique Substring Length: %d\n\n",
		inputString, maxSubstring)

	// Binary Search Example
	searchArr := []int{1, 3, 5, 7, 9, 11, 13, 15}
	searchTarget := 9
	searchResult := binarySearch(searchArr, searchTarget)
	fmt.Printf("Binary Search Example:\nArray: %v\nTarget: %d\nFound at index: %d\n\n",
		searchArr, searchTarget, searchResult)

	// Recursion Example
	factNum := 5
	fmt.Printf("Recursion Example:\nFactorial of %d = %d\n\n",
		factNum, factorial(factNum))

	// Merge Sort Example
	unsorted := []int{38, 27, 43, 3, 9, 82, 10}
	fmt.Printf("Merge Sort Example:\nUnsorted: %v\n", unsorted)
	fmt.Printf("Sorted: %v\n\n", mergeSort(unsorted))

	// Kompleksitas Time Analysis
	fmt.Println("Complexity Analysis:")
	fmt.Println("- Two Pointers: O(n)")
	fmt.Println("- Sliding Window: O(n)")
	fmt.Println("- Binary Search: O(log n)")
	fmt.Println("- Merge Sort: O(n log n)")
	fmt.Println("- Factorial Recursion: O(n)")
}

/*
OUTPUT EJAKUSI:

Two Pointers Example:
Array: [2 5 7 11 15]
Target: 18
Result: [1 2]

Sliding Window Example:
Input: abcabcbb
Longest Unique Substring Length: 3

Binary Search Example:
Array: [1 3 5 7 9 11 13 15]
Target: 9
Found at index: 4

Recursion Example:
Factorial of 5 = 120

Merge Sort Example:
Unsorted: [38 27 43 3 9 82 10]
Sorted: [3 9 10 27 38 43 82]

Complexity Analysis:
- Two Pointers: O(n)
- Sliding Window: O(n)
- Binary Search: O(log n)
- Merge Sort: O(n log n)
- Factorial Recursion: O(n)
*/
