package twopointer

// 区间指针
func moveZeroes(nums []int) {
	// q 指向未处理区的头端，p 指向非0区的尾端
	// 创建指针 q 就是用来遍历数组的
	p, q := -1, 0

	// 注意：为什么这里没有显式确保不会发生 p 越界？
	// 因为初始时 p 比 q 小 1（即 q 在前，p 在后），且两个指针要么一起移动，要么只移动 q
	// 所以只要确保 q 不越界，p 自然不会越界
	for q < len(nums) {
		v := nums[q]
		if v == 0 {
			// Q 右移一步，相当于 0 区尾端右移一步
			q++
		} else if v != 0 {
			// 交换 p+1 和 q 位置的值，p 和 q 都右移一步
			// 即非0区尾端向右扩一步；0区头、尾端分别向右缩一步、扩一步
			tmp := nums[p+1]
			nums[p+1] = nums[q]
			nums[q] = tmp
			p++
			q++
		}
	}
}
