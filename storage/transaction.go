package storage

type Transaction interface {
	Ref() interface{}
	Commit() error
	Rollback() error
	Begin() (Transaction, error)
	Query() interface{}
}
