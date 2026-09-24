package slidingwin

import "math"

// 算法思路：
// 1. 创建两个指针 left、right，分别表示滑动窗口的左右边界。初始时滑动窗口还不存在或说为空，因此
//    可以这么表示：[-1,-1],[0,-1],[-1,0],[0,0)。我觉得 [0,0) 理解起来最自然，因此选择它。
// 3. 重复如下步骤，直到 right >= n
//      方法一：
//      1. 右移，Right。更新滑窗计数
//      2. 检查是否覆盖。如果有覆盖，重复下述步骤，直到滑动窗口长度小于 t 长度或者不再覆盖了
// 			a. 计算最小窗口
//          b. 更新滑窗计数，右移 left
//      方法二：
//      1. 先把 s[right] 包含进滑动窗口
//      2. 再循环检查是否有覆盖，直到覆盖了或者 right>=n。重复下述步骤：
//          a. 更新滑窗计数，右移 right
//      3. 检查是否有覆盖。若覆盖，则计算最小窗口，更新滑窗计数，右移 left

// 滑动窗口表示为 [0,0)。针对判断是否覆盖做了改进：不再是比较两个计数数组，而是处理过程中更新
// 记录滑动窗口包含的 t 中字符的个数。注意，t 中的重复字符也要计数。
//
// 跟后一个实现相比，这里不再采用两个字频数组或叫字符计数数组不断累加计数，而是只采用一个字频数组，
// 根据目标字符串 t 初始化它，然后在处理过程中递减直到 0
func minWindow(s string, t string) string {
	if len(t) == 0 || len(s) < len(t) {
		return ""
	}

	tLen := len(t)
	// 滑动窗口还差多少个 t 中字符
	distance := tLen
	minLen := math.MaxInt
	minLeft := -1
	cCounts := make([]int, 128) // 128 ASCII表字符个数
	for _, c := range t {
		cCounts[c]++
	}

	for left, right := 0, 0; right < len(s); {
		// 迭代开始时，滑动窗口还不包含 s[right]
		// 若滑窗中对应字符的个数多于 t 中的，多出的不减
		if cCounts[s[right]] > 0 {
			distance--
		}
		cCounts[s[right]]-- // 结果可能是负数
		right++

		for right-left >= tLen && distance == 0 { // 覆盖了
			if right-left < minLen {
				minLen = right - left
				minLeft = left
			}

			// 因为负数对应的重复字符是多出的，因此大于等于 0 才增加 distance
			if cCounts[s[left]] >= 0 {
				distance++
			}

			// 字符进入滑动窗口时要减去，因此，出去时就应该加上
			cCounts[s[left]]++
			left++
		}
	}

	if minLen == math.MaxInt {
		return ""
	}
	return s[minLeft : minLeft+minLen]
}

// 滑动窗口表示为 [0,0)。针对判断是否覆盖做了改进：不再是比较两个计数数组，而是处理过程中更新
// 记录滑动窗口包含的 t 中字符的个数。注意，t 中的重复字符也要计数
func minWindow2(s string, t string) string {
	if len(t) == 0 || len(s) < len(t) {
		return ""
	}

	tLen := len(t)
	// 滑动窗口包含的 t 中字符的个数
	distance := 0
	minLen := math.MaxInt
	minLeft := -1
	wCounts := make([]int, 128) // 128 ASCII表字符个数
	tCounts := make([]int, 128)
	for _, c := range t {
		tCounts[c]++
	}

	for left, right := 0, 0; right < len(s); {
		// 迭代开始时，滑动窗口还不包含 s[right]
		// 若滑窗中对应字符的个数多于 t 中的，多出的不加 distance
		if wCounts[s[right]] < tCounts[s[right]] {
			distance++
		}
		wCounts[s[right]]++
		// if wCounts[s[right]] <= tCounts[s[right]] {
		// 	distance++
		// }
		right++

		for right-left >= tLen && distance == tLen {
			if right-left < minLen {
				minLen = right - left
				minLeft = left
			}

			// 滑窗中对应字符的个数可能多于 t，注意多出的不减 distance
			if wCounts[s[left]] <= tCounts[s[left]] {
				distance--
			}

			wCounts[s[left]]--
			left++
		}
	}

	if minLen == math.MaxInt {
		return ""
	}
	return s[minLeft : minLeft+minLen]
}

// 滑动窗口表示为 [0,0)，建模自然直观
func minWindow3(s string, t string) string {
	if len(t) == 0 || len(s) < len(t) {
		return ""
	}

	tLen := len(t)
	minLen := math.MaxInt
	minLeft := -1
	wCounts := make([]int, 128) // ASCII 表共 128 个字符
	tCounts := make([]int, 128)
	for _, c := range t {
		tCounts[c]++
	}

	for left, right := 0, 0; right < len(s); {
		// 迭代开始时，滑动窗口还不包含 s[right]
		c := s[right]
		if tCounts[c] > 0 {
			wCounts[c]++
		}
		right++ // 让滑动窗口包含s[right]

		for right-left >= tLen && cover(wCounts, tCounts) {
			// 埋有 bug：虽然第一行计算了minLen，但第二行没有判断当前窗口是否更小，就直接赋值了
			// minLen = min(right-left, minLen)
			// minLeft = left
			if right-left < minLen {
				minLen = right - left
				minLeft = left
			}

			wCounts[s[left]]--
			left++
		}

		// 上面的写法更简洁，去除了 if 语句
		// 		if right-left >= tLen && cover(wCounts, tCounts) {
		// 			for {
		// 				if right-left < minLen {
		// 					minLen = right - left
		// 					minLeft = left
		// 				}
		//
		// 				wCounts[s[left]]--
		// 				left++
		//
		// 				// 循环退出条件写在这里，是为了模拟 do-while 循环
		// 				if right-left < tLen || !cover(wCounts, tCounts) {
		// 					break
		// 				}
		// 			}
		// 		}
	}

	if minLen == math.MaxInt {
		return ""
	}
	return s[minLeft : minLeft+minLen]
}

// 滑动窗口表示为 [0,0]。虽然代码实现理解起来还算通畅，但建模跟实际情形不符，不太自然
func minWindow4(s string, t string) string {
	if len(t) == 0 || len(s) < len(t) {
		return ""
	}

	minLen := math.MaxInt
	minLeft, minRight := -1, -1
	tLen := len(t)
	// 输入有大写字母，大小写字母编码中间含有 6 个其他字符
	wCounts := make([]int, 26*2+6)
	tCounts := make([]int, 26*2+6)

	for _, c := range t {
		tCounts[c-'A']++
	}

	for left, right := 0, 0; right < len(s); right++ {
		// 对滑动窗口中的字符计数
		wCounts[s[right]-'A']++

		// right >= tLen-1 才有必要判断，滑动窗口是否覆盖 T
		if right < tLen-1 {
			continue
		}

		if cover(wCounts, tCounts) {
			for right-left+1 >= tLen && cover(wCounts, tCounts) {
				if right-left+1 < minLen {
					minLeft, minRight = left, right
					minLen = right - left + 1
				}
				wCounts[s[left]-'A']--
				left++
			}
		}
	}

	if minLen == math.MaxInt {
		return ""
	}
	return s[minLeft : minRight+1]
}

func cover(wCounts, tCounts []int) bool {
	for i, v := range tCounts {
		if v == 0 {
			continue
		}
		if v > wCounts[i] {
			return false
		}
	}
	return true
}
