package binary_search

// https://leetcode.cn/problems/find-smallest-letter-greater-than-target/description/
// 返回 letters 中大于 target 的最小的字符。如果不存在这样的字符，则返回 letters 的第一个字符。

func nextGreatestLetter(letters []byte, target byte) byte {
	n := len(letters)
	if target < letters[0] || target >= letters[n-1] {
		return letters[0]
	}
	//过了上面的检查，必然成立：letter[0]<=target<letters[n-1]
	low, high := 0, n-1
	tIdx := -1
	for low <= high {
		mid := (low + high) / 2
		if target == letters[mid] {
			tIdx = mid
			break
		} else if letters[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	if tIdx != -1 {
		for i := tIdx + 1; i < n; i++ {
			if letters[i] > target {
				return letters[i]
			}
		}
	} else {
		return letters[low]
	}
	return letters[0]
}

func nextGreatestLetter2(letters []byte, target byte) byte {
	n := len(letters)
	if target < letters[0] || target >= letters[n-1] {
		return letters[0]
	}
	//过了上面的检查，必然成立：letter[0]<=target<letters[n-1]
	low, high := 0, n-1
	for low < high {
		mid := (low + high) / 2
		if target <= letters[mid] {
			high = mid
		} else {
			low = mid + 1
		}
	}
	// 循环结束时 low == high，且 letters[low/high] 要么等于 target，
	// 要么是最接近 target 的前、后字母
	for i := low; i < n; i++ {
		if letters[i] > target {
			return letters[i]
		}
	}
	return letters[0]
}
