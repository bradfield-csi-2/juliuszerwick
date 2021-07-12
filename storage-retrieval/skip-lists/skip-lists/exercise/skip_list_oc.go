package main

type linkedNode struct {
	item Item
	next *linkedNode
	prev *linkedNode
	up   *linkedNode
	down *linkedNode
}

type list struct {
	head *linkedNode
	tail *linkedNode
}

type skipListOC struct {
	lists []*list
}

func newSkipListOC() *skipListOC {
	head := &linkedNode{}
	tail := &linkedNode{}
	head.next = tail
	tail.prev = head
	newList := list{
		head: head,
		tail: tail,
	}

	return &skipListOC{
		lists: []list{newList},
	}
}

func (o *skipListOC) Get(key string) (string, bool) {
	// First, get the length of the skipListOC list.
	// The length - 1 will be the max level to start at.
	maxLevel := len(o) - 1

	// Iterate through linked lists in skip list from highest level to lowest.
	// If we finish searching the lowest linked list without success, return "", false
	for i := maxLevel; i >= 0; i-- {
		// Iterate through linked list to find
	}
	return "", false
}

func (o *skipListOC) Put(key, value string) bool {
	return false
}

func (o *skipListOC) Delete(key string) bool {
	return false
}

// Below this line is if you have time.
func (o *skipListOC) RangeScan(startKey, endKey string) Iterator {
	return &skipListOCIterator{}
}

type skipListOCIterator struct {
}

func (iter *skipListOCIterator) Next() {
}

func (iter *skipListOCIterator) Valid() bool {
	return false
}

func (iter *skipListOCIterator) Key() string {
	return ""
}

func (iter *skipListOCIterator) Value() string {
	return ""
}
