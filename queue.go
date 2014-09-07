package worldweather

import (
	"errors"
	"sync"
)

//something something template support
//just cast this to any kind of data you want to put in the Stack
type Item interface{}

//A minimal Stack implemented using map
type Stack struct {
	syncLock *sync.Mutex
	elements map[int]Item
	current  int
}

func NewStack() *Stack {
	return &Stack{
		syncLock: &sync.Mutex{},
		elements: make(map[int]Item),
		current:  -1,
	}
}

//Add Item to Stack and return the current index
//of the item
func (self *Stack) Push(item Item) int {
	self.syncLock.Lock()

	self.current = self.current + 1
	self.elements[self.current] = item

	self.syncLock.Unlock()

	return self.current
}

//Pop an item from the stack
func (self *Stack) Pop() (item Item, err error) {
	self.syncLock.Lock()
	if self.current < 0 {
		err = errors.New("Stack empty")
	}

	item = self.elements[self.current]
	delete(self.elements, self.current)

	self.current = self.current - 1

	self.syncLock.Unlock()

	return
}
