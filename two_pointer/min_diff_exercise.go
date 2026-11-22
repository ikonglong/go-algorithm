package twopointer

import (
	"math"
	"slices"
)

func smallestDifference_(a []int, b []int) int {
	slices.Sort(a)
	slices.Sort(b)
	pa, pb := 0, 0
	min := diffAbs(a[pa], b[pb])

	for pa < len(a) && pb < len(b) {
		diff := diffAbs(a[pa], b[pb])
		if diff < min {
			min = diff
		}

		// 为下个迭代做准备
		if a[pa] > b[pb] {
			pb++
		} else if a[pa] < b[pb] {
			pa++
		} else {
			return 0
		}
	}

    return min
}

func diffAbs(n1, n2 int) int {
	return int(math.Abs(float64(n1-n2)))
}
