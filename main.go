package main

import (
	"fmt"

	"github.com/MichalMitros/sportradar-task/scoreboard"
)

func main() {
	scoreBoard := scoreboard.New()
	scoreBoard.Use(scoreboard.LoggingMiddleware)

	_ = scoreBoard.HandleEvent(scoreboard.Event{Type: scoreboard.StartEvent, HomeTeam: "Mexico", AwayTeam: "Canada"})
	_ = scoreBoard.HandleEvent(scoreboard.Event{Type: scoreboard.StartEvent, HomeTeam: "Spain", AwayTeam: "Brazil"})
	_ = scoreBoard.HandleEvent(scoreboard.Event{Type: scoreboard.StartEvent, HomeTeam: "Germany", AwayTeam: "France"})
	_ = scoreBoard.HandleEvent(scoreboard.Event{Type: scoreboard.StartEvent, HomeTeam: "Uruguay", AwayTeam: "Italy"})
	_ = scoreBoard.HandleEvent(scoreboard.Event{Type: scoreboard.StartEvent, HomeTeam: "Austria", AwayTeam: "Australia"})

	_ = scoreBoard.HandleEvent(
		scoreboard.Event{Type: scoreboard.UpdateEvent, HomeTeam: "Mexico", AwayTeam: "Canada", HomeScore: 0, AwayScore: 5},
	)
	_ = scoreBoard.HandleEvent(
		scoreboard.Event{Type: scoreboard.UpdateEvent, HomeTeam: "Spain", AwayTeam: "Brazil", HomeScore: 10, AwayScore: 2},
	)
	_ = scoreBoard.HandleEvent(
		scoreboard.Event{Type: scoreboard.UpdateEvent, HomeTeam: "Germany", AwayTeam: "France", HomeScore: 2, AwayScore: 2},
	)
	_ = scoreBoard.HandleEvent(
		scoreboard.Event{Type: scoreboard.UpdateEvent, HomeTeam: "Uruguay", AwayTeam: "Italy", HomeScore: 6, AwayScore: 6},
	)
	_ = scoreBoard.HandleEvent(scoreboard.Event{
		Type:      scoreboard.UpdateEvent,
		HomeTeam:  "Austria",
		AwayTeam:  "Australia",
		HomeScore: 3,
		AwayScore: 1,
	})

	fmt.Println(scoreBoard.GetSummary())

	_ = scoreBoard.HandleEvent(scoreboard.Event{Type: scoreboard.FinishEvent, HomeTeam: "Germany", AwayTeam: "France"})

	fmt.Println(scoreBoard.GetSummary())
}
