package backtrack

// https://leetcode.cn/problems/permutations-ii/
// 输入数组包含重复数据，导致排列中有重复答案

func permuteUnique_(nums []int) [][]int {
	result := make([][]int, 0)
	opts := countFreq_(nums)
	n := len(opts)
	// bug
	// n := len(nums)

	var backtrack func(k int, path []int)
	backtrack = func(k int, path []int) {
		if k == n {
			ans := make([]int, n)
			copy(ans, path)
			result = append(result, ans)
			return
		}
		// 注意，`for i, opt := range opts` 中 opt 是元素的副本，不是元素本身。
		// 因此，修改它并不会修改元素本身
		for i := range opts {
			// 前面阶段已经用完的数就不能再选择了
			if opts[i].count == 0 {
				continue
			}
			// 尝试一个选择
			path = append(path, opts[i].num)
			// 注意，必须通过下标访问元素本身，
			// 通过 for range 中的循环变量无法修改元素本身
			opts[i].count--
			backtrack(k+1, path)
			// 恢复状态
			opts[i].count++
			path = path[0 : len(path)-1]
		}
	}

	backtrack(0, nil)
	return result
}

func countFreq_(nums []int) []NumFreq_ {
	aMap := make(map[int]int)
	for _, v := range nums {
		aMap[v] = aMap[v] + 1
	}

	numFreq := make([]NumFreq_, 0, len(nums))
	for k, v := range aMap {
		numFreq = append(numFreq, NumFreq_{num: k, count: v})
	}
	return numFreq
}

type NumFreq_ struct {
	num   int
	count int
}
