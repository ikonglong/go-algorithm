package binary_search

// https://leetcode.cn/problems/search-a-2d-matrix/
// 算法思路，将二维有序矩阵转换为一维有序数组，转换方式有两种：
// 1. 把二维矩阵转换成一个新的一维数组，在一维数组上二分查找
// 2. 不物理地把矩阵转换为一维数组，只是二分查找采用概念层面的那个一维数组的下标区间

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	n := len(matrix[0])
	len := m * n
	lo, hi := 0, len-1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		// bug：除以 m，对 m 取模
		row := mid / n
		col := mid % n
		item := matrix[row][col]
		if item == target {
			return true
		} else if item < target {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return false
}
