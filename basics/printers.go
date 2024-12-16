package basics

import (
	"flag"
	"fmt"
)

var (
	red   = "\033[31m"
	green = "\033[32m"
	reset = "\033[0m"
)

func MotdDisp(wordList, targetURL string, workers, expiration int, quiet bool) {
	fmt.Printf("Starting MyGB\n\n")
	fmt.Println("---")
	fmt.Println("Target:", targetURL)
	fmt.Println("List:", wordList)
	fmt.Println("Workers:", workers)
	fmt.Printf("Timeout: %ds\n", expiration)
	fmt.Println("Quiet mode:", quiet)
	fmt.Printf("---\n\n")
}

func StartDisp(t *Timer) {
	fmt.Printf("Starting scan...\n\n")
	t.init()
}

func EndDisp(t *Timer) {
	fmt.Printf("\nScan done in %s\n", t.getDuration())
}

func HTTPDisp(word string, code int) {
	if code == 200 {
		fmt.Printf(green+"/%-15s %d\n"+reset, word, code)
	} else {
		fmt.Printf(red+"/%-15s %d\n"+reset, word, code)
	}
}

func HTTPErr(word string, err error) {
	fmt.Printf("Request error /%s: %v\n", word, err)
}

func ReadingErr(err error) {
	fmt.Println("File reading error:", err)
}

func helpDisp() {
	fmt.Println("Usage of mygb:")
	flag.PrintDefaults()
}

func noFlagErr() {
	fmt.Println("No flag has been provided (Use -h for help)")
}

func dtFlagsErr() {
	fmt.Println("-d and -t flags are required (Use -h for help)")
}

func wordListErr() {
	fmt.Println("Can't find the specified word list")
}

func workersErr() {
	fmt.Println("Can't work with 0 workers (Use -h for help)")
}

func expirationErr() {
	fmt.Println("Expiration time too short (Use -h for help)")
}

func fileOpenerErr(err error) {
	fmt.Println("Error opening file:", err)
}
