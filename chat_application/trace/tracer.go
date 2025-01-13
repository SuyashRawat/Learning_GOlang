package trace

import (
	"fmt"
	"io"
)

// Tracer is the interface that describes an object capable of
// tracing events throughout code.
type Tracer interface {
	Trace(...interface{})
}
type tracer struct {
	out io.Writer
}

<<<<<<< HEAD
// New returns a new instance of tracer
func New(w io.Writer) Tracer {
	return &tracer{out: w}
}

// Trace writes the trace output to the specified io.Writer
=======
>>>>>>> 1f07827734644fd605fe5b3740ab0d4a63861888
func (t *tracer) Trace(a ...interface{}) {
	t.out.Write([]byte(fmt.Sprint(a...)))
	t.out.Write([]byte("\n"))
}
<<<<<<< HEAD
=======
func New(w io.Writer) Tracer {
	return &tracer{out: w}
}
>>>>>>> 1f07827734644fd605fe5b3740ab0d4a63861888
