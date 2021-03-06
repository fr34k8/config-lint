package assertion

import "fmt"

var (
	isDebug = false
)

// SetDebug turns verbose logging on or off
func SetDebug(b bool) {
	isDebug = b
}

// Debugf prints a formatted string when verbose logging is turned on
func Debugf(format string, args ...interface{}) {
	if isDebug == false {
		return
	}
	fmt.Printf(format, args...)
}
