package strategy

import (
	"fmt"
)

// 上下文类cache
type cache struct {
	storage map[string]string
	// cache中保留对抽象策略接口的引用
	evictionAlgo
	cap    int
	maxCap int
}

// 初始化cache方法
func initCache(e evictionAlgo) *cache {
	return &cache{
		storage:      make(map[string]string),
		evictionAlgo: e,
		cap:          0,
		maxCap:       2,
	}
}

// cache的设置器方法，用于修改策略成员变量，以期调用不同策略。
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

// 上下文类cache的淘汰方法，内部调用不同的淘汰策略
func (c *cache) evict() {
	c.evictionAlgo.evict(c)
	c.cap--
}

// 抽象出的统一策略接口
type evictionAlgo interface {
	// 包含一个淘汰方法，通过传入上下文类（cache)指针，以期达到具体策略算法可访问到cached对象
	evict(c *cache)
}

// 各算法策略抽象成独立类
type fifo struct {
}

// 各独立类实现统一的策略接口
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
