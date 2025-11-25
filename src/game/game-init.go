package game

import (
	"errors"
	"regexp"
	"strings"
)

func NewPlayer(name string, color string) Player {
	player := Player{name, color, false}
	return player
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

func TraitementPlayerInit(name string) error {

	name = strings.TrimSpace(name)
	reg := regexp.MustCompile(`^[A-Za-z]{4,20}$`)

	if !reg.MatchString(name) {
		return errors.New("nom invalide : 4 à 20 lettres uniquement (ex : Bobby)")
	}

	return nil
}
