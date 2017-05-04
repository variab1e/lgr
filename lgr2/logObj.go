package lgr

import (
	"image/color"
	"io"
	"log"
)


type LoggerConfigI interface {
	GetFilters() []Filters
}

type LoggerConfig struct {
	Level Level
	Name  string
	//Prefix     string
	Prefix           []interface{}
	Writer           *io.Writer				// Writer is a pointer to a Writer that already exists in Output .Writer
	Logger           **log.Logger
	color            *color.Color
	PrintDebug       bool
	Flags            int
	BlockFilters     Filters
	HighlightFilters Filters
}

type Filters []Filter

type Filter struct {
	Keywords	[]string
	Level			int
}

type Output struct {
	Name		string
	Writer	io.Writer
	Filters	[]Filter
}

// defaults // move these //

var (
	defaultOutputs = []Output{
		


// Writer acts as a modifier pre-output for the logs.
// Here we can add additional information (such the function the log is in)
// or styling, such as coloration
func (lt LoggerConfig) Writer(p []byte) (n int, err error) {

// NOTE:
// THIS CURRENTLY ONLY IS DOING STDOUT!!
// I THINK YOU WANT MORE THAN THAT
//
// BEWARE OF COLOR, ON MULTIPLE-THREADS, CAN BE BAD NEWS
//  when multiple writers are writing to the same output
//  for example, the console, the output can become corrupt
//  most likely because of ANSI injection sequences conflicting
//  but I havent had time to investigate this yet

		// get function http://moazzam-khan.com/blog/golang-get-the-function-callers-name/
		// get calling function
		// callStack is an array of calling entities
		callStack := make([]uintptr, 1)
		// callerName is a string containing the calling functions name
		// this will be printed in the log message
		var callerName string
		// Skip 2 levels to get the caller
		if runtime.Callers(2, callStack) == 0 {
				// No caller found
				callerName = "****NOT*FOUND****"
		}
 
		caller := runtime.FuncForPC(callStack[0]-1)
		if caller == nil {
				// caller was nil
				callerName = "nil"
		}
 
		// Print the file name and line number
		fileNameLine := caller.FileLine(callStack[0]-1)
 
		// Print the name of the function
		callerName = caller.Name()

		// split the received message on colons
		var str string = string(p[:])
		var strs []string = strings.SplitN(str,":",6)
		var msg string = str
		// the first 5 of which are time, etc in a normal message
		if len(strs) >= 6 {
				// but the 6th is the message type, ie. DEBUG
				// which can be used to map back to the logger
				msg = strs[5]
		} 
		// and now we can check if it should be colorized, etc.
		if !lt.PrintDebug { 
				lt.color.Print(msg)
		} else {
				lt.color.Print(str)
		}
		return len(p), nil
}

func (lt LoggerConfig) Write(p []byte) (n int, err error) {

func refreshLoggerConfigs(){
	// see log flag constants
	// https://golang.org/pkg/log/#pkg-constants
	for _, n := range LoggerConfigs {
				
				// if the log level is less than the outputThreshold (stdout)
				// and less than logThreshold (file output)
				// than don't log anything
		if n.Level < outputThreshold && n.Level < logThreshold {
			n.Handle = ioutil.Discard
		} else if n.Level >= outputThreshold && n.Level >= logThreshold {
			// if greater than or equal to both, log to both
			n.Handle = io.MultiWriter(FileHandle, n)
		} else if n.Level >= outputThreshold && n.Level < logThreshold {
			// if only outputThreshold is greater, only log to console
			n.Handle = n
		} else {
			// else (the only option remaining is logThreshold is greater)
			// log to FileLogger only
			n.Handle = FileHandle
		}
		*n.Logger = log.New(n.Handle, n.Prefix, n.Flags)
	}
}


func New() (logger LoggerConfig) {



	logger.Writer = 
}






//Filter lets you add Terms to the Filter
func (log *lgr) Filter() {

}

/**


3 log functions

Log -> internal
Printf -> Console , meant for user, web etc,
(Error? -> User, non crashing)
Critical -> System Crash

**/
