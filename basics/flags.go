package basics

import (
	"flag"
	"os"
	"strings"
)

func FlagsManager() (d, t string, w, e int, q bool) {
	var wordList, targetURL string
	var workers, expiration int
	var quiet, help, clear bool

	flag.StringVar(&wordList, "d", "", "Path to dictionnary file (required)")
	flag.StringVar(&targetURL, "t", "", "Target to enumerate (required)")
	flag.IntVar(&workers, "w", 1, "Number of workers to run")
	flag.IntVar(&expiration, "e", 10, "Time before HTTP request expiration in seconds")
	flag.BoolVar(&quiet, "q", false, "When set to true, only show HTTP 200")
	flag.BoolVar(&clear, "c", false, "Clears terminal at startup (Only for Linux)")
	flag.BoolVar(&help, "h", false, "Display this help")

	flag.Parse()

	flagsCheck(help, clear, wordList, targetURL, workers, expiration)

	targetURL = targetURLBuilder(targetURL)

	return wordList, targetURL, workers, expiration, quiet
}

func flagsCheck(help, clear bool, wordList, targetURL string, workers, expiration int) {
	if help {
		helpDisp()
		os.Exit(0)
	}

	if flag.NFlag() == 0 {
		noFlagErr()
		os.Exit(1)
	}

	if wordList == "" || targetURL == "" {
		dtFlagsErr()
		os.Exit(1)
	}

	if _, err := os.Stat(wordList); os.IsNotExist(err) {
		wordListErr()
		os.Exit(1)
	}

	if workers < 1 {
		workersErr()
		os.Exit(1)
	}

	if expiration < 1 {
		expirationErr()
		os.Exit(1)
	}

	if clear {
		clearTerm()
	}
}

func targetURLBuilder(oldTargetURL string) (newTargetURL string) {
	var httpTargetURL string

	if strings.HasPrefix(oldTargetURL, "http") {
		httpTargetURL = oldTargetURL
	} else {
		httpTargetURL = "http://" + oldTargetURL
	}

	if strings.HasSuffix(httpTargetURL, "/") {
		newTargetURL = httpTargetURL
	} else {
		newTargetURL = httpTargetURL + "/"
	}

	return newTargetURL
}
