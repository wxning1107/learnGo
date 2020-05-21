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

	e.pre.next = e.next
	if e == cache.tail {
		cache.tail = e.pre
	} else {
		e.next.pre = e.pre
	}

	e.pre = nil

	e.next = cache.head
	cache.head.pre = e

	cache.head = e
}

func main() {
	cache := NewCache(2)
	cache.Put("1", "one")       // 放入元素one，此时one在队列头部
	fmt.Println(cache.Get("1")) // 此处输出“one”，此时one在队列头部
	cache.Put("2", "two")       // 放入元素two，此时two在队列头部
	fmt.Println(cache.Get("1")) // 此处输出“one”，此时one在队列头部
	cache.Put("3", "three")     // 放入元素three，总元素个数为3，因此最近最少使用的元素“2”会被删除
	fmt.Println(cache.Get("2")) // 此处输出nil
	fmt.Println(cache.Get("3")) // 此处输出“three”
	fmt.Println(cache.Get("3")) // 此处输出“three”
	fmt.Println(cache.Get("1")) // 此处输出“one”，此时最近最少使用的元素为“3”
	cache.Put("2", "two")       // 放入元素three，总元素个数为3，因此最近最少使用的元素“3”会被删除
	fmt.Println(cache.Get("3")) // 此处输出nil
	fmt.Println(cache.Get("1")) // 此处输出one
}
