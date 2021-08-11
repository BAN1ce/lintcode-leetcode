package array

// 剑指 Offer 53 - II. 0～n-1中缺失的数字
// https://leetcode-cn.com/problems/que-shi-de-shu-zi-lcof/
func missingNumber(nums []int) int {

	l := len(nums)
	for i := 0; i < l; i++ {
		if nums[i] != i {
			return i
		}
	}
	return len(nums)-1
}
