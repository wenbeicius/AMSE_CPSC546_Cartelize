package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/go-kit/kit/log"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	service "github.com/tryu-fullerton-edu/AMSE_CPSC546_Cartelize/accounts/cmd/service"
)

func main() {

	logger := createLogger()
	startLogger := log.With(logger, "tag", "start")
	startLogger.Log("msg", "created logger")

	db, err := createDatabase()

	if err != nil {
		startLogger.Log("msg", "failed to connect to database", "err", err)
		os.Exit(1)
	}

	defer db.Close()
	startLogger.Log("msg", "Connected to database")

	service.Run()
}

func createDatabase() (*sqlx.DB, error) {
	user := "postgres"
	password := "admin"
	db, err := sqx.Connect("postgres", fmt.Sprintf("user=%s password=%s dbname=cartelizedb sslmode=disable",
		user, password))

	if err != nil {
		return nil, err
	}

	ddl, err := ioutil.ReadFile("CreateTables.sql")
	if err != nil {
		return nil, err
	}

	_, err := db.Exec(string(ddl))

	return db, err
}
