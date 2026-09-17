package backtrack

// https://leetcode.cn/problems/sudoku-solver/
// leetcode 报告有个测试用例计算超时
func solveSudoku_not_good(board [][]byte) {
	// 数独表格有多少行，就对应多少个决策阶段
	// 选择列表：每个阶段处理的行是确定的，不可选择。只有空列可选择
	// 路径：不断填充的数独表格就是路径
	var backtrack func(k int, board [][]byte)
	backtrack = func(k int, board [][]byte) {
		// 过程每一步中会检查某个位置是否可放，不可放就被剪枝了，
		// 所以能走到最后一步就是可行解
		if k == len(board) {
			return
		}

		// k 就是当前阶段准备填充的行。len(board[k]) 恒等于 9。
		// 对于选择列表
		// 【错误的理解】：该行的空列与1-9个数字的配对组合，就是所有可能的选择
		// 【正确的理解】：该行的空列
		for i := 0; i < len(board[k]); i++ {
			if board[k][i] != '.' {
				continue
			}

			for v := 1; v <= 9; v++ {
				if isVOk(board, k, i, v) {
					board[k][i] = byte('0' + v)
					break
				}
				// 如果不 OK，继续尝试下一个值
			}

			// 如果仍为空，表明无法找到有效填充值，剪枝
			if board[k][i] == '.' {
				continue
			}

			// 注意，前面的 for 循环已经为当前单元格填出了有效值
			backtrack(k+1, board)
			// 恢复
			board[k][i] = '.'
		}
	}

	backtrack(0, board)
}

func isVOk(board [][]byte, row, col, v int) bool {
	// 检查当前这一行是否有重复的数字
	for i := 0; i < len(board[row]); i++ {
		if board[row][i] != '.' && board[row][i] == byte('0'+v) {
			return false
		}
	}

	// 检查当前这一列是否有重复的数字
	for i := range len(board) {
		if board[i][col] != '.' && board[i][col] == byte('0'+v) {
			return false
		}
	}

	for i := 0; i < len(board); i = i + 3 {
		for j := 0; j < len(board[row]); j = j + 3 {
			if row >= i && row < i+3 && col >= j && col < j+3 {
				for m, n := i, j; m < i+3 && n < j+3; m, n = m+1, n+1 {
					if board[m][n] != '.' && board[m][n] == byte('0'+v) {
						return false
					}
				}
			}
		}
	}

	return true
}
