package totalamountofpoints

import (
	"strconv"
	"strings"
)

func Points(games []string) int {

	var points int

	for i := 0; i < len(games); i++ {
		teamScores := strings.Split(games[i], ":")
		points += countOfPoints(teamScores)
	}

	return points
}

func countOfPoints(teamScores []string) int {
	const (
		win  = 3
		loss = 0
		tie  = 1
	)

	firstTeamScore, _ := strconv.Atoi(teamScores[0])
	secondTeamScore, _ := strconv.Atoi(teamScores[1])

	if firstTeamScore > secondTeamScore {
		return win
	} else if firstTeamScore < secondTeamScore {
		return loss
	} else {
		return tie
	}

}
