package workers

import (
	"LANG_KENNY_GO25/basics"
)

// Call the main functions that allow the program to work
func Execute() {
	// Gets the flags
	f := basics.FlagsManager()

	// Displays parameters
	basics.MotdDisp(f)

	// Shows that the work begins
	basics.WorkStartDisp()

	// Does the work
	worker(f)

	// Shows that the work is finished
	basics.WorkEndDisp()
}
