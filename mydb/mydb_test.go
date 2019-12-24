package mydb

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func TestOpen1(t *testing.T) {
	db, err := Open("sqlite3", ":memory:;:memory:;:memory:")
	assert.Nil(t, err)
	defer db.Close()
	err = db.Ping()
	assert.Nil(t, err)
	assert.Equal(t, 2, len(db.replicas))
}

func TestOpen2(t *testing.T) {
	db, err := Open("sqlite3", ":memory:")
	assert.Nil(t, err)
	defer db.Close()
	err = db.Ping()
	assert.Nil(t, err)
	assert.Equal(t, 1, len(db.replicas))
	assert.True(t, db.masterCanRead)
}

func TestClose(t *testing.T) {
	db, err := Open("sqlite3", ":memory:;:memory:;:memory:")
	assert.Nil(t, err)

	err = db.Close()
	assert.Nil(t, err)

	err = db.Ping()
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "sql: database is closed")
}

func TestNewDB(t *testing.T) {
	masterDB, _, err := sqlmock.New()
	assert.Nil(t, err)
	defer masterDB.Close()

	db, err := NewDB(masterDB, nil, nil)
	assert.Nil(t, err)

	result := db.readReplicaRoundRobin()
	assert.NotNil(t, result)
}

func TestReadReplicaRoundRobin1(t *testing.T) {
	masterDB, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer masterDB.Close()
	readDB1, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	readDB2, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer readDB2.Close()

	db, err := NewDB(masterDB, readDB1, readDB2)
	assert.Nil(t, err)

	result := db.readReplicaRoundRobin()

	assert.NotNil(t, result)
	assert.Equal(t, readDB2, result)
}

func TestReadReplicaRoundRobin2(t *testing.T) {
	masterDB, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer masterDB.Close()
	readDB1, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer readDB1.Close()
	readDB2, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer readDB2.Close()

	db, err := NewDB(masterDB, readDB1, readDB2)
	assert.Nil(t, err)

	result1 := db.readReplicaRoundRobin()
	result2 := db.readReplicaRoundRobin()
	result3 := db.readReplicaRoundRobin()
	result4 := db.readReplicaRoundRobin()

	assert.NotNil(t, result1)
	assert.NotNil(t, result2)
	assert.NotNil(t, result3)
	assert.NotNil(t, result4)
	assert.Equal(t, readDB2, result1)
	assert.Equal(t, readDB1, result2)
	assert.Equal(t, readDB2, result3)
	assert.Equal(t, readDB1, result4)
}

func TestPrepare1(t *testing.T) {
	masterDB, masterMock, err := sqlmock.New()
	assert.Nil(t, err)

	replicaDB1, replicaMock1, err := sqlmock.New()
	assert.Nil(t, err)

	replicaDB2, replicaMock2, err := sqlmock.New()
	assert.Nil(t, err)

	masterMock.ExpectPrepare("SELECT").WillBeClosed()
	replicaMock1.ExpectPrepare("SELECT").WillBeClosed()
	replicaMock2.ExpectPrepare("SELECT ").WillBeClosed()

	replicaMock1.
		ExpectQuery("SELECT").
		WillReturnRows(sqlmock.NewRows([]string{"YEAH", "YO", "YEP"}))

	db, err := NewDB(masterDB, replicaDB1, replicaDB2)
	assert.Nil(t, err)

	result, err := db.Prepare("SELECT")
	assert.Nil(t, err)

	realStmt := result.(*stmt)

	assert.NotNil(t, realStmt.masterStmt)
	assert.Equal(t, 2, len(realStmt.replicaStmts))
	assert.Equal(t, 3, len(realStmt.allStmt()))
	_, err = realStmt.Query("item")
	assert.NotNil(t, err)
}
