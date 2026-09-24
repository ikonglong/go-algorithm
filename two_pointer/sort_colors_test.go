package slidingwin

import (
	"fmt"
	"testing"
)

// https://leetcode.cn/problems/sort-colors/description/

// 区间指针
// |-------|-------|-----------------|
//    红       白     未处理区又是蓝区
//         p       q
// 对于未处理区，将其中的红蓝都移出去后，它就是蓝区
// 初始时红区、白区都不存在，因此将 p 和 q 都初始化为 0，
// 表示[x, p) 是红区，[p, q) 是白区，[q, n)是未处理区又是蓝区。
// p、q 分别是红区、白区的右侧开边界（不包含），初始化为 0 贴切地表达了不存在，挺好的建模

// 算法思路：
// 1. 创建一个指针 ptr，表示头部的红色范围，但头部范围不包含这个指针，即为开区间
// 2. 初始时，ptr 指向 0，即头部范围为 [0, 0)，空范围
// 3. 从 i=0 开始遍历数组。如果nums[i]为 0，就跟 nums[ptr] 交换
// 4. 从 ptr 开始再次遍历数组，如果ums[i]为 1，就跟 nums[ptr ]交换
func sortColors(nums []int) {
	if len(nums) == 0 || len(nums) == 1 {
		return
	}

	ptr := 0
	// 为什么 i 不能从 1 开始？
	// 首先 ptr 是头部范围的右边界，但不包含，因此未处理部分应该从 ptr 开始。
	// 其次，从 i=1 开始埋有 bug。如果从 1 开始，nums[0] 是 0 时，
	// ptr 虽然位置正确，但没有被“计入”0区间，因为 ptr 本该指向 1，但它仍是 0。
	// 从 i=0 开始时，if nums[i] == 0 {...} 会将 ptr 指向正确的位置。
	// 例如：nums=[0,1,2]，第一个循环从 i=1 开始，未发现新的 0，循环结束后 ptr
	// 仍是 0。第二个循环从 ptr 开始找 1，找到第一个 nums[i]=1 时跟 nums[ptr] 交换，
	// 就把 1 换到 nums[ptr] 了，但此时 ptr=0 且 nums[ptr]=0
	for i := range nums {
		if nums[i] == 0 {
			nums[ptr], nums[i] = nums[i], nums[ptr]
			ptr++
		}
	}
	for i := ptr; i < len(nums); i++ {
		if nums[i] == 1 {
			nums[ptr], nums[i] = nums[i], nums[ptr]
			ptr++
		}
	}
}

// 算法思路：
// 这个思路的巧妙之处在于，即使初始时0区、1区都还不存在，使用其尾元素下一个位置也可以表示它们，
// 并且这个建模既能自然地表示还不存在的分区，也能自然地表示已存在的分区
//
//  1. 创建两个指针 p0、p1，分别表示 0 区、1 区尾元素下一个位置。
//     初始时还不存在 0 区、1 区，因此将 P0、P1 置为 0，表示这两个区还不存在。
//     注意初始时，零区和一区的边界还混沌未分，叠在一起
//  2. 遍历 nums
//       若当前元素等于 0
// 	       交换 nums[p0] 和 nums[i]。
//
//         若 p0 < p1，则表明 0 区和 1 区的边界已经分清了。
//         这意味着 nums[p0] 是 0 区尾元素的下一个位置，也就是 1 区的头。
//         而前一次交换将 nums[p0] 换到了 nums[i]
//		     要交换 nums[i]、nums[p1]，这样就把起初的 nums[p0] 放到了 nums[p1]
//       若当前元素等于 1，则交换 nums[p1] 和 nums[i]
//

func sortColors2(nums []int) {
	if len(nums) == 0 || len(nums) == 1 {
		return
	}

	p0, p1 := 0, 0
	for i := range nums {
		if nums[i] == 0 {
			nums[p0], nums[i] = nums[i], nums[p0]
			if p0 < p1 {
				nums[i], nums[p1] = nums[p1], nums[i]
			}
			p0++
			// 若 nums=[0,2,1]，第一次执行此分支时，p0++将nums[0]归入0区。
			// p1仍指向0。0区一旦存在，虽然1区还不存在，p1也应该一起移动到下
			// 一个可表示1区不存在的位置，即0区尾元素下一个位置。此时p0,p1仍重叠
			p1++
		} else if nums[i] == 1 {
			nums[p1], nums[i] = nums[i], nums[p1]
			p1++
		}
	}
}

// 算法思路：
// 跟方法二类似，也需要两个指针 p0, p2，分别指向 0 区右边界（不包含）、2 区左边界（不包含）。
// p2 指向最后一个未处理的元素。这样它们既能自然地表示还不存在的区间，也能自然地表示已存在的区间。
//
//  1. 初始时，p0、p2 分别为 0、n-1
//
//  2. 从 i=0 开始遍历数组。循环条件是 i<=p2。遍历过程中，
//     若 nums[i] 为 0，跟 nums[p0] 交换
//     若 nums[i] 为 2，跟 nums[p2] 交换
//
//  3. 这样做是正确的吗？可以发现，对于第二种情况，将 nums[i] 与 nums[p2] 交换之后，
//     新的 nums[i] 可能仍是 2，也可能是 0/1。如果此时结束交换，开始遍历下一个元素 nums[i+1]，
//     就不会再考虑 nums[i] 了，这样我们就会得到错误的答案。
//
//     因此，当找到 2 时，需要不断地将其与 nums[p2] 交换，直到新的 nums[i] 不为 2。此时，如果
//     nums[i] 为 0，那么对应着第一种情况；如果 nums[i] 为 1，那么就不需要进行任何后续的操作。
//
// 注意，这个算法不是交换 nums[p0] 和 nums[p2]，而是交换 nums[i] 跟 nums[p0|p2]。
// 交换 nums[p0] 和 nums[p2] 的思路一次遍历搞不定
func sortColors3(nums []int) {
	if len(nums) == 0 || len(nums) == 1 {
		return
	}

	p0, p2 := 0, len(nums)-1
	for i := 0; i <= p2; i++ {
		for ; i <= p2 && nums[i] == 2; p2-- {
			nums[p2], nums[i] = nums[i], nums[p2]
		}
		// 此时 nums[i] 可能是 0 或 1
		if nums[i] == 0 {
			nums[p0], nums[i] = nums[i], nums[p0]
			p0++
		}
		// 如果是 1，什么都不做
	}
}

func TestSortColor(t *testing.T) {
	nums := []int{2, 0, 2, 1, 1, 0}
	func() {
		sortColors3(nums)
		fmt.Printf("nums:%v\n", nums)
	}()
}
