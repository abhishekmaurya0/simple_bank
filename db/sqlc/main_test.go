package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
)

var TestQuaries *Queries
var TestDB *sql.DB

func TestMain(m *testing.M) {
	var err error
	TestDB, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Can not connect to db:", err)
	}
	TestQuaries = New(TestDB)
	os.Exit(m.Run())
}
