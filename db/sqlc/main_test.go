package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
	"github.com/wasinaphatlilawatthananan/go-postgres/util"
)

var testQueries *Queries
var testDB *sql.DB


func TestMain(m *testing.M)  {
	var err error
	config,err := util.LoadConfig("../..")
	if err !=nil{
		log.Fatal("cannot connect to db:",err)
	}
	testDB,err = sql.Open(config.DBDriver,config.DBSource)

	if err !=nil{
		log.Fatal("cannot connect to db:",err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
	
}