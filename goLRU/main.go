package main

import (
	"fmt"
	"sync"
)

var lock sync.RWMutex

type Entry struct {
	Key   string
	Value interface{}
	pre   *Entry
	next  *Entry
}

type Cache struct {
	cache    map[string]*Entry
	capacity int
	head     *Entry
	tail     *Entry
}

func NewCache(cap int) *Cache {
	return &Cache{cache: make(map[string]*Entry), capacity: cap}
}

func (cache *Cache) Put(key string, value interface{}) interface{} {
	lock.Lock()
	defer lock.Unlock()

	if val, ok := cache.cache[key]; ok {
		cache.moveToHead(val)
		return nil
	}

	e := &Entry{Key: key, Value: value, next: cache.head}
	if cache.head != nil {
		cache.head.pre = e
	}
	cache.head = e
	if cache.tail == nil {
		cache.tail = e
	}
	cache.cache[key] = e

	if len(cache.cache) <= cache.capacity {
		return nil
	}

	removedEntry := cache.tail
	cache.tail = cache.tail.pre
	removedEntry.pre = nil
	cache.tail.next = nil
	delete(cache.cache, removedEntry.Key)

	return removedEntry.Value
}

func (cache *Cache) Get(key string) interface{} {
	lock.Lock()
	defer lock.Unlock()

	if val, ok := cache.cache[key]; ok {
		cache.moveToHead(val)
		return val.Value
	}

	return nil
}

func (cache *Cache) moveToHead(e *Entry) {
	if e == cache.head {
		return
	}

	if e == cache.tail {
		cache.tail = e.next
		e.next.pre = nil
	} else {
		e.pre.next = e.next
		e.next.pre = e.pre
	}

	e.pre = cache.head
	e.next = nil
	cache.head = e
}

func main() {
	cache := NewCache(2)
	cache.Put("1", "one")
	cache.Put("2", "two")
	cache.Put("3", "three")
	cache.Put("2", "two")
	fmt.Println(cache.Get("1"))
}
