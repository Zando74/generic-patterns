package main

import "github.com/Zando74/generic-patterns/behavioral"

type Item struct {
	behavioral.Iterable
	Power int
}

type ListCollection struct { // an iterable collection of items in a list form
	*behavioral.Iterator[Item]
}

func (c *ListCollection) Push(item *Item) {

	wrappedItem := &behavioral.Iterator[Item]{Item: item}

	if c.Iterator == nil {
		c.Iterator = wrappedItem
	} else {
		c.Extend(wrappedItem)
	}
}

func (c *ListCollection) Pop() *Item {

	if c.Iterator == nil {
		return nil
	}

	penultimate := c.Penultimate()
	penultimate.SetNext(nil)

	return penultimate.Item
}

func NewListCollection(items ...*Item) *ListCollection {
	collection := &ListCollection{}

	for _, item := range items {
		collection.Push(item)
	}

	return collection
}

func MainIteratorExample() {

	item1 := &Item{Power: 1}
	item2 := &Item{Power: 2}
	item3 := &Item{Power: 3}

	ItemCollection := NewListCollection(item1, item2, item3)

	ItemCollection.Push(&Item{Power: 4})

	iterate := ItemCollection.InitIterator() // Initialize the iterator

	for it := iterate(); it != nil; it = iterate() { // Iterate over the collection
		println(it.Item.Power) // it contains the current item
	}

	ItemCollection.Pop()

	iterate = ItemCollection.InitIterator()
	for it := iterate(); it != nil; it = iterate() {
		println(it.Item.Power)
	}

}
