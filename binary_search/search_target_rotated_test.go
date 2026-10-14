package binary_search

// https://leetcode.cn/problems/search-in-rotated-sorted-array/
// 注意，数组中的值互不相同
// 算法思路：
// 一般二分查找处理的输入只有一种模式/样式，因此循环体中的 if else if 处理的是
// 一种样式的输入中 nums[mid] 跟 target 的不同关系。但是这里二分查找处理的输入有三种模式：
//  1. 旋转断点在左边
//  2. 旋转断点在右边
//  3. 旋转断点跨越左右两边，即尾元素和首元素组成旋转断点。换句话说，普通有序数组是旋转数组的
//     一种特例，即尾元素是首元素的前驱元素，这样就构成环了。【这种输入当前题目貌似没有】
// 这样，二分查找循环体中的第一层 if else if 得先判断输入是哪一种模式，再根据模式特征决定如何
// 判断 target 在哪一边。虽然 target 可能在包含旋转断点的那边，但是通过不断地二分，最终 target
// 一定会落在一个有序区间，使用二分查找就可以找到它了

func searchTargetRotated(nums []int, target int) int {
	n := len(nums)
	lo, hi := 0, n-1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if nums[mid] == target {
			return mid
		} else if nums[lo] <= nums[mid] {
			// 若 [lo, mid] 有序，则 [mid+1, hi] 循环有序，target 可能在任何一边
			if nums[lo] <= target && target < nums[mid] { // target < nums[mid]
				hi = mid - 1
			} else { // nums[mid] < target
				lo = mid + 1
			}
		} else {
			// 若 [lo, mid] 循环有序，则意味着 [mid+1, hi] 有序，target 可能在任何一边
			// if nums[mid] < target && target <= nums[hi] {
			if nums[mid+1] <= target && target <= nums[hi] { // nums[mid] < target
				lo = mid + 1
			} else { // targe < nums[mid]
				hi = mid - 1
			}
		}
	}
	return -1
}

// 从上面的实现版本简化得到该版本，原本我认为该版本更简洁，且完全 OK。
// 但这个版本测试失败，跟我的理解和预期完全不符，有待搞清楚，我的理解、认识错在哪里？
//
// 首先，用例子说明下面这个实现为什么不行。例如：[4,5,6,7,0,1,2]，target=0，旋转使
// 得二分查找的一般规则在这里不成立了。
//
// 其次，一般二分查找处理的输入只有一种模式/样式，因此循环体中的 if else if 处理的是
// 一种样式的输入中 nums[mid] 跟 target 的不同关系。但是这里二分查找处理的输入有三种
// 模式，循环体中的第一层 if else if 得先判断输入是哪一种模式，再根据模式特征决定如何
// 判断 target 在哪一边。

func searchRotated2(nums []int, target int) int {
	n := len(nums)
	lo, hi := 0, n-1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return -1
}
