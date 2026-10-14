package binary_search

// https://leetcode.cn/problems/find-minimum-in-rotated-sorted-array/
// 注意，若旋转 7 次，则可以得到 [0,1,2,4,5,6,7]；数组中的值互不相同
// 算法思路：
// 1. 如果输入非升序（也就是包含旋转断点），则旋转断点在哪边，最小值就在哪边
// 2. 如果输入是升序的，或者说旋转后是升序的，则不包含旋转断点，或者说，旋转
//    断点跨越左右两边，即尾元素和首元素组成旋转断点
// 如果旋转断点明确在右边，那么最小值一定在右边；否则，最小值就在左边，不管输入是哪种情形？

func findMin(nums []int) int {
	n := len(nums)
	lo, hi := 0, n-1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if lo == hi {
			return nums[lo] // ?
		}
		// 注意旋转数组跟山形峰值数组不一样，千万别搞混淆了
		// 例如：旋转数组 [4,5,6,7,0,1,2] 中 7->0 这样的下坡只有一个，且只有一对相邻元素是降序。
		// 但峰值数组 [1,2,4,7,5,3,1] 下坡比较长，多对相邻元素都是降序

		// mid-1>=0 => mid>=1 => mid>0 => 想到看看 mid==0 是什么情况
		// 升序数组本质上也是循环有序数组的一个特例。因此，可以认为最后一个元素是首元素的前驱。
		// 若 mid=0，其前驱就是 hi，对于升序数组来说，同样满足若前驱大于 num[mid]，num[mid]就是最小值
		// 这样就可以把下面条件中的两种情形统一并简化为：nums[(mid-1+n)%n] > nums[mid]
		if (mid > 0 && nums[mid-1] > nums[mid]) || (mid == 0 && nums[hi] > nums[mid]) {
			return nums[mid]
		} else if nums[mid] > nums[hi] {
			lo = mid + 1
		} else {
			hi = mid - 1
		}

		// 注意，只有根据右侧才能正确判断最小值在哪边
		// 若 nums[mid] > nums[hi]，则旋转断点在 mid 右侧，最小值在右边。
		// 若 nums[mid] <= nums[hi]，则左半边是升序或包含旋转断点，最小值在左边。

		// 如果 nums[lo] > nums[mid]，那么最小值一定在 mid 左边。但是 nums[lo]<=nums[mid]
		// 并不能推出最小值在 mid 右边，因为右边可能是升序的，也可能包含旋转断点；仅当包含旋转断点
		// 时，最小值才在右边。例如：[1,2,3,4,5]，mid=3，最小值就不在右边。因此下面的写法不行

		// else if nums[lo] > nums[mid] { // 表明 [lo, mid] 非升序（即循环有序），那么 [mid+1, hi] 升序
		// 	hi = mid - 1
		// } else { // nums[lo]<=nums[mid] 表明 [lo, mid] 升序，但不能退出 [mid+1, hi] 循环有序或叫包含旋转断点
		// 	lo = mid + 1
		// }
	}
	return -1
}
