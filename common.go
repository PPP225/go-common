package common

import (
	"encoding/json"
	"fmt"
	"log"
	"regexp"
	"runtime"
	"time"
)

// Check check for error and panics, if not nil
func Check(e error) {
	if e != nil {
		panic(e)
	}
}

// Log writes to stdout if shoudLog flag is true
func Log(shoudLog bool, format string, a ...interface{}) {
	if shoudLog {
		fmt.Printf(format, a...)
	}
}

// ErrorTraced mimics panic() output.
//
// Usage: return ErrorTraced(err, inputAgrs...)
func ErrorTraced(prev error, args ...string) error {
	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	return fmt.Errorf("%s(%v)\n\t%s:%d\n%s", frame.Function, args, frame.File, frame.Line, prev)
	//log.Printf("%s:%d %s\n", frame.File, frame.Line, frame.Function)
}

// Trace logs current code line trace
func Trace() {
	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	log.Printf("%s,:%d %s\n", frame.File, frame.Line, frame.Function)
}

// CheckEmail checks e-mail format correctness
func CheckEmail(email string) (ok bool) {
	reEmail := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return reEmail.MatchString(email)
}

// ElapsedMs returns elapsed time between now and arg as ms
func ElapsedMs(start time.Time) float64 {
	return (float64)(time.Now().Sub(start).Nanoseconds()) / 1e6
}

// PrettyPrint prints objects prettier
func PrettyPrint(v interface{}) (err error) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err == nil {
		fmt.Println(string(b))
	}
	return
}

// TODO: add retry function https://stackoverflow.com/questions/47606761/repeat-code-if-an-error-occured
