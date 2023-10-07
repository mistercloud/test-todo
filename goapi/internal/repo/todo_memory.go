package repo

import (
	"errors"
	"github.com/mistercloud/test-todo/goapi/internal/entity"
)

type ToDoMemory struct {
	Items []entity.Item
	maxId entity.ItemId
}

func NewToDoMemory() *ToDoMemory {
	return &ToDoMemory{
		Items: []entity.Item{},
		maxId: 0,
	}
}

func (c *ToDoMemory) Add(title string) (entity.Item, error) {
	if title == `` {
		return entity.Item{}, errors.New("empty title")
	}
	c.maxId++
	item := entity.Item{Id: c.maxId, Title: title}
	c.Items = append(c.Items, item)

	return item, nil
}

func (c *ToDoMemory) Remove(id entity.ItemId) error {
	i := -1
	for n := range c.Items {
		if c.Items[n].Id == id {
			i = n
		}
	}
	if i == -1 {
		return errors.New("no Id found")
	}
	c.Items = append(c.Items[:i], c.Items[i+1:]...)
	return nil
}

func (c *ToDoMemory) List() []entity.Item {
	ret := make([]entity.Item, len(c.Items))
	i := 0
	for _, item := range c.Items {
		ret[i] = item
		i++
	}
	return ret
}
