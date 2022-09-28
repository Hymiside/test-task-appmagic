package cache

import "errors"

type Cache struct {
	items map[string]Item
}

type Item struct {
	value interface{}
}

func NewCache() *Cache {
	items := make(map[string]Item)
	return &Cache{items: items}
}

func (c *Cache) Set(key string, value interface{}) {
	c.items[key] = Item{value: value}
}

func (c *Cache) Get(key string) (interface{}, error) {
	data := c.items[key].value
	if data == nil {
		return nil, errors.New("error no data")
	}
	return data, nil
}
