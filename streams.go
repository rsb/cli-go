package cli

import (
	"fmt"
	"io"
	"os"
)

// Streams represents the 3 modes by which data travels via the cli.
// In: 	the standard input os.Stdin of the app.
// Out:	the standard output os.Stdout of the app.
// Err: the standard error os.Stderr of the app.
//
// These can all be controlled by the user, but left on touched the defaults
// are listed as above
type Streams struct {
	in  io.Reader
	out io.Writer
	err io.Writer
}

// NewStreams constructor used to create in/out and err streams
func NewStreams(in io.Reader, out, err io.Writer) Streams {
	return Streams{
		in:  in,
		out: out,
		err: err,
	}
}

func (ds *Streams) SetIn(in io.Reader) {
	ds.in = in
}

func (ds *Streams) In() io.Reader {
	if ds.in == nil {
		ds.in = os.Stdin
	}
	return ds.in
}

func (ds *Streams) SetOut(out io.Writer) {
	ds.out = out
}

func (ds *Streams) Out() io.Writer {
	if ds.out == nil {
		ds.out = os.Stdout
	}
	return ds.out
}

func (ds *Streams) SetError(err io.Writer) {
	ds.err = err
}

func (ds *Streams) Error() io.Writer {
	if ds.err == nil {
		ds.err = os.Stderr
	}

	return ds.err
}

// Print is a convenience method to Print to the Streams output
func (ds *Streams) Print(i ...interface{}) {
	_, _ = fmt.Fprint(ds.Out(), i...)
}

// Println is a convenience method to Println
func (ds *Streams) Println(i ...interface{}) {
	ds.Print(fmt.Sprintln(i...))
}

// Printf is a convenience to Printf
func (ds *Streams) Printf(format string, i ...interface{}) {
	ds.Print(fmt.Sprintf(format, i...))
}

// PrintErr is a convenience method to Fprint to the defined Err output
func (ds *Streams) PrintErr(i ...interface{}) {
	_, _ = fmt.Fprint(ds.Error(), i...)
}

// PrintErrln is a convenience method to Println to the defined Err output
func (ds *Streams) PrintErrln(i ...interface{}) {
	ds.PrintErr(fmt.Sprintln(i...))
}

// PrintErrf is a convenience method to Print
func (ds *Streams) PrintErrf(i ...interface{}) {
	ds.Print(fmt.Sprintln(i...))
}
