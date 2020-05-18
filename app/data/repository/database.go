package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jinzhu/gorm/dialects/postgres" // postgres
)

const (
	host     = "balarama.db.elephantsql.com"
	port     = 5432
	user     = "mkduxrjc"
	password = "Aw34WfUwQ8Q-y-8NCXCedmSrnPfHmFkL"
	dbname   = "mkduxrjc"
)

var psqlInfo string = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
// var psqlInfo string = fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable", user, password, host, port, dbname)

// DataBaseHelper ...
// tentar PGX https://medium.com/avitotech/how-to-work-with-postgres-in-go-bad2dabd13e4
type DataBaseHelper struct {
	ctx context.Context
	tx  *sql.Tx
	db  *sql.DB
	err error
}

var err error

var dbHelper *DataBaseHelper = &DataBaseHelper{}

// TODO Fix that
var dataBaseHelper *DataBaseHelper = GetDb()

// GetDb ..
func GetDb() *DataBaseHelper {
	if dbHelper.db == nil {
		fmt.Println(psqlInfo)
		dbHelper.db, dbHelper.err = sql.Open("postgres", psqlInfo)
		if dbHelper.err != nil {
			fmt.Println(dbHelper.err)
		}
	}

	return dbHelper
}

// BeginTransaction ..
func (dbh *DataBaseHelper) BeginTransaction() {
	log.Println("Begin transaction")
	dbh.ctx = context.Background()
	dbh.tx, err = dbh.db.BeginTx(dbh.ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})

	if err != nil {
		log.Fatal("Error begin-transaction: ", err)
	}
}

// CommitTransaction ...
func (dbh *DataBaseHelper) CommitTransaction() {
	log.Println("Commiting transaction")
	dbh.tx.Commit()
}

// RollbackTransaction ..
func (dbh *DataBaseHelper) RollbackTransaction() {
	log.Println("Rollback transaction")
	dbh.tx.Rollback()
}

// CloseDb ..
func CloseDb() {
	log.Println("Close db")
	dbHelper.db.Close()
}
