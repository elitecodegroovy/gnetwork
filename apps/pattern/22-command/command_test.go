package command

import (
	"strings"
	"testing"
	"time"
)

func TestCommand(t *testing.T) {
	var timeCommand Command
	timeCommand = &TimePassed{time.Now()}

	var helloCommand Command
	helloCommand = &HelloMessage{}

	time.Sleep(time.Second)

	t1 := timeCommand.Info()
	if !strings.Contains(t1, "s") {
		t.Fatal("helloCommand failed!")
	}
	msg := helloCommand.Info()
	if !strings.Contains(msg, "Hi, Golang!") {
		t.Fatal("helloCommand failed!")
	}
}
