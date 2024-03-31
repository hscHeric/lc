package main

func findMin(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	middle := len(nums) / 2
	if middle > 0 && middle < len(nums)-1 && nums[middle] < nums[middle-1] && nums[middle] < nums[middle+1] {
		return nums[middle]
	}
	min1 := findMin(nums[:middle])
	min2 := findMin(nums[middle:])

	return min(min1, min2)
}
