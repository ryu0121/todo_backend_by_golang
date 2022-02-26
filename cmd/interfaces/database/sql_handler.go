package database

type SqlHandler interface {
	Query(string, ...interface{}) (Rows, error)
	QueryRow(string, ...interface{}) Row
	Exec(string, ...interface{}) (Result, error)
}

type Rows interface {
	Scan(...interface{}) error
	Next() bool
	Close() error
}

type Row interface {
	Scan(...interface{}) error
}

type Result interface {
	LastInsertId() (int64, error)
}
