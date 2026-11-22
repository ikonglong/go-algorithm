package twopointer

// https://leetcode.cn/problems/diao-zheng-shu-zu-shun-xu-shi-qi-shu-wei-yu-ou-shu-qian-mian-lcof/

// 算法思路：
// 1. 创建两个指针 low、high，分别指向第一个和最后一个元素
// 2. 使用 low 和 high 遍历数组，直到 low 大于等于 high
//		a. low 从左向右找偶数，找到就停住
//		b. high 从右向左找奇数，找到就停住
//      c. 当 low、high 分别停在偶数、奇数上时，交换它们
//    	d. low 右移、high 左移

func trainingPlan(actions []int) []int {
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
func trainingPlan2(actions []int) []int {
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

// 算法思路：
// 1. [0, p) 是奇数区间，[p, n-1] 既是待处理区间，又表示偶数区间。初始时 p = 0，即奇数区间为空
//    	注意理解如下几点：
//    	a. p 是【下一个奇数要存放的位置】，是奇数区间的右开边界，也是待处理区间的左闭边界
//    	b. [p, n-1] 既是待处理区间，也表示偶数区间，因为当将其中所有的奇数都交换至奇数区
//         间后，它就是偶数区间了
// 2. 用 i 遍历待处理区间
// 		a. i 向右扫描，找到奇数就停住
//		b. 此时，必然是 p 停在偶数上，i 停在奇数上，交换它们
//      c. 右移 p，让它指向【下一个奇数要存放的位置】
//      d. 交换后必须右移 i，确保 i >= p。若发生 i < p，逻辑就不对了
func trainingPlan3(actions []int) []int {
	n := len(actions)
	p := 0 // 下一个奇数要存放的位置
	for i:=p; i < n && p <= i; {
		// 找到奇数，停住
		if actions[i]&1 != 1 {
			i++
			continue
		}

		actions[p], actions[i]=actions[i], actions[p]
		p++
		// 交换后必须右移 i，确保 i >= p。若发生 i < p，逻辑就不对了
		i++
	}
	return actions
}

// 上一个版本的简洁写法，用来对比理解
func trainingPlan4(actions []int) []int {
	p := 0
	for i := 0; i < len(actions); i++ {
		// 找到奇数，停住
		if actions[i]&1 != 1 {
			// i++
			continue
		}

		actions[p], actions[i] = actions[i], actions[p]
		p++
	}
	return actions
}
