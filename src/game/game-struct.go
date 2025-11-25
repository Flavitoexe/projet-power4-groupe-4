package game

type Player struct {
	Name   string
	Color  string
	HasWon bool
}

type GameState struct {
	Grid        Grille
	Players     [2]Player
	CurrentTurn int
	IsEvenTurn  bool
	Error       string
}

type GameResult struct {
	Players  [2]Player
	Winner   Player
	Date     string
	NbrTours int
}
