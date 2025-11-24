package game

type Grille struct {
	Board [][]string
}

func addPion(grille *Grille, col int, player Player) bool {

	nLigne := -1

	for i := len(grille.Board) - 1; i >= 0; i-- {
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

func checkWin(grille *Grille, players [2]Player) Player {

	rows := len(grille.Board)
	cols := len(grille.Board[0])
	var winner Player

	// Vérification de victoire
	for i := rows - 1; i >= 0; i-- {
		for j := 0; j < cols; j++ {
			if (grille.Board[i][j] != " " && j+3 < cols && grille.Board[i][j] == grille.Board[i][j+1] && grille.Board[i][j+1] == grille.Board[i][j+2] && grille.Board[i][j+2] == grille.Board[i][j+3]) ||
				(grille.Board[i][j] != " " && i > 2 && grille.Board[i][j] == grille.Board[i-1][j] && grille.Board[i-1][j] == grille.Board[i-2][j] && grille.Board[i-2][j] == grille.Board[i-3][j]) ||
				(grille.Board[i][j] != " " && i > 2 && j+3 < cols && grille.Board[i][j] == grille.Board[i-1][j+1] && grille.Board[i-1][j+1] == grille.Board[i-2][j+2] && grille.Board[i-2][j+2] == grille.Board[i-3][j+3]) ||
				(grille.Board[i][j] != " " && i > 2 && j-3 >= 0 && grille.Board[i][j] == grille.Board[i-1][j-1] && grille.Board[i-1][j-1] == grille.Board[i-2][j-2] && grille.Board[i-2][j-2] == grille.Board[i-3][j-3]) {
				caseColor := grille.Board[i][j]

				// Attribution de la victoire au joueur correspondant en fonction de la couleur
				switch caseColor {
				case players[0].Color:
					winner = players[0]
				case players[1].Color:
					winner = players[1]
				}
			}
		}
	}

	// Vérification égalité
	var full = 0
	for i := 0; i < cols; i++ {
		if grille.Board[0][i] != " " {
			full++
		}
	}

	if full == cols {
		winner = Player{"Egalité", "None", true}
	}

	return winner
}

func GamePlay(grille *Grille, col int, players [2]Player, currentTurn *int, isEvenTurn *bool) Player {

	var player Player
	if *currentTurn%2 != 0 {
		player = players[0]
	} else {
		player = players[1]
	}

	if !addPion(grille, col, player) {
		// fmt.Println("Colonne pleine")
		return Player{"fullCol", "None", false}
	}

	*currentTurn++
	*isEvenTurn = !*isEvenTurn

	winner := checkWin(grille, players)

	if winner.Name != "" {
		return winner
	}

	return Player{}
}
