package db

import (
  "database/sql"
  "fmt"
  "log"
  _ "github.com/lib/pq"
  "github.com/go-gorp/gorp"
)

type DB struct {
  *sql.DB
}

const (
  DbUser = "macbook"
  DbPassword = ""
  DbName = "postgres"
)

var db *gorp.DbMap

func Init() {
  dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DbUser, DbPassword, DbName)

  var err error

  db, err = ConnectDB(dbinfo)

  if err != nil {
    log.Fatal(err)
  }
}

func ConnectDB(dataSourceName string) (*gorp.DbMap, error) {
	db, err := sql.Open("postgres", dataSourceName)

  if err != nil {
		return nil, err
	}

  if err = db.Ping(); err != nil {
		return nil, err
	}

  dbmap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}
	//dbmap.TraceOn("[gorp]", log.New(os.Stdout, "golang-gin:", log.Lmicroseconds)) //Trace database requests
	return dbmap, nil
}

func GetDB() *gorp.DbMap {
	return db
}
