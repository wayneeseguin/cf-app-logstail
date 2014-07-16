package main

import (
	"fmt"
	"github.com/ActiveState/tail"
	"os"
	"strings"
	"sync"
)

var logs []string

func init() {
	logstail := os.Getenv("logstail")
	if len(logstail) > 1 {
		logs = strings.Split(logstail, ":")
	}
}

func tailLog(logFile, stream string) {
	var s *os.File
	if stream == "2" {
		s = os.Stderr
	} else {
		s = os.Stdout
	}

	t, err := tail.TailFile(logFile, tail.Config{Follow: true})
	if err != nil {
		fmt.Printf("ERROR: Can not tail log file %s, err: %s", logFile, err)
	} else {
		for line := range t.Lines {
			fmt.Fprintf(s, line.Text+"\n")
			s.Sync()
		}
	}
}

func main() {
	var w sync.WaitGroup
	w.Add(len(logs))
	for _, log := range logs {
		split := strings.Split(log, "|")
		logFile := split[0]
		stream := "1"
		if len(split) > 1 {
			stream = split[1]
		}
		go func() {
			tailLog(logFile, stream)
			w.Done()
		}()
	}
	w.Wait()
}
