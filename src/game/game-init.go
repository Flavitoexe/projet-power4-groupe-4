package game

import (
	"errors"
	"regexp"
	"strings"
)

func traitementPlayerInit(name string) error {

	name = strings.TrimSpace(name)
	reg := regexp.MustCompile(`^[A-Za-z]{4,20}$`)

	if !reg.MatchString(name) {
		return errors.New("nom invalide : 4 à 20 lettres uniquement (ex : Bobby)")
	}

	return nil
}
