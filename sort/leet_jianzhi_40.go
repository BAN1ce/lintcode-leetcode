package sort

import "sort"

// 剑指 Offer 40. 最小的k个数
// https://leetcode-cn.com/problems/zui-xiao-de-kge-shu-lcof/
func getLeastNumbers(arr []int, k int) []int {

	sort.Slice(arr, func(i, j int) bool {
		return  arr[i] > arr[j]
	})

	return arr[0:k]
}