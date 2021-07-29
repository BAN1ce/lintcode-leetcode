package sliding_window

// 219 存在重复元素 II
// https://leetcode-cn.com/problems/contains-duplicate-ii/
func containsNearbyDuplicate(nums []int, k int) bool {

	m := make(map[int][]int)

	for i, v := range nums {

		if m[v] != nil {
			if i-m[v][len(m[v])-1] <= k {
				return true
			}
		}
		m[v] = append(m[v], i)

	}

	return false

}
