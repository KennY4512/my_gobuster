package basics

import (
	"flag"
	"fmt"
)

// Timer object to handle time counter
var timer Timer

// Variables for color display
var (
	red   = "\033[31m"
	green = "\033[32m"
	reset = "\033[0m"
)

// Basic display functions
func MotdDisp(f Flags) {
	fmt.Printf("Starting MyGB\n\n")
	fmt.Println("---")
	fmt.Println("Target:", f.TargetURL)
	fmt.Println("List:", f.WordList)
	fmt.Println("Workers:", f.Workers)
	fmt.Printf("Timeout: %ds\n", f.Expiration)
	fmt.Println("Quiet mode:", f.Quiet)
	fmt.Println("CSV output:", f.Csv)
	fmt.Printf("---\n\n")
}

func WorkStartDisp() {
	fmt.Printf("Starting scan...\n\n")
	timer.init()
}

func WorkEndDisp() {
	fmt.Printf("\nScan done in %s\n", timer.getDuration())
}

func HTTPDisp(word string, code int) {
	if code == 200 {
		fmt.Printf(green+"/%-15s %d\n"+reset, word, code)
	} else {
		fmt.Printf(red+"/%-15s %d\n"+reset, word, code)
	}
}

func helpDisp() {
	fmt.Println("Usage of mygb:")
	flag.PrintDefaults()
}

// Error display functions
func HTTPErr(word string, err error) {
	fmt.Printf("Request error for /%s: %v\n", word, err)
}

func FileReadErr(err error) {
	fmt.Println("Error reading file:", err)
}

func fileOpenErr(err error) {
	fmt.Println("Error opening file:", err)
}

func fileCreateErr(err error) {
	fmt.Println("Error creating file:", err)
}

func fileWriteErr(err error) {
	fmt.Println("Error writing file:", err)
}

func fileFlushErr(err error) {
	fmt.Println("Error writing buffer:", err)
}

func noFlagErr() {
	fmt.Println("No flag has been provided (Use -h for help)")
}

func dtFlagsErr() {
	fmt.Println("-d and -t flags are required (Use -h for help)")
}

func wordListErr() {
	fmt.Println("Word list not found")
}

func workersErr() {
	fmt.Println("Workers number too low (Use -h for help)")
}

func expirationErr() {
	fmt.Println("Expiration time too short (Use -h for help)")
}
