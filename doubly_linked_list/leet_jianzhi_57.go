package doubly_linked_list

// 剑指 Offer 57. 和为s的两个数字
// https://leetcode-cn.com/problems/he-wei-sde-liang-ge-shu-zi-lcof/
func twoSum(nums []int, target int) []int {

	m := make(map[int]int)
	for _, v := range nums {
		m[v] += 1
	}
	for v := range m {
		tmp := target - v
		if count, ok := m[tmp]; ok {
			if tmp != v {
				return []int{target - v, v}
			}
			if tmp == v && count > 1 {
				return []int{target - v, v}
			}
		}
	}
	return nil
}
