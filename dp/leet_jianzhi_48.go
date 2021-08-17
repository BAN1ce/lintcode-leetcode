package dp

func lengthOfLongestSubstring(s string) int {

	m := make(map[int32]int)

	l := len(s)

	if l == 0 {
		return 0
	}
	maximum := 0
	tmp := 0
	index := 0
	ok := false
	for i, v := range s {

		if index, ok = m[v]; !ok {
			index = -1
		}
		if i-index > tmp {
			tmp = tmp + 1
		} else {
			tmp = i - index
		}
		maximum = max(maximum, tmp)
		m[v] = i
	}
	return maximum

}
