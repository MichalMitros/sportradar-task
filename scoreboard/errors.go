package scoreboard

import "errors"

// ErrUnknownEventType is returned when event type is unknown.
var ErrUnknownEventType = errors.New("unknown event type")

// ErrGameAlreadyOngoing is returned if a game provided in event is already ongoing.
var ErrGameAlreadyOngoing = errors.New("the game is already ongoing")

// ErrGameNotFound is returned when event changes state of game which is not ongoing.
var ErrGameNotFound = errors.New("game not found")

// ErrNegativeScore is returned when event contains negative scores.
var ErrNegativeScore = errors.New("score cannot be negative")

// ErrHomeTeamAlreadyPlaying is returned if home team is already tracked on scoreboard.
var ErrHomeTeamAlreadyPlaying = errors.New("home team already playing")

// ErrAwayTeamAlreadyPlaying is returned if away team is already tracked on scoreboard.
var ErrAwayTeamAlreadyPlaying = errors.New("away team already playing")
