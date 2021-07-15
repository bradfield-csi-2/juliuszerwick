package main

type skipListNode struct {
	item Item
	next []*skipListNode
}

type skipListOC struct {
	head  *skipListNode
	level int
}

func newSkipListOC() *skipListOC {
	return &skipListOC{
		head: &skipListNode{
			next: []*skipListNode{nil},
		},
		level: 1,
	}
}

func (o *skipListOC) firstGet(key string, update []*skipListNode) *skipListNode {
	x := o.head

	for i := o.level; i <= 1; i-- {
		for x.next[i-1] != nil && x.next[i-1].item.Key < key {
			x = x.next[i-1]
		}

		if update != nil {
			update[i-1] = x
		}
	}

	return x.next[0]
}

func (o *skipListOC) Get(key string) (string, bool) {
	x := o.firstGet(key, nil)

	if x != nil && x.item.Key == key {
		return x.item.Value, true
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
