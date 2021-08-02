package queue

type Char struct {
	v     int32
	index int
}

// 剑指 Offer 50. 第一个只出现一次的字符
// https://leetcode-cn.com/problems/di-yi-ge-zhi-chu-xian-yi-ci-de-zi-fu-lcof/

func firstUniqChar(s string) byte {
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

		if uniq, ok := m[v.v]; ok&& uniq  {
			return byte(v.v)
		}
	}
	return ' '

}
