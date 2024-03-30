package palindromenumber

func isPalindrome(x int) bool {
	original_num := x
	reverse_num := 0

	for x > 0 {
		reverse_num = reverse_num*10 + x%10
		x = x / 10
	}

	if original_num == reverse_num {
		return true
	} else {
		return false
	}
}
