package schedule

import (
	"time"
)

const (
	EverySecond = time.Second * 1
	EveryMinute = time.Minute
	EveryHour   = time.Hour
	EveryDay    = time.Hour * 24
	EveryWeek   = time.Hour * 24 * 7
)

type Task struct {
	Rate time.Duration
	Cmd  []string

	latestRunTime time.Time
	channel       <-chan time.Time
}
