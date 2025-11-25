package game

func GetScoreboard(partieF GameState, dateStr string) GameResult {
	var results GameResult
	results.Players = partieF.Players
	results.NbrTours = partieF.CurrentTurn
	if partieF.Players[0].HasWon {
		results.Winner = partieF.Players[0]
	} else if partieF.Players[1].HasWon {
		results.Winner = partieF.Players[1]
	} else {
		results.Winner = Player{"Eaglité", "", false}
	}
	results.Date = dateStr
	return results
}
