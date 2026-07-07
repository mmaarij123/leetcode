func lengthOfLongestSubstring(s string) int {
	lastSeen := make(map[byte]int)

	left := 0
	ans := 0

	for right := 0; right < len(s); right++ {
		if idx, exists := lastSeen[s[right]]; exists && idx >= left {
			left = idx + 1
		}

		lastSeen[s[right]] = right

		if right-left+1 > ans {
			ans = right - left + 1
		}
	}

	return ans
}