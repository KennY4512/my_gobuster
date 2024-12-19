package basics

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"time"
)

// Structure implementing flags
type Flags struct {
	WordList, TargetURL string
	Workers, Expiration int
	Quiet, Csv          bool
}

// Structure implementing a timer
type Timer struct {
	initTime time.Time
}

func (timer *Timer) init() {
	timer.initTime = time.Now()
}

func (timer Timer) getDuration() time.Duration {
	return time.Since(timer.initTime).Round(time.Millisecond)
}

// Structure implementing a CSV file creator
type CSV struct {
	dataSlice []string
}

func (c *CSV) Add(word string, httpCode int) {
	c.dataSlice = append(c.dataSlice, word+","+fmt.Sprintf("%d", httpCode)) // Format the line before adding it to a temporary slice
}

func (c CSV) Write() {

	// Creates a CSV file whose name is the current timestamp
	file, err := os.Create("mygb_" + time.Now().Format("20060102_150405") + ".csv")
	if err != nil {
		fileCreateErr(err)
		return
	}
	defer file.Close()

	// Creats a new writter (buffer)
	writer := bufio.NewWriter(file)

	// Adds CSV header to the buffer
	_, err = writer.WriteString("word,httpcode\n")
	if err != nil {
		fileWriteErr(err)
		return
	}

	// Adds the content of the slice to the buffer
	for _, line := range c.dataSlice {
		_, err = writer.WriteString(line + "\n")
		if err != nil {
			fileWriteErr(err)
			return
		}
	}

	// Writes buffer content to csv file
	err = writer.Flush()
	if err != nil {
		fileFlushErr(err)
		return
	}
}

// Clears the terminal
func clearTerm() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// Opens a given file correctly
func FileOpener(wordList string) (file *os.File) {
	file, err := os.Open(wordList)
	if err != nil {
		fileOpenErr(err)
		defer file.Close()
		os.Exit(1)
	}
	return file
}
