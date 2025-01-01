package scoreboard

import (
	"log"
)

// LoggingMiddleware logs details about events.
func LoggingMiddleware(event Event, next func(event Event) error) error {
	err := next(event)

	switch {
	case err != nil:
		log.Printf("Can't handle %s event: %v\n", event.Type, err)
	case event.Type == StartEvent:
		log.Printf("Started game: %s vs %s\n", event.HomeTeam, event.AwayTeam)
	case event.Type == UpdateEvent:
		log.Printf("Updated score: %s vs %s - %d:%d\n", event.HomeTeam, event.AwayTeam, event.HomeScore, event.AwayScore)
	case event.Type == FinishEvent:
		log.Printf("Finished game: %s vs %s - %d:%d\n", event.HomeTeam, event.AwayTeam, event.HomeScore, event.AwayScore)
	}

	return err
}
