package storage

import "errors"

type Transaction interface {
	Ref() interface{}
	Commit() error
	Rollback() error
	Begin() (Transaction, error)
	Query() interface{}
}

// Context is a generic definition for an SQL connection state
type Context interface {
	ConnStr() string
	Closed() bool
	Open() error
	Close()
	Tx() (Transaction, error)
	UnmanagedTx() (Transaction, error)
}

var ErrAlreadyExists = errors.New("Duplicated key detected")
var ErrNotFound = errors.New("Record not found")
var ErrTooManyRecordFound = errors.New("Requesting single record returns more than one")
