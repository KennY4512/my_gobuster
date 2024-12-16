package basics

import (
	"os"
	"os/exec"
	"time"
)

func clearTerm() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

type Timer struct {
	initTime time.Time
	duration time.Duration
}

func (t *Timer) init() {
	t.initTime = time.Now()
}

func (t *Timer) getDuration() time.Duration {
	t.duration = time.Since(t.initTime).Round(time.Millisecond)
	return t.duration
}

func FileOpener(wordList string) (file *os.File) {
	file, err := os.Open(wordList)

	if err != nil {
		fileOpenerErr(err)
		defer file.Close()
		os.Exit(1)
	}

	return file
}
