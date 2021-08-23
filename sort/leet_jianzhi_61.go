package sort

// 剑指 Offer 61. 扑克牌中的顺子
// https://leetcode-cn.com/problems/bu-ke-pai-zhong-de-shun-zi-lcof/
func isStraight(nums []int) bool {

	kingNum := 0

	m := make(map[int]bool)
	min := 14
	max := 0
	for _,v := range nums {
		if v == 0 {
			kingNum++
		} else {
			if _, ok := m[v]; ok {
				return false
			} else {
				if v > max {
					max = v
				}
				if v < min {
					min = v
				}
				m[v] = true
			}
		}

	}
	if (max - min) <=4 {
		return true
	}

	return false
}
