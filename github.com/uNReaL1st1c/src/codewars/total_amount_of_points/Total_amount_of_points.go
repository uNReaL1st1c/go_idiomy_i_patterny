package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	games := make([]string, 0, 10)
	games = append(games, "2:2")
	games = append(games, "3:1")
	games = append(games, "3:2")
	games = append(games, "2:2")
	games = append(games, "1:1")
	games = append(games, "0:3")
	games = append(games, "1:2")

	fmt.Println(Points(games))

}

func Points(games []string) int {

	var points int

	for i := 0; i < len(games); i++ {
		teamScores := strings.Split(games[i], ":")
		points += CountOfPoints(teamScores)
	}

	return points
}

func CountOfPoints(teamScores []string) int {
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
