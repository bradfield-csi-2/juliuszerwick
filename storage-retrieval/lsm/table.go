package table

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Item struct {
	Key, Value string
}

// Given a sorted list of key/value pairs, write them out according to the format you designed.
func Build(path string, sortedItems []Item) error {
	t := &Table{
		Data: sortedItems,
	}

	// Build Table Index by setting map keys to Item.Key and map values (offsets) to Item's position in slice.
	for i, v := range sortedItems {
		t.Index[v.Key] = i
	}

	// Encode values into buffer
	// We use JSON just to create an MVP and can use the encoding/gob package to encode/decode binary data later.
	jsonData, err := json.Marshal(t)
	if err != nil {
		log.Printf("JSON Marshal error: %v\n", err)
		return err
	}

	// Write to file
	err = ioutil.WriteFile(path, jsonData, 0644)
	if err != nil {
		log.Printf("WriteFile error: %v\n", err)
		return err
	}

	return nil
}

// A Table provides efficient access into sorted key/value data that's organized according
// to the format you designed.
//
// Although a Table shouldn't keep all the key/value data in memory, it should contain
// some metadata to help with efficient access (e.g. size, index, optional Bloom filter).
type Table struct {
	// TODO
	//idx  *SSIndex
	Index map[string]int
	Data  []Item
}

//type SSIndex struct {
//	key    string
//	offset int
//}

// Prepares a Table for efficient access. This will likely involve reading some metadata
// in order to populate the fields of the Table struct.
//
// Note to self: Perhaps when we "load a table" we're really only loading the SSIndex?
// That leaves the actual reading from the file to the Get function which will find
// the offset in the file to read?
// That would also mean each offset has to include the len(SSIndex) which would be the
// "metadata" at the start of the file.
func LoadTable(path string) (*Table, error) {
	f, err := ioutil.ReadFile(path)
	if err != nil {
		log.Printf("WriteFile error: %v\n", err)
		return nil, err
	}

	t := &Table{}
	err = json.Unmarshal([]byte(f), &t)
	if err != nil {
		log.Printf("JSON decode error: %v\n", err)
		return nil, err
	}

	return t, nil
}

func (t *Table) Get(key string) (string, bool, error) {
	// Use key to lookup offset in SSIndex.
	offset := t.Index[key]

	// Access Item value in Data slice using offset.
	item := t.Data[offset]

	// Double check that provided key matches found Item.Key
	if item.Key != key {
		return "", false, nil
	}

	return item.Value, true, nil
}

func (t *Table) RangeScan(startKey, endKey string) (Iterator, error) {
	return nil, nil
}

type Iterator interface {
	// Advances to the next item in the range. Assumes Valid() == true.
	Next()

	// Indicates whether the iterator is currently pointing to a valid item.
	Valid() bool

	// Returns the Item the iterator is currently pointing to. Assumes Valid() == true.
	Item() Item
}
