package twosum

func twoSum(nums []int, target int) []int {
	ans := make([]int, 2)
	d := make(map[int]int)

	for i, num := range nums {
		t := target - num
		if val, ok := d[t]; ok {
			ans[0] = val
			ans[1] = i
			break
		}
		d[num] = i
	}
	return ans
}
