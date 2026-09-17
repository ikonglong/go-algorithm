package backtrack

// https://leetcode.cn/problems/wildcard-matching/description/
func isMatch(s string, p string) bool {
	isMatched := false

	// 模式字符串包含多少个字符，就有多少个决策阶段
	// 决策阶段：一个阶段处理一个模式字符，需要决定这个字符跟字符串 s 中哪些连续字符匹配
	// 选择列表：一个模式字符跟字符串 s 的所有匹配选择
	// 路径：模式字符串 p、字符串 s 分别匹配到哪里了？
	//      分别用下一个要匹配的字符索引 pi、si 来表示
	var backtrack func(s string, p string, si, pi int)
	backtrack = func(s string, p string, si, pi int) {
		// 结束条件
		if pi == len(p) {
			if pi == len(p) && si == len(s) {
				isMatched = true
			}
			return
		}

		// pi 的取值始终满足 0 <= pi <= len(p)。
		// 每次递归最多让 pi 增加 1；当 pi == len(p) 时立即返回，
		// 因此下面一定有 pi < len(p)，访问 p[pi] 是安全的。

		if p[pi] == '*' {
			// 对 s="aa"、p="a*"，第一次走 p[pi] == s[si] 分支。第二次
			// 走这里，此时 si=1,pi=1。如果 for k=0;k<1;k++{...}，就只递归
			// 一次：backtrack(si=1, pi=2)，在这次递归调用中，满足结束条件，
			// 但不满足 si == length，结果为 false。因此正确的写法：
			// k<=len(s)-si，这样才能让 si 也能达到 len(s)
			//
			// k=0时, si 和 pi+1 配对就表达了p[pi]没有匹配任何字符的情形
			for k := 0; k <= len(s)-si; k++ {
				backtrack(s, p, si+k, pi+1)
			}
		} else if p[pi] == '?' {
			// 匹配零个字符
			// 若题目要求，问号必须匹配一个字符
			// backtrack(s, p, si, pi+1)

			// 匹配 1 个字符
			// 因为这里的意图是要匹配一个字符，因此需要 si+1<len(s)

			//if si+1 < len(s) {
			// 上面的条件有 bug。例如：s="a",p="?*"
			// 第一次调用 backtrack(s, p, 0, 0)，执行到这里时：
			// si + 1 < len(s) => 1 < 1 // false => 没有递归，
			// 执行结束，结果为 false

			// 该条件表达的是问题领域逻辑：'?'必须消耗一个真实存在的字符。
			// 它的副作用是具有剪枝效果，阻止程序进入si > len(s)的无效状态。
			// 由于这个分支没有直接写 `s[si]`，因此该判断不是在保护某个当前的下标访问
			if si < len(s) {
				// 注意理解，'?'跟字符的匹配就是si和pi的匹配。发起下次递归时，
				// 匹配就已经完成了，因为si+1和pi+1是下一对准备处理的字符。这个实现
				// 很隐晦，容易误解为'?'跟字符的匹配是si+1和pi+1的匹配。
				backtrack(s, p, si+1, pi+1)
			}
		} else if si < len(s) && p[pi] == s[si] { // 普通字符，直接检测是否相同
			// 如果不加 si < len(s)，那么就可能下标越界
			// 例如：s = "adceb", p = "*a*b"
			// 第一个*会枚举所有可能的匹配。若匹配所有字符，下一次递归是backtrack(si=5, pi=1)。
			// 这次递归中，si == len(s),p[pi] == 'a'，继续执行 p[pi] == s[si] 就越界了
			backtrack(s, p, si+1, pi+1)
		}
		// 剪枝，这里不写任何代码，甚至下面的代码块跟不写任何代码效果完全相同
		// else {
		// 	isMatched = false
		// }
	}

	backtrack(s, p, 0, 0)
	return isMatched
}

// 这个实现版本可能更易懂
func isMatch2(s string, p string) bool {
	isMatched := false

	// 模式字符串包含多少个字符，就有多少个决策阶段
	// 决策阶段：一个阶段处理一个模式字符，需要决定这个字符跟字符串 s 中哪些连续字符匹配
	// 选择列表：一个模式字符跟字符串 s 的所有匹配选择
	// 路径：模式字符串 p、字符串 s 分别匹配到哪里了？
	//      分别用下一个要匹配的字符索引 pi、si 来表示
	var backtrack func(s string, p string, si, pi int)
	backtrack = func(s string, p string, si, pi int) {
		// 结束条件
		if pi == len(p) {
			if pi == len(p) && si == len(s) {
				isMatched = true
			}
			return
		}

		// pi 的取值始终满足 0 <= pi <= len(p)。
		// 每次递归最多让 pi 增加 1；当 pi == len(p) 时立即返回，
		// 因此下面一定有 pi < len(p)，访问 p[pi] 是安全的。

		if p[pi] == '*' {
			for k := 0; k <= len(s)-si; k++ {
				backtrack(s, p, si+k, pi+1)
			}
			return
		}

		if si >= len(s) { // 确保 p 中的问号和普通字符有要消耗的一个字符
			return
		}

		if p[pi] == '?' || p[pi] == s[si] {
			backtrack(s, p, si+1, pi+1)
		}
		// 剪枝，这里不写任何代码，甚至下面的代码块跟不写任何代码效果完全相同
		// else {
		// 	isMatched = false
		// }
	}

	backtrack(s, p, 0, 0)
	return isMatched
}

func isMatch_bad(s string, p string) bool {
	isMatched := false // 最终结果

	// 模式字符串包含多少个字符，就有多少个决策阶段
	// 决策阶段：一个阶段处理一个模式字符，需要决定这个字符跟字符串 s 中哪些连续字符匹配
	// 选择列表：一个模式字符跟字符串 s 的所有匹配选择
	// 路径：模式字符串 p、字符串 s 分别匹配到哪里了？
	//      分别用下一个要匹配的字符索引 pi、si 来表示
	var backtrack func(s string, p string, si, pi int)
	backtrack = func(s string, p string, si, pi int) {
		// 结束条件

		// --- 尝试 1 ---
		// if pi == len(p) || si == len(s) {
		// 	isMatched = true
		// 	return
		// }
		// ------
		// si、pi 任何一个匹配完了就应该结束，这点是确定的。
		// 但任何一个匹配完了还无法确定两者是否匹配

		// --- 尝试 2 ---
		// if pi == len(p) && si == len(s) {
		// 	isMatched = true
		// 	return
		// }
		// ------
		// 埋有 bug，它只有在两个字符串同时结束时才返回。但模式已经结束、
		// 原字符串尚未结束时，也必须结束当前递归，只是结果应为“不匹配”。
		// ------
		// s = "aa", p = "a"
		// backtrack(si=0, pi=0)
		// 满足 p[pi] == s[si] 分支
		//   ↓
		// backtrack(si=1, pi=1)
		//   pi == len(p) 但 si < len(s)
		//   p[pi] == '*' 就数组越界了

		// --- 尝试 3 ---
		// if pi == len(p) || si == len(s) {
		// 	if pi == len(p) && si == len(s) {
		// 		isMatched = true
		// 	}
		// 	return
		// }
		// ------
		// 埋有 bug：s = "", p = "******", expected: true, actual: false

		if p[pi] == '*' {
			// k=0时, si 和 pi+1 配对就表达了p[pi]没有匹配任何字符的情形
			// 对 s="aa"、p="a*"，第一次走 p[pi] == s[si] 分支。第二次
			// 走这里，此时 si=1,pi=1。如果 for k=0;k<1;k++{...}，就只递归
			// 一次：backtrack(si=1, pi=2)，在这次递归调用中，满足结束条件，
			// 但不满足 si == length，结果为 false。因此正确的写法：
			// k<=len(s)-si，这样才能让 si 也能达到 len(s)
			for k := 0; k <= len(s)-si; k++ {
				backtrack(s, p, si+k, pi+1)
				// 递归调用前没有修改当前调用节点的任何状态，
				// 哪怕是局部变量，无需恢复
			}
		} else if p[pi] == '?' {
			// 不匹配任何字符
			// 题目要求，问号必须匹配一个字符
			// backtrack(s, p, si, pi+1)
			// 匹配一个字符
			backtrack(s, p, si+1, pi+1)
			// 同样，这两个定位调用前没有修改当前调用节点的任何状态，
			// 因此无需恢复
		} else if p[pi] == s[si] { // 普通字符，直接比较是否相同
			backtrack(s, p, si+1, pi+1)
		}
		// 剪枝，这里不写任何代码，甚至下面的代码块跟不写任何代码效果完全相同
		// else {
		// 	isMatched = false
		// }
	}

	backtrack(s, p, 0, 0)
	return isMatched
}
