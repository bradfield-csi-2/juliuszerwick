package main

type linkNode struct {
	item Item
	next *linkNode
	prev *linkNode
	up   *linkNode
	down *linkNode
}

type list struct {
	head *linkNode
	tail *linkNode
}

type skipListOC struct {
	lists []*list
}

func newSkipListOC() *skipListOC {
	head := &linkNode{}
	tail := &linkNode{}
	head.next = tail
	tail.prev = head
	newList := &list{
		head: head,
		tail: tail,
	}

	return &skipListOC{
		lists: []*list{newList},
	}
}

func (o *skipListOC) Get(key string) (string, bool) {
	// First, get the length of the skipListOC list.
	// The length - 1 will be the max level to start at.
	maxLevel := len(o.lists) - 1

	// Iterate through linked list to find key and value.
	currentList := o.lists[maxLevel]
	headNode := currentList.head
	node := headNode.next

	// If first node in list has a key greater than provided key,
	// move onto the next list in the OC.
	//if node.item.Key > key {
	//	// If we are already at the bottom list, then the key does not exit.
	//	// Break out of for loop and return "", false
	//	if i == 0 {
	//		break
	//	}
	//	continue
	//}

	for node.down != currentList.tail {
		for node != currentList.tail && key > node.item.Key {
			node = node.next
		}

		if node.item.Key != key && node.down != currentList.tail {
			//	if i == 0 {
			//		break
			//	}

			node = node.down
		}
	}

	for node != currentList.tail && key > node.item.Key {
		node = node.next
	}

	if node != currentList.tail && node.item.Key == key {
		return node.item.Value, true
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
