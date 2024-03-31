package main

func findMin(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	low := 0
	high := n - 1
	ans := nums[0]
	for low <= high {
		mid := low + (high-low)/2
		if nums[low] <= nums[high] {
			ans = min(ans, nums[low])
			return ans
		} else {
			if nums[low] <= nums[mid] {
				ans = min(ans, nums[mid])
				low = mid + 1
			} else {
				ans = min(ans, nums[mid])
				high = mid - 1
			}
		}
	}
	return ans
}
