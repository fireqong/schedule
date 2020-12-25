package schedule

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"reflect"
	"time"
)

type Schedule struct {
	LogFile string

	tasks []*Task
}

func (s *Schedule) AddTask(task *Task) {
	s.tasks = append(s.tasks, task)
}

func (s *Schedule) Run() error {
	if s.LogFile == "" {
		return fmt.Errorf("set log file path first")
	}

	cases := make([]reflect.SelectCase, len(s.tasks))
	i := 0

	for _, task := range s.tasks {
		task.latestRunTime = time.Now()
		task.channel = time.Tick(task.Rate)
		cases[i] = reflect.SelectCase{Dir: reflect.SelectRecv, Chan: reflect.ValueOf(task.channel)}
		i += 1
	}

	for {
		chosen, _, _ := reflect.Select(cases)
		task := s.tasks[chosen]

		task.latestRunTime = time.Now()
		cmd := exec.Command(task.Cmd[0], task.Cmd[1:]...)
		err := cmd.Run()

		if err != nil {
			log := fmt.Sprintf("[%v] cmd %v faild.\r\n", time.Now().String(), task.Cmd)
			ioutil.WriteFile(s.LogFile, []byte(log), os.ModePerm)
		}

	}
}
