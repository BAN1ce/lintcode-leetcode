package dp

// 剑指 Offer 10- I. 斐波那契数列
// https://leetcode-cn.com/problems/fei-bo-na-qi-shu-lie-lcof/

func fib(n int) int {

	m:= make(map[int]int)

	m[0]=0
	m[1]=1
	for i:=2;i<=n;i++{
		m[i] = (m[i-1]+m[i-2])%(1e9+7)
	}

	return m[n]
}