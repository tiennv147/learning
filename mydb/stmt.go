package mydb

import (
	"database/sql"
)

// Stmt is an aggregate prepared statement.
// It holds a prepared statement for each underlying physical db.
type Stmt interface {
	Close() error
	Exec(...interface{}) (sql.Result, error)
	Query(...interface{}) (*sql.Rows, error)
	QueryRow(...interface{}) *sql.Row
}

type stmt struct {
	db           *DB
	masterStmt   *sql.Stmt
	replicaStmts []*sql.Stmt
}

func (s *stmt) allStmt() []*sql.Stmt {
	var stmts []*sql.Stmt
	stmts = append(stmts, s.masterStmt)
	if len(s.replicaStmts) > 0 {
		stmts = append(stmts, s.replicaStmts...)
	}
	return stmts
}

// Close closes the statement by concurrently closing all underlying
// statements concurrently, returning the first non nil error.
func (s *stmt) Close() error {
	stmts := s.allStmt()
	return scatter(len(stmts), func(i int) error {
		return stmts[i].Close()
	})
}

// Exec executes a prepared statement with the given arguments
// and returns a Result summarizing the effect of the statement.
// Exec uses the master as the underlying physical db.
func (s *stmt) Exec(args ...interface{}) (sql.Result, error) {
	return s.masterStmt.Exec(args...)
}

// Query executes a prepared query statement with the given
// arguments and returns the query results as a *sql.Rows.
// Query uses a slave as the underlying physical db.
func (s *stmt) Query(args ...interface{}) (*sql.Rows, error) {
	return s.replicaStmts[s.db.replica()].Query(args...)
}

// QueryRow executes a prepared query statement with the given arguments.
// If an error occurs during the execution of the statement, that error
// will be returned by a call to Scan on the returned *Row, which is always non-nil.
// If the query selects no rows, the *Row's Scan will return ErrNoRows.
// Otherwise, the *sql.Row's Scan scans the first selected row and discards the rest.
// QueryRow uses a slave as the underlying physical db.
func (s *stmt) QueryRow(args ...interface{}) *sql.Row {
	return s.replicaStmts[s.db.replica()].QueryRow(args...)
}
