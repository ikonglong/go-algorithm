package hashmap

// https://leetcode.cn/problems/two-sum/
// 只需返回一个解，且可以任意顺序返回元素下标
// 注意：
// 1. 一个元素只能使用一次
// 2. 如果 a、b 的和为 target，那么 (a, b)、(b, a) 是同一个解，不能出现两次

// 解法：
// 1. 暴力
// 2. 哈希表
// 3. 排序+双指针

func twoSum(nums []int, target int) []int {
	// map, key: element, value: 元素下标
	// 保存下标是为了去除重复
	numMap := make(map[int]int)
	for i, v := range nums {
		numMap[v] = i
	}

	for i := range nums {
		b := target - nums[i]
		if idx, found := numMap[b]; found {
			// 如果要求返回所有解，要去除重复，则需这么写
			// if i > idx { ... }
			if idx != i { // 检查另一个元素的下标是否是当前下标
				return []int{i, idx}
			}
		}
	}

	return nil
}
