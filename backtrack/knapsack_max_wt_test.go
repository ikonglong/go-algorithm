package backtrack

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func knapsackMaxWt(items []int, wCap int) int {
	// 全局结果
	maxW := math.MinInt

	// 决策阶段：每处理一个物品
	// 选择列表：每个阶段处理一个物品，处理时只有两种选择：装或不装
	// 路径：因为不需要计算最终装入哪些物品使得重量最大，因此不需要物品列表，只需要一个累计重量
	// 还需要一个额外参数（背包的剩余容量），以便在递归过程中进行剪枝
	var backtrack func(items []int, k int, cw int, wCap int)
	backtrack = func(items []int, k, cw, wCap int) {
		// cw == wCap 用来对递归树或决策树进行剪枝。
		// 如果当前重量已经达到背包容量，意味着已经得到最优解，无需进行后续计算了
		if cw == wCap || k == len(items) {
			if cw > maxW {
				maxW = cw
			}
			return
		}

		// 这里的选择列表很简单,决定当前物品装还是不装，因此，不需要 for 循环

		// 每个阶段的选择列表基本独立，互不影响，因此选择列表不需要更改
		// 路径 cw 只是一个局部数值变量，这里选择不更改它，
		// 因此递归调用后也无需恢复

		// 不装。选择不装就无需关心装后是否超重
		backtrack(items, k+1, cw, wCap)

		if cw+items[k] <= wCap {
			// 装
			backtrack(items, k+1, cw+items[k], wCap)
			// 由于都是局部变量且没有被更改，因此无需恢复
		} else {
			// 对 cw+items[k] > wCap 情形进行剪枝
			// 不要将剪枝和选择“不装”混为一谈了。
			// 剪枝不是“不装”，“不装”是一种决策，即选择不将当前物品装入背包，但要继续推进下一阶段。
			// 而剪枝是：选择装，但若装后超重，则直接返回，不继续推进下一个阶段
		}
	}

	backtrack(items, 0, 0, wCap)
	return maxW
}

// knapsackMaxWt([3, 2, 6, 8], 7), expected: 6, actual: math.MinInt
//
// 问题分析：
// 1. 注释中说每件物品有“装”和“不装”两种选择，但代码只实现了“装”
// 2. 装 3 → 装 2 → 尝试装 6 时超重 → return
//    k=0, cw=0
//    k=1, cw=3
//    k=2, cw=5
//    5 + 6 > 7，因此直接返回。
//    由于没有任何递归到达 cw == wCap || k == len(items)，maxW 从未更新，最终仍是 math.MinInt

func knapsackMaxWt_Bad(items []int, wCap int) int {
	// 全局结果
	maxW := math.MinInt

	// 决策阶段：每处理一个物品
	// 选择列表：每个阶段处理一个物品，处理时只有两种选择：装或不装
	// 路径：因为不需要计算最终装入哪些物品使得重量最大，因此不需要物品列表，只需要一个累计重量
	// 还需要一个额外参数（背包的剩余容量），以便在递归过程中进行剪枝
	var backtrack func(items []int, k int, cw int, wCap int)
	backtrack = func(items []int, k, cw, wCap int) {
		// cw == wCap 用来对递归树或决策树进行剪枝。
		// 如果当前重量已经达到背包容量，意味着已经得到最优解，无需进行后续计算了
		if cw == wCap || k == len(items) {
			if cw > maxW {
				maxW = cw
			}
			return
		}

		// 这里的选择列表很简单,决定当前物品装还是不装，因此，不需要 for 循环

		// 导致 bug 的错误思考路径：
		// 为啥会一路不装选择呢？因为将剪枝和选择“不装”选择混为一谈了。
		// 剪枝不是“不装”，“不装”是一种决策，即选择不将当前物品装入背包，但要继续推进下一阶段。
		// 而剪枝是：选择装，但若装后超重，则直接返回，不继续推进下一个阶段
		//
		// 下面的条件满足时同时包含：
		// cw+items[k] > wCap 且选择不装。选择不装时无需关心装后是否超重
		// cw+items[k] > wCap 且选择装&剪枝，因为装后超重
		if cw+items[k] > wCap { // 剪枝
			return
		}

		// 每个阶段的选择列表基本独立，互不影响，因此选择列表也不需要更改
		// 路径 cw 只是一个局部数值变量，这里选择不更改它，
		// 因此递归调用后也无需恢复
		//
		// 只有 cw+items[k] <= wCap 才执行
		backtrack(items, k+1, cw+items[k], wCap)
	}

	backtrack(items, 0, 0, wCap)
	return maxW
}

func TestKnapsackMaxWt(t *testing.T) {
	r1 := knapsackMaxWt([]int{3, 2, 6, 8}, 7)
	assert.Equal(t, 6, r1)

}
