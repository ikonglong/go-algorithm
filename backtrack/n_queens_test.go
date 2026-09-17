package backtrack

const QUEEN byte = 'Q'
const EMPTY byte = '.'

func solveNQueens(n int) [][]string {
	// string 本质就是 []byte
	result := make([][]string, 0)

	board := make([][]byte, n) // n 行
	for i := range board {
		board[i] = make([]byte, n) // n 列
		for j := range board[i] {
			board[i][j] = EMPTY
		}
	}

	var backtrack func(int, [][]byte)
	// 阶段：棋盘行。有多少个棋盘行，就有多少个阶段
	// 选择列表：确定行后，列就是选择列表，虽然这里还有其他约束限制选择。
	// 这里决定基于棋盘动态计算可用选择，而不是用一个单独的参数表示可用选择列表
	// 路径，表示棋盘的二维数组
	backtrack = func(row int, board [][]byte) {
		// k 从 0 开始。k=n-1 时，表示处理最后一行。k==n-1 不能作为递归的基线条件，
		// 因为还没执行下面的主体逻辑，最后一行就未处理
		if row == n {
			result = append(result, copySolution(board))
			return
		}

		for col := range n {
			// 如果当前位置不能放，表示「当前这个候选选择」不可行，
			// 应跳过该列，尝试当前行的下一列
			//
			// if isOk(board, row, col) {
			// 	board[row][col] = QUEEN
			// }
			// backtrack(row+1, board)
			//
			// 在归的过程中，可回溯到任意一个阶段，因此需要
			// 撤销前面做出的选择，恢复「调用节点」状态，以尝试其他选择
			// if board[row][col] == QUEEN {
			// 	board[row][col] = EMPTY
			// }

			// 纠正错误后的版本
			if !isOk(board, row, col) {
				continue
			}
			board[row][col] = QUEEN
			backtrack(row+1, board)
			// 单元格只可能处于两种状态：queen 或 empty，因此无需检测是否为 queen
			board[row][col] = EMPTY
		}
	}

	backtrack(0, board)
	return result
}

func isOk(board [][]byte, row, col int) bool {
	// 先检查列是否能放
	// 为什么i<row？因为对于指定列，只需检查到row-1行，第row行还没有放
	for i := 0; i < row; i++ {
		if board[i][col] == QUEEN {
			return false
		}
	}

	// 检查反对角线是否能放。从前一行开始检查
	// 题目要求检查经过当前位置的所有可能的反对角线，而不是棋盘中央的反对角线
	// if row == col {
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] == QUEEN {
			return false
		}
	}
	// }

	// 检查正对角线是否能放。从前一行开始检查
	// 题目要求检查经过当前位置的所有可能的正对角线，而不是棋盘中央的正对角线
	n := len(board)
	// if row+col == n-1 {
	for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
		if board[i][j] == QUEEN {
			return false
		}
	}
	// }

	return true
}

func copySolution(board [][]byte) []string {
	// r := make([]string, len(board))
	r := make([]string, 0)
	for i := range board {
		// string(board[i]) 从字节数组创建了 string，
		// 因为 string 是不可变字节序列，所以创建时实际上复制了字节数组
		r = append(r, string(board[i]))
	}
	return r
}
