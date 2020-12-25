# Schedule #

[![Build Status](https://travis-ci.org/fireqong/schedule.svg?branch=master)](https://travis-ci.org/fireqong/schedule)
[![GoDoc](https://godoc.org/github.com/fireqong/schedule?status.svg)](https://godoc.org/github.com/fireqong/schedule)
[![Go Report](https://goreportcard.com/badge/github.com/fireqong/schedule)](https://goreportcard.com/report/github.com/fireqong/schedule)

## Install ##

Use `go get` to install this library.

    go get github.com/fireqong/schedule
    
## API document ##

See [GoDoc](https://godoc.org/github.com/fireqong/schedule) for full document.

## Example ##

```go
// Fetchall fetches URLs in parallel and reports their times and sizes.
package main

import "github.com/fireqong/schedule"

func main() {
	s := &schedule.Schedule{
		LogFile: "test.log",
	}

	s.AddTask(&schedule.Task{Rate: schedule.EverySecond, Cmd: []string{"php", "-r", "file_put_contents('test.txt', 1, FILE_APPEND);"}})
	if err := s.Run(); err != nil {
		panic(err.Error())
	}
}
```

## License ##

This library is licensed under MIT license. See LICENSE for details.