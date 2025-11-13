package game

type Grille struct {
	Board [][]string
}

func InitGrille() Grille {
	rows, cols := 6, 7
	board := make([][]string, rows)
	for i := range board {
		board[i] = make([]string, cols)
		for j := range board[i] {
			board[i][j] = " "
		}
	}
	return Grille{Board: board}
}
