package doubly_linked_list

import "strings"

// 剑指 Offer 58 - I. 翻转单词顺序
// https://leetcode-cn.com/problems/fan-zhuan-dan-ci-shun-xu-lcof/
func reverseWords(s string) string {

	words := strings.Split(s, " ")

	str := ""


	for i:=len(words)-1;i>=0;i--{
		if words[i] == ""{
			continue
		}
		str+= strings.Trim(words[i]," ")+" "
	}
	return strings.Trim(str," ")
}
