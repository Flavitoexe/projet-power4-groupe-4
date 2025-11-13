package game
<<<<<<< HEAD
=======

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
>>>>>>> 144efd2f5aa6fc6c69162ecb9530789ad949a80f
