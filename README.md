# leetcode&lintcode题解 (Go语言)

**每道题一个文件，但是题解中可能调用其它题的函数。**

## 思路

## 文件目录

```

array
    ├── leet_jian_53.go // 剑指 Offer 53 - I. 在排序数组中查找数字 I
    ├── leet_jianzhi_03.go // 剑指 Offer 03. 数组中重复的数
    ├── leet_jianzhi_04.go // 剑指 Offer 04. 二维数组中的查找
    ├── leet_jianzhi_11.go // 剑指 Offer 11. 旋转数组的最小数字
    └── leet_jianzhi_53_v2.go // 剑指 Offer 53 - II. 0～n-1中缺失的数字

dp
    ├── leet_jianzhi_10_v1.go // 剑指 Offer 10- I. 斐波那契数列
    ├── leet_jianzhi_10_v2.go // 剑指 Offer 10- II. 青蛙跳台阶问题
    └── leet_jianzhi_63.go // 剑指 Offer 63. 股票的最大利润
    ├── leet_jianzhi_42.go // 剑指 Offer 42. 连续子数组的最大和
    ├── leet_jianzhi_46.go // 剑指 Offer 46. 把数字翻译成字符串
    ├── leet_jianzhi_47.go // 剑指 Offer 47. 礼物的最大价值


doubly_linked_list 
    ├── leet_146.go // LRU 缓存机制
    └── leet_1472.go // 设计浏览器历史记录
    ├── leet_jianzhi_21.go // 剑指 Offer 21. 调整数组顺序使奇数位于偶数前面
    ├── leet_jianzhi_25.go 剑指 Offer 25. 合并两个排序的链表
    └── leet_jianzhi_52.go 剑指 Offer 52. 两个链表的第一个公共节点
    ├── leet_jianzhi_57.go // 剑指 Offer 57. 和为s的两个数字
    └── leet_jianzhi_58.go // 剑指 Offer 58 - I. 翻转单词顺序

    
linked_list
    ├── leet_83.go // 删除排序链表中的重复元素
    ├── leet_jianzhi_06.go // 从尾到头打印链表
    ├── leet_jianzhi_18.go // 剑指 Offer 18. 删除链表的节点
    ├── leet_jianzhi_22.go // 剑指 Offer 22. 链表中倒数第k个节点
    └── leet_jianzhi_24.go // 反转链表
    └── leet_jianzhi_35.go // 复杂链表的复


queue
    ├── leet_387.go // 字符串中的第一个唯一字符 & 剑指 Offer 50. 第一个只出现一次的字符
    ├── leet_jianzhi_09.go 用两个栈实现队列
    └── leet_jianzhi_50.go // 剑指 Offer 50. 第一个只出现一次的字符
search
    ├── leet_jianzhi_12.go // 剑指 Offer 12. 矩阵中的路径
    └── leet_jianzhi_13.go // 剑指 Offer 13. 机器人的运动范围

sliding_window
    ├── leet_1876.go // 长度为三且各字符不同的子字符串
    ├── leet_217.go // 存在重复元素
    ├── leet_219.go // 存在重复元素II
    ├── leet_3.go // 无重复字符的最长子串
    └── leet_643.go // 子数组最大平均数 I
    ├── leet_718.go // 最长重复子数组
sort
    ├── leet_jianzhi_40.go // 剑指 Offer 40. 最小的k个数
    ├── leet_jianzhi_41.go // 剑指 Offer 41. 数据流中的中位数
    ├── leet_jianzhi_45.go // 剑指 Offer 45. 把数组排成最小的数
    └── leet_jianzhi_61.go // 剑指 Offer 61. 扑克牌中的顺子

stack
    └── leet_jianzhi_30.go // 包含min函数的栈
string
    ├── leet_offer_05.go // 剑指 Offer 05. 替换空格
    └── leet_offer_58.go // 剑指 Offer 58 - II. 左旋转字符串

tree
    ├── 94.go
    ├── leet_101.go // 对称二叉树 （层次遍历）
    ├── leet_226.go // 翻转二叉树
    ├── leet_429.go // N叉树的层序遍历
    ├── leet_637.go // 二叉树的层平均值
    ├── leet_jianzhi_26.go // 剑指 Offer 26. 树的子结构
    ├── leet_jianzhi_27.go // 剑指 Offer 27. 二叉树的镜像
    ├── leet_jianzhi_28.go // 剑指 Offer 28. 对称的二叉树 (使用递归)
    ├── leet_jianzhi_32_v1.go // 剑指 Offer 32 - I. 从上到下打印二叉树
    ├── leet_jianzhi_32_v3.go // 剑指 Offer 32 - III. 从上到下打印二叉树 III
    ├── leet_jianzhi_34.go // 剑指 Offer 34. 二叉树中和为某一值的路径
    ├── leet_jianzhi_54.go // 剑指 Offer 54. 二叉搜索树的第k大节点
    ├── lint_1744.go // 递增顺序查找树
    ├── lint_1746.go // 二叉搜索树结点最小距离
    ├── lint_66.go // 二叉树前序遍历
    ├── lint_68.go //  二叉树后序遍历
    ├── lint_69.go //  二叉树层次遍历 & 剑指 Offer 32 - II. 从上到下打印二叉树 II
    └── lint_97.go // 二叉树的最大深度
    
```