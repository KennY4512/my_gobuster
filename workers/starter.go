package workers

import (
	"LANG_KENNY_GO25/basics"
)

// Call the main functions that allow the program to work
func Execute() {
	// Gets the flags
	wordList, targetURL, workers, expiration, quiet, csv := basics.FlagsManager()

	// Displays parameters
	basics.MotdDisp(wordList, targetURL, workers, expiration, quiet, csv)

	// Shows that the work begins
	basics.WorkStartDisp()

	// Does the work
	worker(wordList, targetURL, workers, expiration, quiet, csv)

	// Shows that the work is finished
	basics.WorkEndDisp()
}
