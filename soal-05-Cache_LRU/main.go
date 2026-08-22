package main

import "fmt"

type LRUCache struct {
	capacity int
	queue    []int
	cache    map[int]int
}

func NewLRLCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		cache:    make(map[int]int),
	}
}

func (c *LRUCache) Get(key int) (int, bool) {
	val, ok := c.cache[key]
	return val, ok
}

func (c *LRUCache) Put(key int, val int) {
	fmt.Println("len",len(c.queue))
	if len(c.queue) == c.capacity {
		staleKey := c.queue[0]
		fmt.Println("stale key",staleKey)
		delete(c.cache, staleKey)
		c.queue = c.queue[1:]
	}
	c.queue = append(c.queue, key)
	c.cache[key] = val

}

func main() {
	lruCache := NewLRLCache(2)
	lruCache.Put(1, 10)
	val, ok := lruCache.Get(1)
	fmt.Println(val, ok)
	lruCache.Put(11, 2)
	val2, ok2 := lruCache.Get(11)
	fmt.Println(val2, ok2)
	fmt.Println("b",lruCache.cache)
	lruCache.Put(1000, 159999)
	fmt.Println("a",lruCache.cache)
	val3, ok3 := lruCache.Get(1000)
	fmt.Println(val, ok)
	fmt.Println(val2, ok2)
	fmt.Println(val3, ok3)
	
}
