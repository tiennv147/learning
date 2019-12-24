package mydb

import (
	"context"
	"database/sql"
	"errors"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

// DB is a logical database with multiple underlying physical databases
// forming a single master multiple replicas topology.
// Reads and writes are automatically directed to the correct physical db.
type DB struct {
	master        *sql.DB
	replicas      []*sql.DB
	mu            sync.RWMutex
	count         uint64
	masterCanRead bool
}

// Open concurrently opens each underlying physical db.
// driverName must be a semi-comma separated list of DSNs with the first
// one being used as the master and the rest as replicas.
func Open(driverName string, sources string) (*DB, error) {
	conns := strings.Split(sources, ";")
	if len(conns) == 0 {
		return nil, errors.New("empty servers list")
	}
	db := &DB{}
	for i, c := range conns {
		if len(c) == 0 { // trailing
			continue
		}
		s, err := sql.Open(driverName, c)
		if err != nil {
			return nil, err
		}

		err = s.Ping()
		if err != nil {
			return nil, err
		}

		if i == 0 { // first is the master
			db.master = s
		} else {
			db.replicas = append(db.replicas, s)
		}
	}
	if len(db.replicas) == 0 {
		db.replicas = append(db.replicas, db.master)
		db.masterCanRead = true
	}
	return db, nil
}

func NewDB(master *sql.DB, replicas ...*sql.DB) (*DB, error) {
	if master == nil {
		return nil, errors.New("no master for write")
	}
	db := &DB{
		master: master,
	}
	for _, replica := range replicas {
		if replica == nil {
			continue
		}
		db.replicas = append(db.replicas, replica)
	}

	if len(db.replicas) == 0 {
		db.replicas = append(db.replicas, db.master)
		db.masterCanRead = true
	}

	return db, nil
}

// Ping verifies if a connection to each physical database is still alive,
// establishing a connection if necessary.
func (db *DB) Ping() error {
	dbs := db.getAllDbs()
	return scatter(len(dbs), func(i int) error {
		return dbs[i].Ping()
	})
}

// PingContext verifies if a connection to each physical database is still
// alive, establishing a connection if necessary.
func (db *DB) PingContext(ctx context.Context) error {
	dbs := db.getAllDbs()
	return scatter(len(dbs), func(i int) error {
		return dbs[i].PingContext(ctx)
	})
}

// Query executes a query that returns rows, typically a SELECT.
// The args are for any placeholder parameters in the query.
// Query uses a replica as the physical db.
func (db *DB) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return db.readReplicaRoundRobin().Query(query, args...)
}

// QueryContext executes a query that returns rows, typically a SELECT.
// The args are for any placeholder parameters in the query.
// QueryContext uses a replica as the physical db.
func (db *DB) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	return db.readReplicaRoundRobin().QueryContext(ctx, query, args...)
}

// QueryRow executes a query that is expected to return at most one row.
// QueryRow always return a non-nil value.
// Errors are deferred until Row's Scan method is called.
// QueryRow uses a replica as the physical db.
func (db *DB) QueryRow(query string, args ...interface{}) *sql.Row {
	return db.readReplicaRoundRobin().QueryRow(query, args...)
}

// QueryRowContext executes a query that is expected to return at most one row.
// QueryRowContext always return a non-nil value.
// Errors are deferred until Row's Scan method is called.
// QueryRowContext uses a replica as the physical db.
func (db *DB) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	return db.readReplicaRoundRobin().QueryRowContext(ctx, query, args...)
}

// Begin starts a transaction on the master. The isolation level is dependent on the driver.
func (db *DB) Begin() (*sql.Tx, error) {
	return db.master.Begin()
}

// BeginTx starts a transaction with the provided context on the master.
//
// The provided TxOptions is optional and may be nil if defaults should be used.
// If a non-default isolation level is used that the driver doesn't support,
// an error will be returned.
func (db *DB) BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error) {
	return db.master.BeginTx(ctx, opts)
}

// Close closes all physical databases concurrently, releasing any open resources.
func (db *DB) Close() error {
	err := db.master.Close()
	if err != nil {
		return err
	}
	for i := range db.replicas {
		err = db.replicas[i].Close()
	}
	return err
}

// Exec executes a query without returning any rows.
// The args are for any placeholder parameters in the query.
// Exec uses the master as the underlying physical db.
func (db *DB) Exec(query string, args ...interface{}) (sql.Result, error) {
	return db.master.Exec(query, args...)
}

// ExecContext executes a query without returning any rows.
// The args are for any placeholder parameters in the query.
// Exec uses the master as the underlying physical db.
func (db *DB) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return db.master.ExecContext(ctx, query, args...)
}

// Prepare creates a prepared statement for later queries or executions
// on each physical database, concurrently.
func (db *DB) Prepare(query string) (Stmt, error) {
	mStmt, meErr := db.master.Prepare(query)
	rStmts := make([]*sql.Stmt, len(db.replicas))

	err := scatter(len(rStmts), func(i int) (err error) {
		rStmts[i], err = db.replicas[i].Prepare(query)
		return err
	})
	err = meErr

	if err != nil {
		return nil, err
	}

	return &stmt{
		db:           db,
		masterStmt:   mStmt,
		replicaStmts: rStmts,
	}, nil
}

// PrepareContext creates a prepared statement for later queries or executions
// on each physical database, concurrently.
//
// The provided context is used for the preparation of the statement, not for
// the execution of the statement.
func (db *DB) PrepareContext(ctx context.Context, query string) (*sql.Stmt, error) {
	return db.master.PrepareContext(ctx, query)
}

// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
// Expired connections may be closed lazily before reuse.
// If d <= 0, connections are reused forever.
func (db *DB) SetConnMaxLifetime(d time.Duration) {
	db.master.SetConnMaxLifetime(d)
	for i := range db.replicas {
		db.replicas[i].SetConnMaxLifetime(d)
	}
}

// SetMaxIdleConns sets the maximum number of connections in the idle
// connection pool for each underlying physical db.
// If MaxOpenConns is greater than 0 but less than the new MaxIdleConns then the
// new MaxIdleConns will be reduced to match the MaxOpenConns limit
// If n <= 0, no idle connections are retained.
func (db *DB) SetMaxIdleConns(n int) {
	db.master.SetMaxIdleConns(n)
	for i := range db.replicas {
		db.replicas[i].SetMaxIdleConns(n)
	}
}

// SetMaxOpenConns sets the maximum number of open connections
// to each physical database.
// If MaxIdleConns is greater than 0 and the new MaxOpenConns
// is less than MaxIdleConns, then MaxIdleConns will be reduced to match
// the new MaxOpenConns limit. If n <= 0, then there is no limit on the number
// of open connections. The default is 0 (unlimited).
func (db *DB) SetMaxOpenConns(n int) {
	db.master.SetMaxOpenConns(n)
	for i := range db.replicas {
		db.replicas[i].SetMaxOpenConns(n)
	}
}

// MasterCanRead adds the master physical database to the replicas list if read==true
// so that the master can perform WRITE queries AND READ queries .
func (db *DB) MasterCanRead(read bool) {
	db.mu.Lock()
	defer db.mu.Unlock()
	if read == true && db.masterCanRead == false {
		db.replicas = append(db.replicas, db.master)
		db.masterCanRead = read
	}
	if read == false && db.masterCanRead == true && len(db.replicas) > 1 {
		var replicas []*sql.DB
		for _, s := range db.replicas {
			if s != db.master {
				replicas = append(replicas, s)
			}
		}
		db.replicas = replicas
		db.masterCanRead = read
	}
}

// getAllDbs returns each underlying physical database,
// the first one is the master
func (db *DB) getAllDbs() []*sql.DB {
	var dbs []*sql.DB
	dbs = append(dbs, db.master)
	if len(db.replicas) > 0 {
		dbs = append(dbs, db.replicas...)
	}
	return dbs
}

func (db *DB) replica() int {
	if len(db.replicas) == 1 {
		return 0
	}
	return int(atomic.AddUint64(&db.count, 1) % uint64(len(db.replicas)))
}

func (db *DB) readReplicaRoundRobin() *sql.DB {
	db.mu.RLock()
	defer db.mu.RUnlock()
	return db.replicas[db.replica()]
}
