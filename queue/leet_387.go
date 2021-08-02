package queue

// 字符串中的第一个唯一字符
// https://leetcode-cn.com/problems/first-unique-character-in-a-string/
func firstUniqChar1(s string) int {

	m := make(map[int32]bool)
	index := make([]Char, 0)

	for i, v := range s {
		if _, ok := m[v]; ok {
			m[v] = false
		} else {
			m[v] = true
			index = append(index, Char{
				v:     v,
				index: i,
			})
		}
	}

	for _, v := range index {

		if uniq, ok := m[v.v]; ok && uniq {
			return v.index
		}
	}
	return -1
}
