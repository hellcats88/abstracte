package logging

import "github.com/google/uuid"

// K is used to define a log message ctx parameter, using Name/Value form
type K struct {
	N string
	V interface{}
}

// Context contains all context information for a log message plus extra arguments
type Context struct {
	corrID string
	extra  []K
}

// NewContext creates new logger context from a correlation ID
func NewContext(corrID string) Context {
	return Context{
		corrID: corrID,
	}
}

// NewContextUUID creates new logger context generating
// a new correlation ID from UUID package
func NewContextUUID() Context {
	return Context{
		corrID: uuid.New().String(),
	}
}

// CorrID returns the context correlation ID
func (ctx Context) CorrID() string {
	return ctx.corrID
}

// CloneEmpty create a new independent context with the same original correlation ID without extra parameters
func (ctx Context) CloneEmpty() Context {
	return NewContext(ctx.CorrID())
}

// Clone create a new independent context with the same original correlation ID and extra parameters
func (ctx Context) Clone() Context {
	return Context{
		corrID: ctx.corrID,
		extra:  ctx.extra,
	}
}

// AddExtra appends to the internal memory the input list of extra parameters. Duplicates are admitted
func (ctx *Context) AddExtra(extras ...K) {
	ctx.extra = append(ctx.extra, extras...)
}

// GetExtras returns the list of stored extras in this context
func (ctx Context) GetExtras() []K {
	return ctx.extra
}
