package main

/*
#include <cassert>
#include "leveldb/db.h"

leveldb::DB* db;
leveldb::Options options;
options.create_if_missing = true;
options.error_if_exists = true;

int OpenDB(leveldb::DB* db) {
	leveldb::Status status = leveldb::DB::Open(options, "/tmp/testdb", &db);

	if (!status.ok()) {
		cerr << s.ToString() << endl;
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
		cerr << s.ToString() << endl;
		return 0;
	}

	return 1;
}

std::string Get(leveldb::DB* db, std::string key) {
	std::string value;
	leveldb::Status s = db->Get(leveldb::WriteOptions(), key, &value);

	if (!s.ok()) {
		cerr << s.ToString() << endl;
	}

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
*/

func main() {
	// Open a db connection.

	// Make a Put into the db.

	// Make a Get from the db.

	// Close the db connection.
}

// Open() opens a connection to the database.
func Open() {
}

// Close() closes a connection to the database.
func Close() {
}

// Put() inserts a value into the database.
func Put(key string, value string) error {
}

// Get() retrieves a value from the database.
func Get(key string) (value string) {
}
