package sort_

// https://leetcode.cn/problems/shu-zu-zhong-de-ni-xu-dui-lcof/
// 在递归归并排序的过程中，累计逆序对个数

func reversePairs(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}
	return mergeSort(nums, 0, len(nums)-1)
}

// 返回逆序对个数
func mergeSort(nums []int, p, q int) int {
	if p == q {
		return 0
	}
	mid := (p + q) / 2
	reverseCount := 0
	reverseCount += mergeSort(nums, p, mid)
	reverseCount += mergeSort(nums, mid+1, q)
	reverseCount += mergeArrRanges(nums, p, mid, q)
	return reverseCount
}

func mergeArrRanges(nums []int, p, q, r int) int {
	aux := make([]int, r-p+1)
	i, j := p, q+1
	k := 0
	reverseCount := 0
	for i <= q && j <= r {
		if nums[j] < nums[i] {
			reverseCount += (q - i + 1)
			// 埋有 bug
			//reverseCount += (q - p + 1)
			aux[k] = nums[j]
			j++
		} else {
			aux[k] = nums[i]
			i++
		}
		k++
	}
	for ; i <= q; i++ {
		aux[k] = nums[i]
		k++
	}
	for ; j <= r; j++ {
		aux[k] = nums[j]
		k++
	}
	for i, v := range aux {
		nums[p+i] = v
	}
	return reverseCount
}
