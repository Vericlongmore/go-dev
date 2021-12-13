package leetcode

var dir = [][]int{
	{-1, 0},
	{0, 1},
	{1, 0},
	{0, -1},
}

// 解法二 DFS
func solve1(board [][]byte) {
	for i := range board {
		for j := range board[i] {
			if i == 0 || i == len(board)-1 || j == 0 || j == len(board[i])-1 {
				if board[i][j] == 'O' {
					dfs130(i, j, board)
				}
			}
		}
	}

	for i := range board {
		for j := range board[i] {
			if board[i][j] == '*' {
				board[i][j] = 'O'
			} else if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}
}

func dfs130(i, j int, board [][]byte) {
	if i < 0 || i > len(board)-1 || j < 0 || j > len(board[i])-1 {
		return
	}
	if board[i][j] == 'O' {
		board[i][j] = '*'
		for k := 0; k < 4; k++ {
			dfs130(i+dir[k][0], j+dir[k][1], board)
		}
	}
}
