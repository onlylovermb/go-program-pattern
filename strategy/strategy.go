package strategy

import (
	"fmt"
)

type cache struct {
	storage map[string]string
	evictionAlgo
	cap    int
	maxCap int
}

func initCache(e evictionAlgo) *cache {
	return &cache{
		storage:      make(map[string]string),
		evictionAlgo: e,
		cap:          0,
		maxCap:       2,
	}
}

func (c *cache) setEvictAlgo(e evictionAlgo) {
	c.evictionAlgo = e
}

func (c *cache) add(key, value string) {
	if c.cap == c.maxCap {
		c.evict()
	}
	c.cap++
	c.storage[key] = value
}

func (c *cache) get(key string) (value string, ok bool) {
	value, ok = c.storage[key]
	return
}

func (c *cache) del(key string) {
	if _, ok := c.get(key); ok {
		delete(c.storage, key)
	}
}

func (c *cache) evict() {
	c.evictionAlgo.evict(c)
	c.cap--
}

type evictionAlgo interface {
	evict(c *cache)
}

type fifo struct {
}

func (a *fifo) evict(c *cache) {
	fmt.Println("evicting by fifo strategy!")
}

type lru struct {
}

func (a *lru) evict(c *cache) {
	fmt.Println("evicting by lru strategy!")
}

type lfu struct {
}

func (a *lfu) evict(c *cache) {
	fmt.Println("evicting by lfu strategy!")
}
