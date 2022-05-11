// Simple log replacement for go with Verbosity and Debug levels, and
// multi-stream output support.
//
// (C) by Marco Paganini <paganini AT paganini DOT net>

package logger

import (
	"context"
	"fmt"
	"io"
	"os"
)

// Logger represents a logger object
type Logger struct {
	verbose      int
	debug        int
	outputs      []io.Writer
	mirrorOutput io.Writer
}

// New Creates a new Logger instance.
func New(prefix string) *Logger {
	return &Logger{
		verbose:      0,
		debug:        0,
		outputs:      []io.Writer{os.Stderr},
		mirrorOutput: nil}
}

// key for Context use.
type key int

const (
	keyLogger = key(iota)
)

// SetVerboseLevel sets the verbosity level for this log instance.
func (o *Logger) SetVerboseLevel(n int) {
	o.verbose = n
}

// SetDebugLevel sets the debugging level for this log instance.
func (o *Logger) SetDebugLevel(n int) {
	o.debug = n
}

// SetOutputs sets all logging outputs to the outputs presented
// in the slice of io.Writers.
func (o *Logger) SetOutputs(outputs []io.Writer) {
	o.outputs = outputs
}

// SetMirrorOutput sets the mirror output stream to the writers.
func (o *Logger) SetMirrorOutput(output io.Writer) {
	o.mirrorOutput = output
}

// writeString sends the string to all defined outputs if the requested level
// is less than or equal to the configured verbose level. Note that if the
// mirrorOutput stream is non nil, the message is always written to it,
// independent of the error level.
func (o *Logger) writeString(level int, s string) {
	// Conditionally output to all streams
	if level <= o.verbose {
		for _, w := range o.outputs {
			io.WriteString(w, s)
		}
	}
	// Output to mirror output if non-nil
	if o.mirrorOutput != nil {
		io.WriteString(o.mirrorOutput, s)
	}
}

// Println prints the message to the output streams followed by a newline.
func (o *Logger) Println(v ...interface{}) {
	o.writeString(0, fmt.Sprintln(v...))
}

// Printf uses the formatting string and variables to print a message to the
// output streams.
func (o *Logger) Printf(format string, v ...interface{}) {
	o.writeString(0, fmt.Sprintf(format, v...))
}

// Fatalln prints the message to the output streams followed by a newline
// and calls os.Exit(1).
func (o *Logger) Fatalln(v ...interface{}) {
	o.writeString(0, fmt.Sprintln(v...))
	os.Exit(1)
}

// Fatal is a convenience alias for Fatalln.
func (o *Logger) Fatal(v ...interface{}) {
	o.Fatalln(v...)
}

// Fatalf prints a formatted message to the output streams and calls os.Exit(1).
func (o *Logger) Fatalf(format string, v ...interface{}) {
	o.writeString(0, fmt.Sprintf(format, v...))
	os.Exit(1)
}

// Verboseln prints the message to the output streams, followed by a newline,
// if the current verbose level is greater than or equal the specified
// verbosity level.
func (o *Logger) Verboseln(level int, v ...interface{}) {
	o.writeString(level, fmt.Sprintln(v...))
}

// Verbosef uses the formatting string and variables to print a message to the
// output streams if the current verbose level is greater than or equal to the
// specified verbose level.
func (o *Logger) Verbosef(level int, format string, v ...interface{}) {
	o.writeString(level, fmt.Sprintf(format, v...))
}

// Debugln prints the message to the output streams, followed by a newline,
// if the current debugging level is greater than or equal the specified
// debugging level.
func (o *Logger) Debugln(level int, v ...interface{}) {
	o.writeString(level, fmt.Sprintln(v...))
}

// Debugf uses the formatting string and variables to print a message to the
// output streams if the current debugging level is greater than or equal to the
// specified debugging level.
func (o *Logger) Debugf(level int, format string, v ...interface{}) {
	o.writeString(level, fmt.Sprintf(format, v...))
}

// WithLogger returns a new context with the logger object added to it.
func WithLogger(ctx context.Context, log *Logger) context.Context {
	return context.WithValue(ctx, keyLogger, log)
}

// Logf returns the logger function from the context. If no logger function has
// been set, create a new logger object and return it. Users should not rely on
// this behavior and set their own logger functions.
func Logf(ctx context.Context, log *Logger) *Logger {
	ret, ok := ctx.Value(keyLogger).(*Logger)
	if !ok {
		panic("internal error: No logger function set or wrong type.")
	}
	return ret
}
