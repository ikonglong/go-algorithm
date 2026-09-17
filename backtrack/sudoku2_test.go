package backtrack

// https://leetcode.cn/problems/sudoku-solver/
// 跟 backtrack/sudoku_test.go 是同一道 leetcode 题
// 更好的实现版本
func solveSudoku(board [][]byte) {
	// 用一个二维数组记录数独表格中每一行每个数字是否出现了
	// 为什么第二维是 10 而不是 9 呢？
	// 因为要用第二维的下标表示数独表格中的数字，以便高效查看某个值是否存在。
	// 数独中的数字是 1 - 9，其下标最大是 9（0 不用），因此必须是 10
	var rowFlags [9][10]bool
	var colFlags [9][10]bool
	// 数独表格被划分成了九块（三行三列），每个块包含九个单元格
	// 将每个块看作一个整体，即二维数组中的一个元素，这样每个块就有了确定的下标。
	// 例如，第二个块的下标是 (0, 1)。每个块儿又包含 9 个单元格，这就是第三维
	var blockFlags [3][3][10]bool

	solved := false

	for i := range 9 {
		for j := range 9 {
			if board[i][j] != '.' {
				v := board[i][j] - '0'
				rowFlags[i][v] = true
				colFlags[j][v] = true
				// (i/3, j/3) 就是当前单元格所属块的下标
				blockFlags[i/3][j/3][v] = true
			}
		}
	}

	// 有多少个单元格，就有多少个决策阶段。逐行处理，一行内逐个单元格处理。
	// 选择列表：每个单元格有 1-9 9 个候选选择
	// 路径：board，它记录了每一步的状态
	var backtrack func(row, col int)
	backtrack = func(row, col int) {
		if row == 9 {
			// 前面的步骤中如果无有效填充就被剪枝了。
			// 因此，只有有效的填充才会走到这里，这意味着问题已解决
			solved = true
			return
		}

		// 下面的代码无需担心 row 越界，因为上面 row 等于 9 就结束了

		if board[row][col] == '.' {
			for v := 1; v <= 9; v++ {
				// 检查同一行、同一列以及所属块是否有相同值
				if !rowFlags[row][v] && !colFlags[col][v] && !blockFlags[row/3][col/3][v] {
					board[row][col] = byte('0' + v)
					rowFlags[row][v] = true
					colFlags[col][v] = true
					blockFlags[row/3][col/3][v] = true

					nextRow := row
					nextCol := col + 1
					if nextCol > 8 {
						nextRow++
						nextCol = 0
					}
					backtrack(nextRow, nextCol)

					// row == 9 时只是把 solved 设为 true，然后从最深层递归返回。
					// 返回上一层后，如果没有提前 return ，代码会继续执行「恢复状态」，
					// 于是刚找到的答案（保存在 board）又被逐层擦除了。所有原本为空的位置
					// 最终都恢复成 '.'，所以输出恰好等于输入。因此最终状态是：
					// solved = true      // 我曾经找到过答案
					// board  = 原始棋盘   // 具体答案已被逐层擦除
					if solved {
						return
					}

					// 恢复状态
					board[row][col] = '.'
					rowFlags[row][v] = false
					colFlags[col][v] = false
					blockFlags[row/3][col/3][v] = false
				} // else 就继续尝试下一个数
			}
			// 如果 9 个数字都尝试了，仍无有效填充，自然就不会递归，就剪枝了
		} else { // 已填充数字
			nextRow := row
			nextCol := col + 1
			if nextCol > 8 {
				nextRow++
				nextCol = 0
			}
			backtrack(nextRow, nextCol)
			// 有可能递归树的某个分支已得到解。删除也 OK
			// if solved {
			// 	return
			// }
		}
	}

	backtrack(0, 0)
}
