package array


// 剑指 Offer 53 - I. 在排序数组中查找数字 I
// https://leetcode-cn.com/problems/zai-pai-xu-shu-zu-zhong-cha-zhao-shu-zi-lcof/
func search(nums []int, target int) int {

	m := make(map[int]int)

	for _, v := range nums {
		if _, ok := m[v]; ok {
			m[v] += 1
		} else {
			m[v] = 1
		}
	}
	return m[target]
}
