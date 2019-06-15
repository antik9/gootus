package list

import (
	"sync"
)

type List struct {
	sync.Mutex
	first  *Item
	last   *Item
	length uint
}

type Item struct {
	value     interface{}
	previous  *Item
	next      *Item
	container *List
}

func NewList() List {
	return List{first: nil, last: nil, length: 0}
}

func (self *List) Len() uint {
	return self.length
}

func (self *List) First() *Item {
	return self.first
}

func (self *List) Last() *Item {
	return self.last
}

func (self *List) PushFront(value interface{}) {
	self.Lock()
	defer self.Unlock()
	if self.first == nil {
		self.initFirst(value)
	} else {
		self.first.previous = &Item{value: value, previous: nil, next: self.First(), container: self}
		self.first = self.first.previous
	}
	self.length++
}

func (self *List) PushBack(value interface{}) {
	self.Lock()
	defer self.Unlock()
	if self.last == nil {
		self.initFirst(value)
	} else {
		self.last.next = &Item{value: value, previous: self.Last(), next: nil, container: self}
		self.last = self.last.next
	}
	self.length++
}

func (self *List) initFirst(value interface{}) {
	newItem := Item{value: value, previous: nil, next: nil, container: self}
	self.first = &newItem
	self.last = &newItem
}

func (self *Item) Value() interface{} {
	return self.value
}

func (self *Item) Next() *Item {
	return self.next
}

func (self *Item) Prev() *Item {
	return self.previous
}

func (self *Item) Remove() {
	if self.container != nil {
		self.container.remove(self)
	} else {
		self.remove()
	}
}

func (self *Item) remove() {
	next := self.next
	previous := self.previous

	if next != nil {
		self.next.previous = previous
	}
	if previous != nil {
		self.previous.next = next
	}

	self.next = nil
	self.previous = nil
	self.value = nil
	self.container = nil
}

func (self *List) remove(item *Item) {
	self.Lock()
	defer self.Unlock()

	if item == self.First() {
		self.first = self.first.next
	}
	if item == self.Last() {
		self.last = self.last.previous
	}

	item.remove()
	self.length--
}
