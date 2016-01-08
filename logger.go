// Simple log replacement for go with Verbosity and Debug levels, and
// multi-stream output support.
//
// (C) by Marco Paganini <paganini AT paganini DOT net>

package logger

import (
	"fmt"
	"io"
	"log"
	"os"
)

// Logger represents a logger object
type Logger struct {
	*log.Logger
	verbose int
	debug   int
	streams []*os.File
}

// New Creates a new Logger instance.
func New(prefix string) *Logger {
	o := log.New(os.Stderr, prefix, 0)
	return &Logger{o, 0, 0, []*os.File{os.Stdout}}
}

// SetVerboseLevel sets the verbosity level for this log instance.
func (o *Logger) SetVerboseLevel(n int) {
	o.verbose = n
}

// SetDebugLevel sets the debugging level for this log instance.
func (o *Logger) SetDebugLevel(n int) {
	o.debug = n
}

// SetOutput sets the output streams to the list of writers.
func (o *Logger) SetOutput(streams []*os.File) {
	o.streams = streams
}

// writeString sends the string to all defined outputs.
func (o *Logger) writeString(s string) {
	for _, w := range o.streams {
		io.WriteString(w, s)
	}
}

// Verboseln prints the message to the output streams, followed by a newline,
// if the current verbose level is greater than or equal the specified
// verbosity level.
func (o *Logger) Verboseln(level int, v ...interface{}) {
	if o.verbose >= level {
		o.writeString(fmt.Sprintln(v...))
	}
}

// Verbosef uses the formatting string and variables to print a message to the
// output streams if the current verbose level is greater than or equal to the
// specified verbose level.
func (o *Logger) Verbosef(level int, format string, v ...interface{}) {
	if o.verbose >= level {
		o.writeString(fmt.Sprintf(format, v...))
	}
}

// Debugln prints the message to the output streams, followed by a newline,
// if the current debugging level is greater than or equal the specified
// debugging level.
func (o *Logger) Debugln(level int, v ...interface{}) {
	o.SetFlags(log.Lshortfile)
	if o.debug >= level {
		o.writeString(fmt.Sprintln(v...))
	}
	o.SetFlags(0)
}

// Debugf uses the formatting string and variables to print a message to the
// output streams if the current debugging level is greater than or equal to the
// specified debugging level.
func (o *Logger) Debugf(level int, format string, v ...interface{}) {
	o.SetFlags(log.Lshortfile)
	if o.debug >= level {
		o.writeString(fmt.Sprintf(format, v...))
	}
	o.SetFlags(0)
}
