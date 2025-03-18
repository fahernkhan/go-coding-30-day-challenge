package main

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

// Anda bisa menambahkan fungsi main untuk mencoba twoSum di sini,
// tapi untuk testing, main_test.go sudah cukup.
