package basics

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func FlagsManager() (d, t string, w int, q bool) {
	var wordList, targetURL string
	var workersNumb int
	var quietFlag, help bool

	flag.StringVar(&wordList, "d", "", "Path to dictionnary file (required)")
	flag.StringVar(&targetURL, "t", "", "Target to enumerate (required)")
	flag.IntVar(&workersNumb, "w", 1, "Number of workers to run")
	flag.BoolVar(&quietFlag, "q", false, "When set to true, only show HTTP 200")
	flag.BoolVar(&help, "h", false, "Display this help")

	flag.Parse()

	flagsCheck(help, wordList, targetURL)

	targetURL = targetURLBuilder(targetURL)

	return wordList, targetURL, workersNumb, quietFlag
}

func flagsCheck(help bool, wordList, targetURL string) {
	if help {
		fmt.Println("Usage of mygb:")
		flag.PrintDefaults()
		os.Exit(0)
	}

	if flag.NFlag() == 0 {
		fmt.Println("No flag has been provided (Use -h for help)")
		os.Exit(1)
	}

	if wordList == "" || targetURL == "" {
		fmt.Println("-d and -t flags are required (Use -h for help)")
		os.Exit(1)
	}

	if _, err := os.Stat(wordList); os.IsNotExist(err) {
		fmt.Println("Can't find", wordList)
		os.Exit(1)
	}
}

func targetURLBuilder(oldTargetURL string) (newTargetURL string) {
	if strings.HasPrefix(oldTargetURL, "http") {
		newTargetURL = oldTargetURL
	} else {
		newTargetURL = "http://" + oldTargetURL
	}
	return newTargetURL
}
