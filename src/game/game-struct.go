package game

// Structure joueur
type Player struct {
	Name   string
	Color  string
	HasWon bool
}

// Structure état du jeu pour gérer la partie en cours
type GameState struct {
	Grid        Grille
	Players     [2]Player
	CurrentTurn int
	IsEvenTurn  bool
	Error       string
}

// Structure résultat de la partie pour le scoreboard
type GameResult struct {
	Players  [2]Player
	Winner   Player
	Date     string
	NbrTours int
}
