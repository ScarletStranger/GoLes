package main

import (
	"bufio"
	"fmt"
	"os"
)

type Cache interface {
	Get(k string) (string, bool)
	Set(k, v string)
}

var _ Cache = (*cacheImpl)(nil)

// Доработает конструктор и методы кеша, так чтобы они соответствовали интерфейсу Cache
func newCacheImpl() *cacheImpl {
	return &cacheImpl{}
}

type cacheImpl struct {
	Raw map[string]string
}

func (c *cacheImpl) Get(k string) (string, bool) {
	value, ok := c.Raw[k]
	if ok {
	  return fmt.Sprintf("key: %s, value: %s", k, value), true
	}
	return "", false
  }

func (c *cacheImpl) Set(k, v string) {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	for key, value := range c.Raw {
		c.Raw[value] = text
		key += key
	}
}

func newDbImpl(cache Cache) *dbImpl {
	return &dbImpl{cache: cache, dbs: map[string]string{"hello": "world", "test": "test"}}
}

type dbImpl struct {
	cache Cache
	dbs   map[string]string
}

func (d *dbImpl) Get(k string) (string, bool) {
	v, ok := d.cache.Get(k)
	if ok {
		return fmt.Sprintf("answer from cache: key: %s, val: %s", k, v), ok
	}

	v, ok = d.dbs[k]
	return fmt.Sprintf("answer from dbs: key: %s, val: %s", k, v), ok
}

func main() {
	c := newCacheImpl()
	db := newDbImpl(c)
	fmt.Println(db.Get("test"))
	fmt.Println(db.Get("hello"))
}
