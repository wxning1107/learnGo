package main

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

func newCache(cap int) *Cache {
	return &Cache{capacity: cap, cache: make(map[string]*Entry)}
}

func (cache *Cache) Put(key string, value interface{}) interface{} {
	if entry, ok := cache.cache[key]; ok {
		cache.moveToHead(entry)
		return value
	}

	entry := &Entry{key: key, value: value, next: cache.head}

	return entry
}

func (cache *Cache) Get(key string) interface{} {
	if entry, ok := cache.cache[key]; ok {
		cache.moveToHead(entry)
		return entry.value
	}
	return nil
}

func (cache *Cache) moveToHead(entry *Entry) {

}

func (cache Cache) removeEntry(entry *Entry) {
	if entry == cache.head {
		//cache.head =
	}

}

func main() {

}
