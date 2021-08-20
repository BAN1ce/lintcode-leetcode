package doubly_linked_list

// 剑指 Offer 21. 调整数组顺序使奇数位于偶数前面
// https://leetcode-cn.com/problems/diao-zheng-shu-zu-shun-xu-shi-qi-shu-wei-yu-ou-shu-qian-mian-lcof/
func exchange(nums []int) []int {

	left, right := 0, len(nums)-1

	for i, j := left, right; i < j; {
		if nums[i]%2 == 0 && nums[j]%2 == 1 {
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j--
		} else if nums[i]%2 == 0 {
			j--
		}else {
			i++
		}
	}

	return nums
}
