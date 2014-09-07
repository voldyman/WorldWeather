package worldweather

import (
	"errors"
	"sync"
)

type Item interface{}

type Queue struct {
	syncLock *sync.Mutex
	elements map[int]Item
	current  int
}

func NewQueue() *Queue {
	return &Queue{
		syncLock: &sync.Mutex{},
		elements: make(map[int]Item),
		current:  -1,
	}
}

func (self *Queue) Push(item Item) int {
	self.syncLock.Lock()

	self.current = self.current + 1
	self.elements[self.current] = item

	self.syncLock.Unlock()

	return self.current
}

func (self *Queue) Pop() (item Item, err error) {
	self.syncLock.Lock()
	if self.current < 0 {
		err = errors.New("Queue empty")
	}

	item = self.elements[self.current]
	delete(self.elements, self.current)

	self.current = self.current - 1

	self.syncLock.Unlock()

	return
}
