package worldweather

import (
	"appengine"
)

type WorkerItem struct {
	city       string
	ctx        appengine.Context
	outChannel chan []byte
}

type APIRequestManager struct {
	queue chan *WorkerItem
}

func NewAPIRequestManager() *APIRequestManager {
	return &APIRequestManager{
		queue: make(chan *WorkerItem, 10),
	}
}

func (self *APIRequestManager) AddItem(ctx appengine.Context, cityName string) chan []byte {
	item := &WorkerItem{
		ctx:        ctx,
		city:       cityName,
		outChannel: make(chan []byte),
	}

	self.queue <- item

	return item.outChannel
}

func (self *APIRequestManager) RunWorker() {
	for {
		item := <-self.queue

		data, err := fetchWeatherData(item.ctx, item.city)
		if err != nil {
			close(item.outChannel)
			continue
		}

		item.outChannel <- data
	}
}
