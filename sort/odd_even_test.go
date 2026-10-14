package sort_

// https://leetcode.cn/problems/diao-zheng-shu-zu-shun-xu-shi-qi-shu-wei-yu-ou-shu-qian-mian-lcof/description/
// 算法思路：
// |------------|
// lo           hi
// [0, lo) 是奇数区间，[lo, hi] 是偶数区间，(hi, n-1] 是未处理区间

func trainingPlan(actions []int) []int {
	n := len(actions)
	lo, hi := 0, n-1
	for lo < hi {
		// 移动 low 指针，直到找到一个偶数，用来跟右侧的奇数交换
		if actions[lo]&1 == 1 { // 是奇数
			lo++
			continue
		}
		// 移动 high 指针，直到找到一个奇数，用来跟左侧的偶数交换
		if actions[hi]&1 == 0 { // 是偶数
			hi--
			continue
		}

		actions[lo], actions[hi] = actions[hi], actions[lo]
	}
	return actions
}

func trainingPlan2(actions []int) []int {
	n := len(actions)
	lo, hi := 0, n-1
	for lo < hi {
		// 移动 low 指针，直到找到一个偶数，用来跟右侧的奇数交换
		// 若全部是奇数，low 可能越过右边界。high 是 low 的右边界守卫
		for lo < hi && actions[lo]&1 == 1 {
			lo++
		}
		// 移动 high 指针，直到找到一个奇数，用来跟左侧的偶数交换
		// 若全部是偶数, high 可能越过左边界。low 是 high 的左边界守卫
		for lo < hi && actions[hi]&1 == 0 {
			hi--
		}
		actions[lo], actions[hi] = actions[hi], actions[lo]
	}
	return actions
}
