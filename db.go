package main

import (
	"github.com/cznic/exp/dbm"
	"log"
	"os"
)

const (
	_DB_HASH    = "hash"
	_DB_REVERSE = "reverse"
	_DB_STATUS  = "status"
)

const (
	_DB_STATUS_COUNT   = "count"
	_DB_STATUS_CURRENT = "current"
)

var db *dbm.DB = nil
var dbReverse *dbm.Array
var dbHash *dbm.Array
var dbStatus *dbm.Array

func initDB() {
	const PATH = "dbm.db"
	var dbErr error
	if _, err := os.Stat(PATH); err == nil {
		db, dbErr = dbm.Open(PATH, new(dbm.Options))
	} else {
		db, dbErr = dbm.Create(PATH, &dbm.Options{})
	}
	if dbErr != nil {
		log.Panic(dbErr)
	}

	initArrays := func(pArray **dbm.Array, dbName string) {
		dbValue, dbErr := db.Array(dbName)
		checkErr(dbErr)
		*pArray = &dbValue
	}

	initArrays(&dbHash, _DB_HASH)
	initArrays(&dbReverse, _DB_REVERSE)
	initArrays(&dbStatus, _DB_STATUS)
}

func dbStatusGetCount() int64 {
	val, err := dbStatus.Inc(0, _DB_STATUS_COUNT)
	checkErr(err)
	return val
}

func dbStatusSetCount(count int64) {
	err := dbStatus.Set(count, _DB_STATUS_COUNT)
	checkErr(err)
}

func dbStatusGetCurrent() string {
	val, err := dbStatus.Get(_DB_STATUS_CURRENT)
	checkErr(err)
	if val == nil {
		return ""
	} else {
		return val.(string)
	}
}
