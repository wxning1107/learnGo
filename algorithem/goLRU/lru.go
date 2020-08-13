package main

import "fmt"

type Entry struct {
	key   string
	value interface{}
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
	if entry, ok := cache.cache[key]; ok {
		cache.moveToHead(entry)
		return value
	}

	e := &Entry{key: key, value: value, next: cache.head}
	if cache.head != nil {
		cache.head.pre = e
	}
	cache.head = e
	cache.cache[key] = e

	if cache.tail == nil {
		cache.tail = e
	}

	if cache.capacity < len(cache.cache) {
		delete(cache.cache, cache.tail.key)
		cache.tail.pre.next = nil
		cache.tail = cache.tail.pre

		return value
	}

	return value
}

func (cache *Cache) Get(key string) interface{} {
	if entry, ok := cache.cache[key]; ok {
		cache.moveToHead(entry)
		return entry.value
	}

	return nil
}

func (cache *Cache) moveToHead(entry *Entry) {
	if entry == cache.head {
		return
	}

	entry.pre.next = entry.next

	if entry == cache.tail {
		cache.tail = entry.pre
	} else {
		entry.next.pre = entry.pre
	}

	cache.head.pre = entry

	entry.next = cache.head
	entry.pre = nil

	cache.head = entry
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
