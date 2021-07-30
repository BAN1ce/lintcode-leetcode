package sliding_window


// 1876. 长度为三且各字符不同的子字符串
// https://leetcode-cn.com/problems/substrings-of-size-three-with-distinct-characters/
func countGoodSubstrings(s string) int {

	left := 0
	right := 3

	count := 0
	for ; right <= len(s); {
		tmp := s[left:right]
		if isGood(tmp) {
			count++
		}
		left += 1
		right += 1
	}
	return count
}

func isGood(s string) bool {

	if s[0] == s[1] || s[0] == s[2] || s[1] == s[2] {
		return false
	}
	return true
}
