package worldweather

import (
	"appengine"
)

//WorkerItem can be created and queued to the fetching queue
type WorkerItem struct {
	city       string
	ctx        appengine.Context
	outChannel chan []byte
}

//The manager is responsible for rate limiting the api requests
type APIRequestManager struct {
	queue chan *WorkerItem
}

//Create a new APIRequestsManager
func NewAPIRequestManager() *APIRequestManager {
	return &APIRequestManager{
		queue: make(chan *WorkerItem, 10),
	}
}

//Add item to fetching queue and receive a channel which will
//return json data after fetching it from the API
func (self *APIRequestManager) AddItem(ctx appengine.Context, cityName string) chan []byte {
	item := &WorkerItem{
		ctx:        ctx,
		city:       cityName,
		outChannel: make(chan []byte),
	}

	self.queue <- item

	return item.outChannel
}

//Run the worker which manages fetching the data from API
//should be run in a separate goroutine
func (self *APIRequestManager) RunWorker() {
	for {
		item := <-self.queue

		data, err := FetchWeatherData(item.ctx, item.city)
		if err != nil {
			close(item.outChannel)
			continue
		}

		item.outChannel <- data
	}
}
