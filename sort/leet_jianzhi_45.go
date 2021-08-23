package sort

import (
	"fmt"
	"sort"
	"strconv"
)

type nums []int

func (n nums) Len() int {
	return len(n)
}

func (n nums) Less(i, j int) bool {
	return fmt.Sprintf("%d%d", n[i], n[j]) < fmt.Sprintf("%d%d", n[j], n[i])
}

func (n nums) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

// 剑指 Offer 45. 把数组排成最小的数
// https://leetcode-cn.com/problems/ba-shu-zu-pai-cheng-zui-xiao-de-shu-lcof/
func minNumber(n []int) string {

	sort.Sort(nums(n))

	str := ""

	for _, v := range n {
		str += strconv.Itoa(v)
	}
	return str
}
