package scoreboard_test

import (
	"testing"

	"github.com/MichalMitros/sportradar-task/scoreboard"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUnitNew(t *testing.T) {
	assert.NotNil(t, scoreboard.New(), "should return new ScoreBoard")
}

func TestUnitHandleEvent(t *testing.T) {
	const unknownEventType scoreboard.EventType = "unknown"
	tests := map[string]struct {
		events         []scoreboard.Event
		expectedErrors []error // must be the same length as events.
	}{
		"game not found error": {
			events: []scoreboard.Event{
				{Type: scoreboard.UpdateEvent, HomeTeam: "A", AwayTeam: "B"},
				{Type: scoreboard.FinishEvent, HomeTeam: "C", AwayTeam: "D"},
			},
			expectedErrors: []error{
				scoreboard.ErrGameNotFound,
				scoreboard.ErrGameNotFound,
			},
		},
		"negative score error": {
			events: []scoreboard.Event{
				{Type: scoreboard.StartEvent, HomeTeam: "A", AwayTeam: "B"},
				{Type: scoreboard.UpdateEvent, HomeTeam: "A", AwayTeam: "B", HomeScore: -1},
			},
			expectedErrors: []error{
				nil,
				scoreboard.ErrNegativeScore,
			},
		},
		"game already ongoing error": {
			events: []scoreboard.Event{
				{Type: scoreboard.StartEvent, HomeTeam: "A", AwayTeam: "B"},
				{Type: scoreboard.StartEvent, HomeTeam: "A", AwayTeam: "B"},
			},
			expectedErrors: []error{
				nil,
				scoreboard.ErrGameAlreadyOngoing,
			},
		},
		"home team already playing error": {
			events: []scoreboard.Event{
				{Type: scoreboard.StartEvent, HomeTeam: "A", AwayTeam: "B"},
				{Type: scoreboard.StartEvent, HomeTeam: "A", AwayTeam: "C"},
			},
			expectedErrors: []error{
				nil,
				scoreboard.ErrHomeTeamAlreadyPlaying,
			},
		},
		"away team already playing error": {
			events: []scoreboard.Event{
				{Type: scoreboard.StartEvent, HomeTeam: "A", AwayTeam: "B"},
				{Type: scoreboard.StartEvent, HomeTeam: "C", AwayTeam: "B"},
			},
			expectedErrors: []error{
				nil,
				scoreboard.ErrAwayTeamAlreadyPlaying,
			},
		},
		"unknown event type": {
			events: []scoreboard.Event{
				{Type: unknownEventType, HomeTeam: "A", AwayTeam: "B"},
			},
			expectedErrors: []error{
				scoreboard.ErrUnknownEventType,
			},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			scoreBoard := scoreboard.New()
			events, errors, middleware := testMiddleware(t)
			scoreBoard.Use(middleware)

			for _, event := range tt.events {
				// errors will be stored by middleware - no need to check them here.
				_ = scoreBoard.HandleEvent(event)
			}

			assert.Equal(t, tt.events, *events, "should return correct events history")
			assert.Equal(t, tt.expectedErrors, *errors, "should return correct errors history")
		})
	}
}

func TestUnitGetSummary(t *testing.T) {
	tests := map[string]struct {
		events   []scoreboard.Event
		expected string
	}{
		"no games": {
			expected: "Game Summary:\n",
		},
		"no ongoing games": {
			events: []scoreboard.Event{
				{
					Type:     scoreboard.StartEvent,
					HomeTeam: "A",
					AwayTeam: "B",
				},
				{
					Type:     scoreboard.FinishEvent,
					HomeTeam: "A",
					AwayTeam: "B",
				},
			},
			expected: "Game Summary:\n",
		},
		"many games": {
			events: []scoreboard.Event{
				{Type: scoreboard.StartEvent, HomeTeam: "A", AwayTeam: "B"},
				{Type: scoreboard.StartEvent, HomeTeam: "C", AwayTeam: "D"},
				{Type: scoreboard.UpdateEvent, HomeTeam: "A", AwayTeam: "B", HomeScore: 1, AwayScore: 2},
				{Type: scoreboard.StartEvent, HomeTeam: "E", AwayTeam: "F"},
				{Type: scoreboard.UpdateEvent, HomeTeam: "E", AwayTeam: "F", HomeScore: 1, AwayScore: 2},
				{Type: scoreboard.UpdateEvent, HomeTeam: "C", AwayTeam: "D", HomeScore: 2, AwayScore: 3},
			},
			expected: "Game Summary:\n1. C 2 - D 3\n2. A 1 - B 2\n3. E 1 - F 2\n",
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			scoreBoard := scoreboard.New()

			for _, event := range tt.events {
				err := scoreBoard.HandleEvent(event)
				require.NoError(t, err)
			}

			assert.Equal(t, tt.expected, scoreBoard.GetSummary(), "should return correct summary")
		})
	}
}

// testMiddleware return test middleware with history of events and returned errors (also nils) updated for each event.
func testMiddleware(t *testing.T) (*[]scoreboard.Event, *[]error, scoreboard.Middleware) {
	t.Helper()
	events := []scoreboard.Event{}
	errors := []error{}

	return &events, &errors, func(event scoreboard.Event, next func(event scoreboard.Event) error) error {
		events = append(events, event)
		err := next(event)
		errors = append(errors, err)
		return err
	}
}
