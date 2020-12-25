package schedule

import (
	"testing"
)

func TestSchedule_Run(t *testing.T) {
	s := &Schedule{LogFile: "crontab.log"}
	s.AddTask(&Task{Rate: EverySecond * 4, Cmd: []string{"php", "-r", "123"}})
	s.AddTask(&Task{Rate: EverySecond, Cmd: []string{"php", "-r", "echo 3 . PHP_EOL;"}})
	if err := s.Run(); err != nil {
		panic(err.Error())
	}
}
