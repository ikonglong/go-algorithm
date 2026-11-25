package slidingwin

// https://leetcode.cn/problems/he-wei-sde-lian-xu-zheng-shu-xu-lie-lcof/
// 解读题目：给定一个整数 s，从 1..s 这个整数序列中，找出所有和为 s 的连续整数序列。
// 一个序列内，至少包含两个元素，按升序排列。结果序列之间，按第一个元素升序排列。
//
// 一个有点反直觉的点：在一个升序的连续整数序列上放一个滑动窗口，关注窗口中序列的和。
// 窗口左端右移时，窗口缩小，因此和减小；窗口右端右移时，窗口扩大，和增大
//
// 算法思路：
// 1. 用 small, big 两个指针创建一个滑动窗口，初始时分别为 1, 2
// 2. 循环，条件为 small 小于 big
//      计算窗口内数的和，记为 sum
//      如果 sum < s，big++, sum=sum+big
//      如果 sum > s，sum=sum-small, small++，
//        small++通过移动窗口左端把一个较小的数减掉。
//        big--也能让和减小，但减小的幅度较大。
//      如果 sum == s，将序列添加到结果中，small++, big++

func fileCombination_(s int) [][]int {
	result := make([][]int, 0)
	// Bug: 这么初始化循环一次也不会运行
	// small, big := 1, 1
	// sum := 1
	small, big := 1, 2
	sum := small + big
	for small < big {
		// 该死的！很低级的错误呀
		// sum := small + big
		if sum > s {
			sum -= small
			small++
		} else if sum < s {
			big++
			sum += big
		} else { // sum == s
			seq := make([]int, 0, big-small+1)
			for i := small; i <= big; i++ {
				seq = append(seq, i)
			}
			result = append(result, seq)

			// 移动任何一个、两个都移动都行，只要打破相等状态即可。
			// 后续迭代自然会选择合适的移动
			sum -= small
			small++
			big++
			sum += big
		}
	}
	return result
}
