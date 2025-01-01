// Package scoreboard implements score board for storing games and scores
// and handling events updating their state.
package scoreboard

import (
	"fmt"
	"sort"
	"strings"
	"time"
)

// EventType represents types of events.
type EventType string

const (
	// StartEvent starts new game, requires home and away teams names.
	StartEvent EventType = "start"
	// FinishEvent finishes ongoing game, requires home and away teams names.
	FinishEvent EventType = "finish"
	// UpdateEvent updates ongoing game's scores, requires home and away teams names and scores.
	UpdateEvent EventType = "update"
)

// Match is single ongoing match with scores.
type Match struct {
	homeTeam  string
	awayTeam  string
	homeScore int
	awayScore int
	startTime time.Time
}

// Event modifies state of Match.
type Event struct {
	Type      EventType
	HomeTeam  string
	AwayTeam  string
	HomeScore int
	AwayScore int
}

// Middleware adds functionalities to events handling.
type Middleware func(event Event, next func(event Event) error) error

// ScoreBoard stores matches and handles events.
type ScoreBoard struct {
	matches     map[string]*Match
	middlewares []Middleware
}

// New returns new ScoreBoard.
func New() *ScoreBoard {
	return &ScoreBoard{
		matches:     make(map[string]*Match),
		middlewares: []Middleware{},
	}
}

// Use adds middleware to ScoreBoard.
func (s *ScoreBoard) Use(middleware Middleware) {
	s.middlewares = append(s.middlewares, middleware)
}

// HandleEvent handles events and applies middlewares.
func (s *ScoreBoard) HandleEvent(event Event) error {
	chain := s.buildMiddlewareChain()
	return chain(event)
}

func (s *ScoreBoard) buildMiddlewareChain() func(event Event) error {
	handler := func(event Event) error {
		return s.processEvent(event)
	}

	for i := len(s.middlewares) - 1; i >= 0; i-- {
		currentMiddleware := s.middlewares[i]
		next := handler
		handler = func(event Event) error {
			return currentMiddleware(event, next)
		}
	}

	return handler
}

func (s *ScoreBoard) processEvent(event Event) error {
	switch event.Type {
	case StartEvent:
		return s.startGame(event.HomeTeam, event.AwayTeam)
	case FinishEvent:
		return s.finishGame(event.HomeTeam, event.AwayTeam)
	case UpdateEvent:
		return s.updateScore(event.HomeTeam, event.AwayTeam, event.HomeScore, event.AwayScore)
	default:
		return ErrUnknownEventType
	}
}

// startGame starts new game.
func (s *ScoreBoard) startGame(homeTeam, awayTeam string) error {
	key := buildGameKey(homeTeam, awayTeam)
	if _, exists := s.matches[key]; exists {
		return ErrGameAlreadyOngoing
	}

	// Check if teams are not currently playing.
	for _, match := range s.matches {
		if homeTeam == match.homeTeam || homeTeam == match.awayTeam {
			return ErrHomeTeamAlreadyPlaying
		}
		if awayTeam == match.homeTeam || awayTeam == match.awayTeam {
			return ErrAwayTeamAlreadyPlaying
		}
	}

	s.matches[key] = &Match{
		homeTeam:  homeTeam,
		awayTeam:  awayTeam,
		homeScore: 0,
		awayScore: 0,
		startTime: time.Now(),
	}

	return nil
}

// finishGame finishes ongoing game.
func (s *ScoreBoard) finishGame(homeTeam, awayTeam string) error {
	key := buildGameKey(homeTeam, awayTeam)
	if _, exists := s.matches[key]; !exists {
		return ErrGameNotFound
	}
	delete(s.matches, key)

	return nil
}

// updateScore updates scores of ongoing game.
func (s *ScoreBoard) updateScore(homeTeam, awayTeam string, homeScore, awayScore int) error {
	if homeScore < 0 || awayScore < 0 {
		return ErrNegativeScore
	}

	key := buildGameKey(homeTeam, awayTeam)
	match, exists := s.matches[key]
	if !exists {
		return ErrGameNotFound
	}
	match.homeScore = homeScore
	match.awayScore = awayScore

	return nil
}

// GetSummary returns summary of ScoreBoard as string describing scores in ongoing games.
func (s *ScoreBoard) GetSummary() string {
	matches := make([]*Match, 0, len(s.matches))
	for _, match := range s.matches {
		if match != nil {
			matches = append(matches, match)
		}
	}

	sort.SliceStable(matches, func(i, j int) bool {
		totalScoreI := matches[i].homeScore + matches[i].awayScore
		totalScoreJ := matches[j].homeScore + matches[j].awayScore
		if totalScoreI == totalScoreJ {
			return matches[i].startTime.Before(matches[j].startTime)
		}
		return totalScoreI > totalScoreJ
	})

	var summary strings.Builder
	// WriteString never retruns any error.
	_, _ = summary.WriteString("Game Summary:\n")

	for ix, match := range matches {
		_, _ = summary.WriteString(
			fmt.Sprintf("%d. %s %d - %s %d\n", ix+1, match.homeTeam, match.homeScore, match.awayTeam, match.awayScore),
		)
	}

	return summary.String()
}

// buildGameKey builds map key for the game.
func buildGameKey(homeTeam, awayTeam string) string {
	return fmt.Sprintf("%s vs %s", homeTeam, awayTeam)
}
