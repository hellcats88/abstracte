package logging

// K is used to define a log message ctx parameter, using Name/Value form
type K struct {
	N string
	V interface{}
}

// Context contains all context information for a log message plus extra arguments
type Context interface {
	CorrID() string
	Extras() []K
	Clone() Context
	CloneNoExtra() Context
	AddExtra(extra ...K)
}
