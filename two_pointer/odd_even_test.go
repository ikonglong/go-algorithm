package twopointer

// https://leetcode.cn/problems/diao-zheng-shu-zu-shun-xu-shi-qi-shu-wei-yu-ou-shu-qian-mian-lcof/

// 算法思路：
// 把数组看成三个区间：两端已经归类，中间还未处理。
// pa、pb 从数组两端向中间推进，像两道向内收拢的边界：
//
//     [ 已确认的奇数 |      待处理区间      | 已确认的偶数 ]
//                  pa →               ← pb
//
// pa 左侧 [0, pa) 都是奇数，，pa 是奇数区间的右侧开边界。
// pb 右侧 (pb, n-1] 都是偶数，pb 是偶数区间的左侧开边界。
// [pa, pb] 是待处理区间；初始覆盖整个数组，两端的已归类区间都为空。

func trainingPlan(actions []int) []int {
	// 埋有 bug：pa 未找到偶数时，要右移 pa；此时 pb 还未检查，因此不应该左移
	// for pa, pb := 0, len(actions)-1; pa < pb; pa, pb = pa+1, pb-1 {
	for pa, pb := 0, len(actions)-1; pa < pb; {
		// 犯了很愚蠢的错误：pa&1 == 1
		if actions[pa]&1 == 1 { // 找到奇数，跳过后续逻辑，继续找，直到找到偶数
			pa++
			continue
		}
		if actions[pb]&1 == 0 { // 找到偶数，跳过后续逻辑，继续找，直到找到奇数
			pb--
			continue
		}
		// pa、pb 分别指向偶数、奇数，交换
		actions[pa], actions[pb] = actions[pb], actions[pa]
		pa++
		pb--
	}
	return actions
}

// 虽能工作，但实现挺蠢
//
// 每轮处理：
// 1. pa 向右跳过奇数，找到一个“应该放到右边”的偶数。
// 2. pb 向左跳过偶数，找到一个“应该放到左边”的奇数。
// 3. 若指针未交错，就交换这两个放错位置的数，再让指针各向内移动一步。
// 跳过或交换元素都会扩大已归类区间，直到 pa > pb，中间再无待处理元素。
//
// 例如，用 | 分隔三个区间：
//
//	[1 | 2, 4, 3, 5 | 6]      跳过两端的 1、6，pa 停在 2，pb 停在 5
//	[1, 5 | 4, 3 | 2, 6]      交换 2、5，两个指针向内移动
//	[1, 5, 3 | 空 | 4, 2, 6]  再交换 4、3，待处理区间为空，结束
//
// 两个指针都只向内移动，时间 O(n)，额外空间 O(1)；不保证同类元素的原有顺序。
func trainingPlan2(actions []int) []int {
	// 每个迭代开始时，actions[pa]、actions[pb] 都未处理
	for pa, pb := 0, len(actions)-1; pa <= pb; {
		// 从左向右找偶数
		for pa <= pb && actions[pa]&1 != 0 {
			pa++
		}
		// 从右向左找奇数
		for pa <= pb && actions[pb]&1 == 0 {
			pb--
		}
		// 交换偶数、奇数
		if pa <= pb {
			actions[pa], actions[pb] = actions[pb], actions[pa]
			pa++
			pb--
		}
	}
	return actions
}

// 虽能工作，但实现挺蠢
//
// 算法思路：
//  1. 从 i=0 开始遍历数组。i 既表示奇数区间的右侧开边界，又表示未处理区间的左侧闭边界。
//  2. 每次迭代：
//     从 j=i 开始遍历数组，直到寻找奇数或遍历完成。
//     若未找到，就结束外层循环或直接返回。
//     若找到，就交换 actions[i] 和 actions[j]

func trainingPlan3(actions []int) []int {
	for i := 0; i < len(actions); {
		// 寻找奇数
		j := i
		for ; j < len(actions) && actions[j]&1 == 0; j++ {
		}
		// 没有奇数
		if j == len(actions) {
			break
		}

		actions[i], actions[j] = actions[j], actions[i]
		i++
	}
	return actions
}
