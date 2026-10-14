package binary_search

// https://leetcode.cn/problems/guess-number-higher-or-lower/description/

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
	low := 1
	high := n
	for low <= high {
		mid := (low + high) / 2
		guessV := mid
		gueesR := guess(guessV)
		if gueesR == 0 {
			return guessV
		} else if gueesR == 1 {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func guess(num int) int {
	return 6
}
