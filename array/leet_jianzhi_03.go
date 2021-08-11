package array

// 剑指 Offer 03. 数组中重复的数字
// https://leetcode-cn.com/problems/shu-zu-zhong-zhong-fu-de-shu-zi-lcof/
func findRepeatNumber(nums []int) int {

	m := make(map[int]bool)

	for _, v:=range nums{
		if _, ok:= m[v];ok{
			return v
		}
		m[v] = true
	}
	return 0
}

