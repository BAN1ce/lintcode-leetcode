package array

// 剑指 Offer 11. 旋转数组的最小数字
// https://leetcode-cn.com/problems/xuan-zhuan-shu-zu-de-zui-xiao-shu-zi-lcof/
func minArray(numbers []int) int {


	min := numbers[0]


	for _, v:=range numbers{
		if v < min {
			min =v
		}
	}
	return min
}
