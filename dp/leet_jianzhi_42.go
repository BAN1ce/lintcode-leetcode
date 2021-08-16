package dp

// 剑指 Offer 42. 连续子数组的最大和
// https://leetcode-cn.com/problems/lian-xu-zi-shu-zu-de-zui-da-he-lcof/
func maxSubArray(nums []int) int {

	l := len(nums)
	if l == 0 {
		return 0
	}
	maximum := nums[0]

	m := make([]int, l+1)
	m[0] = nums[0]
	for i := 1; i < l; i++ {
		m[i] = max(nums[i], m[i-1]+nums[i])
		if m[i] >maximum{
			maximum=m[i]
		}
	}

	return maximum
}

func max(a, b int) int {

	if a > b {
		return a
	}

	return b
}
