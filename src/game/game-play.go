package game

import "fmt"

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

func addPion(grille *Grille, col int, player Player) {

	nLigne := -1

	for i := len(grille.Board) - 1; i > 0; i-- {
		if grille.Board[i][col-1] == " " {
			nLigne = i
			break
		}
	}

	if nLigne == -1 {
		fmt.Println("Erreur : Colonne remplie")
		return
	}

	grille.Board[nLigne][col-1] = player.Color
}

func checkWin(grille *Grille, players [2]Player) bool {

	rows := len(grille.Board)
	cols := len(grille.Board[0])

	for i := rows - 1; i >= 0; i-- {
		for j := 0; j < cols; j++ {
			if grille.Board[i][j] != " " && j+3 < cols && grille.Board[i][j] == grille.Board[i][j+1] && grille.Board[i][j+1] == grille.Board[i][j+2] && grille.Board[i][j+2] == grille.Board[i][j+3] {
				caseSign := grille.Board[i][j]

				switch caseSign {
				case players[0].Color:
					fmt.Printf("Bravo ! %s a gagné la partie !\n", players[0].Name)
					return true
				case players[1].Color:
					fmt.Printf("Bravo ! %s a gagné la partie !\n", players[1].Name)
					return true
				}
			}
		}
	}

	return false
}

func GamePlay(grille *Grille, col int, player1 Player, player2 Player) {

	for !checkWin(grille, [2]Player{player1, player2}) {

	}
}
