package sliding_window

import (
	"strings"
)

// 无重复字符的最长子串
// https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/
func lengthOfLongestSubstring(s string) int {

	if s == "" {
		return 0
	}
	left, right := 0, 1

	max := 1
	flag := false
	for ; right < len(s); right++ {

		q := s[left:right]
		next := string(s[right])
		if tmp := strings.Index(q, next); tmp != -1 {
			flag = true
			if (right - left) > max {
				max = right - left
			}
			left += tmp + 1
			// 特殊处理 aab
		} else if right == len(s)-1 {
			if (right - left + 1) > max {
				max = right - left + 1
			}
		}
	}

	if !flag {
		return len(s)
	}
	return max

}
