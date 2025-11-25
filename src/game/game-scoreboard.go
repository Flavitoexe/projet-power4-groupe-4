package game

import "time"

func GetScoreboard(partieF GameState) GameResult {
	var results GameResult

	// Récuperation et formatage de la date
	now := time.Now()
	dateStr := now.Format("02-01-2006 15:04:05")
	results.Date = dateStr

	results.Players = partieF.Players
	results.NbrTours = partieF.CurrentTurn

	// On vérifie l'issue de la partie
	if partieF.Players[0].HasWon {
		results.Winner = partieF.Players[0]
	} else if partieF.Players[1].HasWon {
		results.Winner = partieF.Players[1]
	} else {
		results.Winner = Player{"Egalité", "", false}
	}

	return results
}
