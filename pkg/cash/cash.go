package cash

import (
	"container/list"
	"fmt"

	//"fmt"
	"image/jpeg"
	//"io"
	//"io/ioutil"
	//"fmt"
	"os"
	//"io"
	"image"
	//"image/jpeg"
	//"strconv"

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
}

func NewLru(capacity int) *LRU {
	return &LRU{
		capacity: capacity,
		items:    make(map[string]*list.Element),
		queue:    list.New(),
	}
}

func (c *LRU) Set(key string, value image.Image) bool {
	if element, exists := c.items[key]; exists == true {
		c.queue.MoveToFront(element)
		element.Value.(*Item).Value = value
		return true
	}

	if c.queue.Len() == c.capacity {
		c.MoveToDb(key)
		c.purge()
	}

	item := &Item{
		Key:   key,
		Value: value,
	}

	element := c.queue.PushFront(item)
	c.items[item.Key] = element
	return true
}

func (c *LRU) purge() {
	if element := c.queue.Back(); element != nil {
		item := c.queue.Remove(element).(*Item)
		delete(c.items, item.Key)
	}
}
func (c *LRU) Get(key string) image.Image {
	element, exists := c.items[key]
	if exists == false {
		return nil
	}
	c.queue.MoveToFront(element)
	return element.Value.(*Item).Value
}

func (c *LRU) MoveToDb(key string) error {

	element, exists := c.items[key]
	if exists == false {
		return nil
	}

	img := element.Value.(*Item).Value

	dbImg, err := os.Create("img.jpg")
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
