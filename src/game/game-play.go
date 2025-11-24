package game

import "fmt"

type Grille struct {
	Board [][]string
}

func addPion(grille *Grille, col int, player Player) bool {

	nLigne := -1

	for i := len(grille.Board) - 1; i > 0; i-- {
		if grille.Board[i][col-1] == " " {
			nLigne = i
			break
		}
	}

	if nLigne == -1 {
		return false
	}

	grille.Board[nLigne][col-1] = player.Color
	return true
}

func checkWin(grille *Grille, players [2]Player) bool {

	rows := len(grille.Board)
	cols := len(grille.Board[0])

	for i := rows - 1; i >= 0; i-- {
		for j := 0; j < cols; j++ {
			if grille.Board[i][j] != " " && j+3 < cols && grille.Board[i][j] == grille.Board[i][j+1] && grille.Board[i][j+1] == grille.Board[i][j+2] && grille.Board[i][j+2] == grille.Board[i][j+3] {
				caseColor := grille.Board[i][j]

				switch caseColor {
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

func GamePlay(grille *Grille, col int, players [2]Player, currentTurn *int) {

	for !checkWin(grille, players) {

		if *currentTurn%2 != 0 {
			for !addPion(grille, col, players[0]) {
				fmt.Println(*currentTurn)
				fmt.Println("Colonne pleine, choisissez-en une autre.")
			}
			(*currentTurn)++
		} else {
			for !addPion(grille, col, players[1]) {
				fmt.Println(*currentTurn)
				fmt.Println("Colonne pleine, choisissez-en une autre.")
			}
			(*currentTurn)++
		}
	}

	checkWin(grille, players)
}
