package cash

import (
	"container/list"
	"fmt"
	"image"
	"image/jpeg"
	"os"
	"sync"

	"github.com/mjk712/kartinochki/pkg/config"
	"github.com/mjk712/kartinochki/pkg/lib/e"
)

type Item struct {
	Key   string
	Value image.Image
}

type LRU struct {
	capacity int
	items    map[string]*list.Element
	queue    *list.List
	lock     sync.RWMutex
}

func NewLru(capacity int) *LRU {
	return &LRU{
		capacity: capacity,
		items:    make(map[string]*list.Element),
		queue:    list.New(),
	}
}

func (c *LRU) Set(key string, value image.Image) bool {

	c.lock.Lock()
	if element, exists := c.items[key]; exists == true {
		c.queue.MoveToFront(element)
		element.Value.(*Item).Value = value
		return true
	}

	if c.queue.Len() == c.capacity {
		fmt.Println("mnogo cache")
		c.MoveToDb()
		c.lock.Lock()
		c.purge()
		c.lock.Unlock()
	}

	item := &Item{
		Key:   key,
		Value: value,
	}

	element := c.queue.PushFront(item)
	c.items[item.Key] = element
	c.lock.Unlock()
	return true
}

func (c *LRU) purge() {
	if element := c.queue.Back(); element != nil {
		item := c.queue.Remove(element).(*Item)
		delete(c.items, item.Key)
	}
}
func (c *LRU) Get(key string) (image.Image, bool) {
	c.lock.Lock()
	element, exists := c.items[key]
	if exists == false {
		return nil, false
	}
	c.queue.MoveToFront(element)
	c.lock.Unlock()
	return element.Value.(*Item).Value, true
}

func (c *LRU) MoveToDb() error {

	c.lock.Lock()

	if element := c.queue.Back(); element != nil {

		img := element.Value.(*Item).Value

		dbImg, err := os.Create(config.DbPath() + element.Value.(*Item).Key)
		if err != nil {
			er := e.Wrap("error while move to db", err)
			return er
		}
		defer dbImg.Close()

		if err = jpeg.Encode(dbImg, img, nil); err != nil {
			e.Wrap("error encode", err)
		}
		fmt.Println("FUCK WORK")
		return nil

	}
	c.lock.Unlock()

	return nil
}
