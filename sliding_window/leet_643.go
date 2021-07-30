package sliding_window

// 643. 子数组最大平均数 I
// https://leetcode-cn.com/problems/maximum-average-subarray-i/
func findMaxAverage(nums []int, k int) float64 {

	left := 0
	right := k

	sum := 0
	for i := 0; i < k; i++ {
		sum += nums[i]
	}
	max := float64(sum) / float64(k)

	for ; right < (len(nums)); {

		var tmp float64
		if tmp, sum = avg(nums[left], nums[right], sum, k); tmp > max {
			max = tmp
		}
		left += 1
		right += 1

	}

	return max
}

// 必须使用减左边的数加右边的数，否则会超时
func avg(leftValue, rightValue, sum, k int) (float64, int) {

	sum = sum - leftValue + rightValue
	return float64(sum) / float64(k), sum

}
