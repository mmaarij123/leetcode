func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	start, end := 0, 0

	expand := func(left, right int) {
		for left >= 0 && right < len(s) && s[left] == s[right] {
			left--
			right++
		}

		// Adjust back to the last valid palindrome
		left++
		right--

		if right-left > end-start {
			start = left
			end = right
		}
	}

	for i := 0; i < len(s); i++ {
		// Odd-length palindrome
		expand(i, i)

		// Even-length palindrome
		expand(i, i+1)
	}

	return s[start : end+1]
}