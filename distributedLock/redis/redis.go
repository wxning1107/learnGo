package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/rs/xid"
	"sync"
	"time"
)

//func incr() {
//	client := redis.NewClient(&redis.Options{Addr: "localhost:6379", Password: "cc", DB: 0})
//
//	var lockKey = "counter_lock"
//	var counterKey = "counter"
//
//	resp := client.SetNX(lockKey, 1, time.Second*5)
//	lockSuccess, err := resp.Result()
//	if err != nil || !lockSuccess {
//		fmt.Println(err, "lock result: ", lockSuccess)
//		return
//	}
//
//	getResp := client.Get(counterKey)
//	cntValue, err := getResp.Int64()
//	if err == nil {
//		cntValue++
//
//		resp := client.Set(counterKey, cntValue, 0)
//		_, err := resp.Result()
//		if err != nil {
//			fmt.Println("Set value error!")
//		}
//	}
//	fmt.Println("current counter is ", cntValue)
//}

type lock struct {
	cache      redis.Client
	key        string
	needUnlock bool
	isFirst    bool
}

type Locker interface {
	Lock(key string, expire time.Duration) (locked bool, err error)
	Unlock() error
}

func NewLocker(cache *redis.Client) Locker {
	return &lock{
		cache:      *cache,
		key:        "",
		needUnlock: false,
		isFirst:    false,
	}
}

func (l *lock) Lock(key string, expire time.Duration) (locked bool, err error) {
	if l.isFirst {
		return false, fmt.Errorf("repeat lock")
	}
	l.isFirst = true
	l.key = key

	uuid := xid.New().String()
	locked, err = l.cache.SetNX(l.key, uuid, expire).Result()
	if locked {
		l.needUnlock = true
	}

	return locked, err
}

func (l *lock) Unlock() error {
	if !l.needUnlock {
		return nil
	}

	return l.cache.Del(l.key).Err()
}

var counter int

func incr() {
	var lock = NewLocker(redis.NewClient(&redis.Options{Addr: "localhost:6379", Password: "cc", DB: 0}))
	locked, _ := lock.Lock("key", time.Second*5)
	fmt.Println("get lock failed:", locked)

	counter++

	if locked {
		unlock := lock.Unlock()
		if unlock != nil {
			fmt.Println(unlock)
		}
	}
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			incr()
		}()
	}
	wg.Wait()
	println("current counter is ", counter)
}
