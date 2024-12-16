package workers

import (
	"LANG_KENNY_GO25/basics"
)

func Execute() {
	wordList, targetURL, workers, expiration, quiet := basics.FlagsManager()

	basics.MotdDisp(wordList, targetURL, workers, expiration, quiet)

	var t basics.Timer
	basics.StartDisp(&t)

	Worker(wordList, targetURL, workers, expiration, quiet)

	basics.EndDisp(&t)
}
