package slidingwin

// 算法思路心智模型：
// 创建两个指针 left, right，分别表示结果字符串的左侧闭边界、右侧开边界。
// right 既指向结果字符串的右侧开边界，又指向未处理部分的头元素。初始时，
// 最长重复字符串滑动窗口为 [0, 0)。
//
// 重复下述步骤，遍历整个字符串：
//  1. 更新 s[right] 字符计数，更新滑窗中重复最多的字符计数、移动 right
//  2. 当 right 移动到滑窗不再满足要求（替换额度超支，替换了k+1次）时，
//     停止移动 right，移动 left，使滑窗再次满足要求。
//
// 记忆链：滑窗从满足要求，推进到即将不满足要求时，一定要先破再立，即先转变
// 为不合法的，再重新转变为合法的。不破不立
func characterReplacement2(s string, k int) int {
	if len(s) < 2 {
		return len(s)
	}

	maxLen := 0
	// 另一组变量命名：charFreq, maxFreq
	charCounts := make([]int, 26) // 假设只包含大写英文字母
	maxCharCount := 0             // 维护字符最大出现频率

	// [left, right) 表示子串的滑动窗口
	left, right := 0, 0
	for right < len(s) {
		charCounts[s[right]-'A']++
		maxCharCount = max(maxCharCount, charCounts[s[right]-'A'])
		right++

		// 不选择维护剩余替换额度来判断是否超支，而是通过维护字符最大出现频率来计算是否超支
		if right-left > maxCharCount+k {
			charCounts[s[left]-'A']--
			maxCharCount = max(maxCharCount, charCounts[s[left]-'A'])
			left++
		}

		maxLen = max(maxLen, right-left)
	}
	return maxLen
}

func characterReplacement3(s string, k int) int {
	maxLen := 0
	n := len(s)

	// 理论上看，子串可能被统一为 26 个大写字母中的任何一个，所以每一个都尝试一次
	for targetChar := byte('A'); targetChar <= 'Z'; targetChar++ {
		left, right := 0, 0
		count := k

		for right < n {
			if s[right] != targetChar {
				count--
			}
			right++

			// 想在窗口“刚好结束”时更新 maxLen？对于此题目，做不到，因为count=0时后面的字符可能
			// 都是目标字符，所以认为count=0就刚好结束窗口是错误的。只能在替换超支一次后或说窗口
			// 刚刚非法时更新。注意，还需要在循环结束后补一次更新，因为可能永远不会超支。
			// 例如 "AAAA", k=0、目标为 A，全程不会超支。如果只在超支时更新，就会漏掉答案 4，
			// 因此末尾必须补算。
			// if count < 0 {
			// 	maxLen = max(maxLen, right-1-left)
			// }

			// 注意，count=0时替换额度刚好用完，迭代开始时的s[right]刚被纳入窗口，此时窗口
			// 是满足题目要求的，即合法的。count<0时替换额度超支，迭代开始时的s[right]也是
			// 刚被纳入窗口，是它导致额度超支的，此时窗口是非法的。窗口合法时不能缩窗，因为题
			// 目求最大长度，此时缩窗结果必然小于当前 maxLen，做了也是无用功。只能先通过超支
			// 打破合法状态，再恢复合法，过程中才能找到更大值。此所谓不破不立，先破再立。
			//
			// 如果额度超支，持续缩窗，直到恢复合法。
			for count < 0 {
				if s[left] != targetChar {
					// 之前进入窗口时消耗了额度，移出时需要恢复
					count++
				}
				left++
			}

			// 只要窗口合法（即count>=0），就更新 maxLen。这样就不需要
			// 在循环结束后再补一次更新。这种写法更统一，容易检查正确性，不容易遗漏埋坑
			maxLen = max(maxLen, right-left)
		}

		// 如果上面的循环中仅在刚刚超支时更新 maxLen，循环后就需要补一次更新
		// maxLen = max(maxLen, right-left)
	}
	return maxLen
}

// 窗口首字符不一定就是窗口中出现最多的字符。换句话说，首字符不一定
// 就不是要被替换的字符
func characterReplacement4_bug(s string, k int) int {
	if len(s) == 0 {
		return 0
	}

	n := len(s)
	maxLen := 0
	left, right := 0, 1
	prev := s[left]
	count := k
	for right < n {
		if count > 0 && s[right] != prev {
			count--
		} else if count == 0 && s[right] != prev {
			// if count == 0 && s[right] != prev {
			maxLen = max(maxLen, right-left)
			if s[left] != prev {
				count += 1
			}
			left++
			prev = s[left]
			// 仅仅在右移 left 前检查是否不同，不同就把额度补回一个，这不够。
			// 更新比较目标后，留在窗口里的字符是否需要替换也可能改变，
			// 替换额度需要重新计算。
		}
		right++

		// if count > 0 {
		// 	if s[right] != prev {
		// 		count--
		// 	}
		// 	right++
		// } else {
		// 	if s[right] != prev {
		// 		maxLen = max(maxLen, right-left)
		// 		right++
		// 		left++
		// 		prev = s[left]
		// 		count = k
		// 	} else {
		// 		right++
		// 	}
		// }
	}
	maxLen = max(maxLen, right-left)

	return maxLen
}

func characterReplacement5_bug(s string, k int) int {
	if len(s) == 0 {
		return 0
	}

	n := len(s)
	maxLen := 0
	left, right := 0, 0
	prev := s[left]
	count := k
	for right < n {
		if count > 0 && s[right] != prev {
			count--
		}
		if count == 0 && s[right] != prev {
			maxLen = max(maxLen, right-left+1)
			if s[left] != prev {
				count += 1
			}
			left++
			prev = s[left]
			// count = k
		}
		right++
	}

	return maxLen
}
