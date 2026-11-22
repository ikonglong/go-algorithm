package twopointer

// https://leetcode.cn/problems/diao-zheng-shu-zu-shun-xu-shi-qi-shu-wei-yu-ou-shu-qian-mian-lcof/

func trainingPlan_(actions []int) []int {
	lo, hi := 0, len(actions)-1
	for lo < hi {
		// lo < hi 防止 lo 越界
		// 让 lo 停在偶数上
		for lo < hi && actions[lo]&1 == 1 {
			lo++
		}
		// hi > lo 防止 hi 越界
		// 让 hi 停在奇数上
		for hi > lo && actions[hi]&1 == 0{
			hi--
		}

		actions[lo], actions[hi] = actions[hi], actions[lo]
		// 这两句非必需。因为后续迭代中循环体开头的逻辑会移动它们
		lo++
		hi--
	}
	return actions
}

// 将前一个版本中内嵌的 for 循环拉平
func trainingPlan2_(actions []int) []int {
	lo, hi := 0, len(actions)-1
	for lo < hi {
		// 找到偶数，停住。continue 的目的是跳过剩余逻辑，只做让 pa 停在偶数上这件事
		if actions[lo]&1 == 1 {
			lo++
			// 注意，continue 虽然写在 if 块中，但它影响的并不是
			// if 块的执行，而是最近的 for 循环体逻辑的执行
			continue
		}
		// 找到奇数，停住。continue 的目的是跳过剩余逻辑，只做让 pb 停在奇数上这件事
		if actions[hi]&1 == 0 {
			hi--
			continue
		}

		// 在 pa、pb 分别停在偶数、奇数上之前，不能执行下面的逻辑
		actions[lo], actions[hi] = actions[hi], actions[lo]
		// 这两句非必需。因为后续迭代中循环体开头的逻辑会移动它们
		lo++
		hi--
	}
	return actions
}
