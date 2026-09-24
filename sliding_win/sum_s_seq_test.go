package slidingwin

// https://leetcode.cn/problems/he-wei-sde-lian-xu-zheng-shu-xu-lie-lcof/
// 解读题目：给定一个整数 s，从 1..s 这个整数序列中，找出所有和为 s 的连续整数序列。
// 一个序列内，至少包含两个元素，按升序排列。结果序列之间，按第一个元素升序排列。
//
// 一个有点反直觉的点：在一个升序的连续整数序列上放一个滑动窗口，关注窗口中序列的和，
// 窗口左边前移时，和减小；窗口右边前移时，和增大
//
// 算法思路：
//
// 1. 用 small, big 两个指针创建一个滑动窗口，初始时分别为 1, 2
// 2. 循环，条件为 small 小于 big
//      计算从 small 到 big 的和，即为 sum
//      如果 sum < s，big++, sum=sum+big
//      如果 sum > s，sum=sum-small, small++，
//        small++通过移动窗口左边把一个较小的数减掉。
//        big--也能让和减小，但减小的幅度较大。
//      如果 sum == s，将序列添加到结果中，small++, big++

func fileCombination(s int) [][]int {
	result := make([][]int, 0)
	small, big := 1, 2
	sum := small + big
	// 何时退出循环？small == big 时。为什么？
	// 窗口在一条越往右数越大的数轴上寻找目标和。最窄的合法窗口：[small, big] = [x, x+1]。
	// 当最窄窗口和大于 s 时，只能丢弃 x，于是 small 等于 big，窗口缩成了一个数。按照规则，
	// 要右移 big，但继续向右寻找已经没有意义，因为右移所得结果也必定大于 s。
	//
	// 对于当前问题，假定 big=small+1。small 移动至 1/2*s 位置时：
	// 若 s 为偶数，则 small + big = s
	// 若 s 为奇数，则 small + big > s。自此之后，small + big 也必定大于 s
	for small < big {
		if sum > s {
			sum = sum - small
			small++
		} else if sum < s {
			big++
			sum += big
		} else {
			seq := make([]int, 0, big-small+1)
			for i := small; i <= big; i++ {
				seq = append(seq, i)
			}
			result = append(result, seq)

			// 这里只移动 small 或 big，还是两者都移动？无所谓，都行
			// 若只移动 big，sum > s，下次循环就 small++；若只移动 small，同理
			// 因此选择两个都移动，因为效率好一点
			sum -= small
			small++
			big++
			sum += big
		}
	}
	return result
}
