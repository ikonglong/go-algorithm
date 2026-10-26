package hashmap

// https://leetcode.cn/problems/sum-swap-lcci/
// 给了两个数组，交换其中一个元素后，使得两个数组的元素和相等

func findSwapValues(array1 []int, array2 []int) []int {
	sum1 := 0
	for _, item := range array1 {
		sum1 += item
	}
	sum2 := 0
	a2Map := make(map[int]struct{})
	for _, item := range array2 {
		sum2 += item
		a2Map[item] = struct{}{}
	}
	total := sum1 + sum2
	// 经 LeetCode 测试发现，要求若总和为奇数，直接返回空数组
	// 应该是因为奇数不能平分
	if total&1 == 1 {
		return nil
	}
	diff := total/2 - sum1
	for _, v := range array1 {
		target := v + diff
		if _, found := a2Map[target]; found {
			return []int{v, target}
		}
	}
	return nil
}
