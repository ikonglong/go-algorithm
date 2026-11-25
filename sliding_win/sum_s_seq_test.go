package slidingwin

// https://leetcode.cn/problems/he-wei-sde-lian-xu-zheng-shu-xu-lie-lcof/
// 解读题目：给定一个整数 s，从 1..s 这个整数序列中，找出所有和为 s 的连续整数序列。
// 【一个序列内，至少包含两个元素】，按升序排列。结果序列之间，按第一个元素升序排列。
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

func fileCombination(s int) [][]int {
	result := make([][]int, 0)
	// Bug: 这么初始化循环一次也不会运行
	// small, big := 1, 1
	// sum := 1
	small, big := 1, 2
	sum := small + big
	// 为什么 small == big 时退出循环？
	// 当 small == big 时，这个结果只可能是前一迭代缩小窗口的结果，不可能是扩大的结果。
	// 也就是说，前一迭代因为 sum > s，所以缩小窗口。而当前迭代只能扩大窗口，但只要扩大
	// 窗口，结果必然大于 s；并且自此开始，窗口每次扩缩也总是大于 s
	//
	// 假设当前 small 位于 s/2 位置，此时 big 最小是 small+1，那么：
	// 若 s 为奇数，则 small + big = s
	// 若 s 为偶数，则 small + big > s
	// 因此，自此之后 small + big 必定大于 s
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

			// 移动任何一个、移动两个都行，只要打破相等状态即可。不管如何选择，后续迭代都会继续选择恰当的移动。
			// 例如，若只移动 big，那么 sum > s，下次循环就 small++；若只移动 small，同理
			// 因此选择两个都移动，因为效率好一点
			sum -= small
			small++
			big++
			sum += big
		}
	}
	return result
}
