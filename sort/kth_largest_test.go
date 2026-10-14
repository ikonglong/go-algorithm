package sort_

// https://leetcode.cn/problems/kth-largest-element-in-an-array/

func findKthLargest(nums []int, k int) int {
	if len(nums) < k {
		return -1
	}
	return doFindKthLargest(nums, 0, len(nums)-1, k)
}

func doFindKthLargest(nums []int, lo, hi, k int) int {
	p := partition3(nums, lo, hi)
	if k == p-lo+1 {
		return nums[p]
	} else if k < p-lo+1 {
		return doFindKthLargest(nums, lo, p-1, k) // todo
	} else {
		return doFindKthLargest(nums, p+1, hi, k-(p-lo+1))
	}
}

// partition 循环结束时只知道指针到了分界附近，因此还要判断；
// partition2 和 partition3 的指针约定已经保证了某个指针（分别是 i 和 j）就是枢轴（pivot）的落点。

// 这里 partition 是动词，意思是分割
// 从给定子数组中选择一个元素作为分区点，调整元素位置，
// 使得所有大于等于它的都在其左边，所有小于等于的都在其右边，返回分区点位置
func partition(nums []int, lo, hi int) int {
	i, j := lo, hi-1
	// p(ivot) 枢轴元素
	p := nums[hi]

	// 指针停在“相遇点”，相遇点不一定是枢轴位置
	// 这个版本的循环条件是 i < j。循环结束时，i 和 j 可能相遇，也可能交错。
	// 此时 nums[i] 可能在枢轴左边，也可能在右边，所以不能直接把枢轴换到 i：
	// - 如果 nums[i] <= p，枢轴应该放到 i；
	// - 否则，枢轴应该放到 i+1。
	// 所以循环后要用 if/else 判断。通俗地说，指针停下的位置只是“交界附近”，
	// 还得看这个位置上的数，才能确定枢轴该落在哪一格。
	for i < j {
		// 找小于等于分区点的
		for i < j && nums[i] > p {
			i++
		}
		// 找大于等于分区点的
		for i < j && nums[j] < p {
			j--
		}
		if i < j { // 这个判断容易写漏
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j--
		}
	}
	// 循环结束后，找第一个小于等于分区元素的，交换
	if nums[i] <= p {
		nums[i], nums[hi] = nums[hi], nums[i]
		return i
	} else {
		nums[i+1], nums[hi] = nums[hi], nums[i+1]
		return i + 1
	}
}

// 左边放大于等于分区点的元素，右边放小于等于分区点的元素
func partition2(nums []int, lo, hi int) int {
	i, j := lo, hi-1
	p := nums[hi]

	// i 本身就是分界位置
	// 这个版本一直循环到 i > j。扫描和交换过程中，左边逐渐放好 >= p 的元素，
	// 右边逐渐放好 <= p 的元素。指针交错时，i 已经指向右侧分区的起点，也就是
	// 枢轴应该插入的位置。所以可以直接交换，不用再判断 nums[i] 的大小。
	for i <= j {
		// 跳过左侧已经满足条件的元素
		for i <= j && nums[i] > p {
			i++
		}

		// 跳过右侧已经满足条件的元素
		for i <= j && nums[j] < p {
			j--
		}

		// 交换放错边的元素，然后让两个指针继续向中间移动
		if i <= j {
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j--
		}
	}

	// 循环结束时，i 是右侧分区的起始位置
	nums[i], nums[hi] = nums[hi], nums[i]
	return i
}

// i == j 通常有两类情况：
//
// 1. 两边扫描在一个等于枢轴 p 的元素上相遇。
// 内部位置上，左扫描停在 <= p 的元素，右扫描停在 >= p 的元
// 素；若停在同一个位置，该元素就必须等于 p。
// 例如：nums=[2, 5, 2, 1]，这里 p = 2，两边扫描会在下标 2 相遇。
//
// 2. 左扫描到达 hi，碰上代码里的边界退出。
// 左扫描即使遇到 nums[hi] > p，也会因 i == hi 停下；右扫描也可能
// 停在同一个 hi。【注意，此时 nums[hi] 并不等于枢轴，跟上一情形不同】
// 例如：nums=[1, 2, 3]，这里 p = 1，左右指针都会停在下标 2，于是 i == j。
//
// 总结：相遇点可能是等于枢轴的元素，也可能是 hi 边界处的元素。相遇时所在元素，
// 要么等于枢轴，要么大于枢轴。大于枢轴时，只可能是 nums[hi]：从左向右的扫描到
// 达 hi 时会被边界保护强制停下；右扫描停在 hi 则说明 nums[hi] >= p，因此
// 它可能大于 p，但不会小于 p。

func partition3(nums []int, lo, hi int) int {
	if lo >= hi {
		return lo
	}

	i, j := lo, hi+1
	p := nums[lo]

	// j 是枢轴最终位置
	// 这个版本的枢轴一开始放在 lo，左右扫描分别跳过已经在正确一侧的元素。
	// 扫描交错后，按这套扫描规则，j 正好落在枢轴应该放的位置，于是直接交换，也不需要额外判断。
	for {
		// 左扫描：跳过 > p 的元素，停在 <= p 的元素或 hi
		// for i++; i < hi && nums[i] > p; i++ {
		// }
		for i++; nums[i] > p; i++ {
			if i == hi {
				break
			}
		}

		// 右扫描：跳过 < p 的元素，停在 >= p 的元素或 lo
		// for j--; j > lo && nums[j] < p; j-- {
		// }
		for j--; nums[j] < p; j-- {
			if j == lo {
				break
			}
		}

		if i >= j {
			break
		}

		nums[i], nums[j] = nums[j], nums[i]
	}

	nums[lo], nums[j] = nums[j], nums[lo]
	return j
}
