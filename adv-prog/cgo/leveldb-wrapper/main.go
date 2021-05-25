package main

/*
#cgo CFLAGS: -I/usr/include

#include <stdlib.h>
#include "leveldb/db.h"

leveldb::DB* db;
leveldb::Options options;
options.create_if_missing = true;
options.error_if_exists = true;

int OpenDB(leveldb::DB* db) {
	leveldb::Status status = leveldb::DB::Open(options, "/tmp/testdb", &db);

	if (!status.ok()) {
		return 0;
	}

	return 1;
}

int CloseDB(leveldb::DB* db) {
	delete db;
	return 1;
}

int Put(leveldb::DB* db, std::string key, std::string value) {
	leveldb::Status s = db->Put(leveldb::WriteOptions(), key, value);

	if (!s.ok()) {
		return 0;
	}

	return 1;
}

std::string Get(leveldb::DB* db, std::string key) {
	std::string value;
	leveldb::Status s = db->Get(leveldb::WriteOptions(), key, &value);

	return value;
}
*/
import "C"

import "fmt"

/*
Steps:
	1) Succeed in opening a db connection.
	2) Then support closing the same connection.
	3) Add support to put values into the db and test insertions.
	4) Add support for getting values from the db and test with Put
		 to insert a value and then Get to get the same value.

NOTE: Import "leveldb/c.h" and use structs/functions defined in there and c.cc
			https://github.com/google/leveldb/blob/master/db/c.cc
*/

type Database struct {
	LevelDB *C.leveldb_t
}

func main() {
	// Open a db connection.
	db := C.leveldb_t
	dbName := "cgo-testdb"
	i := Open(dbName)

	if i == 0 {
		log.Fatalln("failed to open db connection")
	}

	// Make a Put into the db.

	// Make a Get from the db.

	// Close the db connection.
}

// Open() opens a connection to the database.
func Open(name string) (*Database, error) {
	// Pass these options into Open() instead.
	opts := C.leveldb_options_t
	dbName := C.CString(name)
	defer C.free(unsafe.Pointer(dbName))

	ldb := C.leveldb_open(opts, dbName)

	return &Database{ldb}, nil
}

// Close() closes a connection to the database.
func Close(db *Database) error {
	C.leveldb_close(db)
}

// Put() inserts a value into the database.
//func Put(db *Database, key string, value string) error {
//	//leveldb_put(db, )
//}

// Get() retrieves a value from the database.
//func Get(key string) (value string) {
//}
