// Package ui provides simple interactive user interface for score board.
// It's designed for demonstration purposes, not to be production-ready.
package ui

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/MichalMitros/sportradar-task/scoreboard"
)

// page is current state of UI.
type page string

const (
	homePage       page = "home"   // Home page
	newGamePage    page = "new"    // Add new game page
	updateGamePage page = "update" // Update game page
	finishGame     page = "finish" // Finish game page
)

// UserInterface is very simple console user interface.
type UserInterface struct {
	scoreBoard  *scoreboard.ScoreBoard
	currentPage page
	reader      *bufio.Reader
}

// New returns new user interface with provided score board.
func New(sb *scoreboard.ScoreBoard) *UserInterface {
	return &UserInterface{
		scoreBoard:  sb,
		currentPage: homePage,
		reader:      bufio.NewReader(os.Stdin),
	}
}

// Show shows user interface.
func (ui *UserInterface) Show() {
	switch ui.currentPage {
	case homePage:
		ui.showMainPage()
	case newGamePage:
		ui.showAddGamePage()
	case updateGamePage:
		ui.showUpdateGamePage()
	case finishGame:
		ui.showFinishGamePage()
	}
}

func (ui *UserInterface) showMainPage() {
	fmt.Println()
	fmt.Print(
		`MENU:
1. Games summary
2. Add game
3. Update scores
4. Finish game
5. Exit
Choose option (1, 2, 3, 4 or 5): `,
	)

	option := readPositiveNumber(ui.reader)

	switch option {
	case 1:
		fmt.Println(ui.scoreBoard.GetSummary())
	case 2:
		ui.currentPage = newGamePage
	case 3:
		ui.currentPage = updateGamePage
	case 4:
		ui.currentPage = finishGame
	case 5:
		os.Exit(0)
	default:
		fmt.Println("\nUnknown option, please try again.")
	}
}

func (ui *UserInterface) showAddGamePage() {
	fmt.Print("Home team name: ")
	homeTeam := readString(ui.reader)
	fmt.Print("Away team name: ")
	awayTeam := readString(ui.reader)

	err := ui.scoreBoard.HandleEvent(scoreboard.Event{
		Type:     scoreboard.StartEvent,
		HomeTeam: homeTeam,
		AwayTeam: awayTeam,
	})
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}

	ui.currentPage = homePage
}

func (ui *UserInterface) showUpdateGamePage() {
	fmt.Print("Home team name: ")
	homeTeam := readString(ui.reader)
	fmt.Print("Away team name: ")
	awayTeam := readString(ui.reader)

	fmt.Print("Home team score: ")
	homeScore := readPositiveNumber(ui.reader)
	fmt.Print("Away team score: ")
	awayScore := readPositiveNumber(ui.reader)

	err := ui.scoreBoard.HandleEvent(scoreboard.Event{
		Type:      scoreboard.UpdateEvent,
		HomeTeam:  homeTeam,
		AwayTeam:  awayTeam,
		HomeScore: homeScore,
		AwayScore: awayScore,
	})
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}

	ui.currentPage = homePage
}

func (ui *UserInterface) showFinishGamePage() {
	fmt.Print("Home team name: ")
	homeTeam := readString(ui.reader)
	fmt.Print("Away team name: ")
	awayTeam := readString(ui.reader)

	err := ui.scoreBoard.HandleEvent(scoreboard.Event{
		Type:     scoreboard.FinishEvent,
		HomeTeam: homeTeam,
		AwayTeam: awayTeam,
	})
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}

	ui.currentPage = homePage
}

func readString(reader *bufio.Reader) string {
	for {
		reader.Reset(os.Stdin)

		if input, err := reader.ReadString('\n'); err == nil {
			input = strings.TrimSpace(input)
			return input
		}
		fmt.Print("Can't read provided text. Please try again: ")
	}
}

func readPositiveNumber(reader *bufio.Reader) int {
	for {
		reader.Reset(os.Stdin)

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Print("Can't read provided text. Please try again: ")
			continue
		}
		input = strings.TrimSpace(input)

		value, err := strconv.Atoi(input)
		if err != nil || value < 0 {
			fmt.Printf("\"%s\" is not a positive number! Please try again:", input)
			continue
		}

		return value
	}
}
