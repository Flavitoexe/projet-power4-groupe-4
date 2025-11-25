package game

import "time"

func GetScoreboard(partieF GameState) GameResult {
	var results GameResult

	now := time.Now()
	dateStr := now.Format("02-01-2006 15:04:05")
	results.Players = partieF.Players
	results.NbrTours = partieF.CurrentTurn

	if partieF.Players[0].HasWon {
		results.Winner = partieF.Players[0]
	} else if partieF.Players[1].HasWon {
		results.Winner = partieF.Players[1]
	} else {
		results.Winner = Player{"Egalité", "", false}
	}

	results.Date = dateStr
	return results
}
