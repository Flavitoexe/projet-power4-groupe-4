package game

import (
	"errors"
	"regexp"
	"strings"
)

type Player struct {
	Name  string
	Color string
}

func NewPlayer(name string, color string) Player {
	player := Player{name, color}
	return player
}

func TraitementPlayerInit(name string) error {

	name = strings.TrimSpace(name)
	reg := regexp.MustCompile(`^[A-Za-z]{4,20}$`)

	if !reg.MatchString(name) {
		return errors.New("nom invalide : 4 à 20 lettres uniquement (ex : Bobby)")
	}

	return nil
}
