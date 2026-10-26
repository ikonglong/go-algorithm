package hashmap

import "slices"

// https://leetcode.cn/problems/3sum/
// 注意：
// 1. 一个元素不能使用两次
// 2. 假设 a+b+c=0，(a,b,c), (b,a,c), (c,a,b), ... 是同一个解，只能出现一次
// 3. 因不同元素的值相同，使得满足要求的不同元素的组合看起来是一样的，也只能出现一个
//    对于这类重复，只要两个值一样，剩余那个肯定一样。因此只需要确保 (a, b) 跟已有
//    解不同，就不会有重复

// 在确定 a 时，若 a 与前一个元素重复，就直接跳过 (continue)。但是当确定 a 时，b 还
// 未确定呢，怎么就判定 (a, b) 这个组合就一定会重复，可以直接跳过？
// 这里跳过重复的 a，不是因为此时已经知道 (a, b) 会重复，而是因为这个重复的 a 后面能尝
// 试的组合，已经被前一个相同值的 a 涵盖了。因为 a 为前一个时，b 的取值范围比当前大。
// 例如 [-1, -1, 0, 1]：第一个 -1 已经可以和 0、1 组成 [-1, 0, 1]。第二个 -1 能尝试
// 的 b、c 只会更少，不可能产生新的值组合。

// 解法：
// 1. 暴力，三层 for 循环
// 2. 先用两层 for 循环确定 a 和 b，再借助哈希表快速查找第三个数是否存在。注意：为了
//    确保一个元素不能使用两次，以及避免重复，需满足 a下标 < b下标 < c下标。

func threeSum(nums []int) [][]int {
	result := make([][]int, 0)
	slices.Sort(nums) // 为了去除第三类重复
	n := len(nums)
	aMap := make(map[int]int)
	// 注意，有重复元素保存的就是最后一个的下标
	// 例如：[-1,0,1,2,-1,-4]
	for i, v := range nums {
		// 保存下标是为了去除第二类重复
		aMap[v] = i
	}

	for i := 0; i < n; i++ {
		a := nums[i]
		// 为了去除第三类重复。确保 (a, b) 跟已有解不同
		if i > 0 && nums[i-1] == a {
			continue
		}
		// bug: j := 1
		for j := i + 1; j < n; j++ {
			b := nums[j]
			// 为了去除第三类重复。确保 (a, b) 跟已有解不同
			// bug：相应地，这里也不能是 j > 1 && ...
			if j > i+1 && nums[j-1] == b {
				continue
			}

			c := -(a + b)
			cIdx, found := aMap[c]
			if !found {
				continue
			}
			if cIdx > j { // 为了去除第二类重复。确保 a下标 < b下标 < c下标
				result = append(result, []int{a, b, c})
			}
		}
	}
	return result
}
