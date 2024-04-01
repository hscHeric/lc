package main

func findPoisonedDuration(timeSeries []int, duration int) int {
	ans := duration

	for i, x := range timeSeries[1:] {
		ans += min(duration, x-timeSeries[i])
	}
	return ans
}
