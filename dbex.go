package main

import (
	"log"

	_ "github.com/go-sql-driver/mysql" // Any sql.DB works

	"learning/mydb"
)

func main() {
	// The first DSN is assumed to be the master and all
	// other to be replicas
	dsns := "tcp://user:password@master/dbname;"
	dsns += "tcp://user:password@replica01/dbname;"
	dsns += "tcp://user:password@replica02/dbname"

	db, err := mydb.Open("mysql", dsns)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("Some physical database is unreachable: %s", err)
	}

	// Read queries are directed to replicas with Query and QueryRow.
	// Always use Query or QueryRow for SELECTS
	// Load distribution is round-robin only for now.
	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM sometable").Scan(&count)
	if err != nil {
		log.Fatal(err)
	}

	// Write queries are directed to the master with Exec.
	// Always use Exec for INSERTS, UPDATES
	result, err := db.Exec("UPDATE sometable SET something = 1")
	if err != nil {
		log.Fatal(err)
	}

	log.Println(result)

	// Prepared statements are aggregates. If any of the underlying
	// physical databases fails to prepare the statement, the call will
	// return an error. On success, if Exec is called, then the
	// master is used, if Query or QueryRow are called, then a replica
	// is used.
	stmt, err := db.Prepare("SELECT * FROM sometable WHERE something = ?")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(stmt)

	// Transactions always use the master
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	// Do something transactional ...
	if err = tx.Commit(); err != nil {
		log.Fatal(err)
	}
}
