package main

import (
	"fmt"
	"os"

	"github.com/MichalMitros/sportradar-task/scoreboard"
	"github.com/MichalMitros/sportradar-task/ui"
)

//nolint:gochecknoglobals // it's ok for simple demonstration purposes
var commands map[string]struct {
	Description string
	Behavior    func()
}

func init() {
	commands = map[string]struct {
		Description string
		Behavior    func()
	}{
		"help": {
			Description: "Show available commands.",
			Behavior:    printHelp,
		},
		"ui": {
			Description: "Run score board with user interface. (default)",
			Behavior:    runUI,
		},
		"demo": {
			Description: "Run simple demo without user interface.",
			Behavior:    runDemo,
		},
	}
}

func main() {
	if len(os.Args) < 2 {
		commands["ui"].Behavior()
		return
	}

	for cmd, action := range commands {
		if os.Args[1] == cmd {
			action.Behavior()
			return
		}
	}

	printHelp()
}

// printHelp prints commands and their descriptions.
func printHelp() {
	fmt.Println("Available commands:")
	for cmd, action := range commands {
		fmt.Printf("%s\t-\t%s\n", cmd, action.Description)
	}
}

// runUI runs score board in interactive mode with simple user interface.
func runUI() {
	scoreBoard := scoreboard.New()
	scoreBoard.Use(scoreboard.LoggingMiddleware)

	view := ui.New(scoreBoard)
	for {
		view.Show()
	}
}

// runDemo runs simple demo with some mocked events for quick demonstration.
func runDemo() {
	scoreBoard := scoreboard.New()
	scoreBoard.Use(scoreboard.LoggingMiddleware)

	// Errors from HandleEvent will be logged and doesn't need to be handled here.
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
