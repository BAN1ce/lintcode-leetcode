package dp

// 剑指 Offer 63. 股票的最大利润
// https://leetcode-cn.com/problems/gu-piao-de-zui-da-li-run-lcof/
func maxProfit(prices []int) int {

	if len(prices) <=1{
		return 0
	}
	var min int
	var max int
	if prices[0] >prices[1]{
		min = prices[1]
	}else {
		min = prices[0]
		max = prices[1]-prices[0]
	}

	for i:=2;i<len(prices);i++{
		if prices[i] > min {
			if tmp:= prices[i]-min;tmp>max{
				max =tmp
			}
		}else {
			min = prices[i]
		}
	}

	return max

}
