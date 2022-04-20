package localcache

import (
	"errors"
	"time"
)

const expiredTime = 30 * time.Second

// ErrNotFound will throw when user get cache by key, but key is not exist in cache
var ErrNotFound = errors.New("fail to get cache by key")

type LocalCache struct {
	store map[string]CacheItem
}

type CacheItem struct {
	value       interface{}
	expireTimer *time.Timer
}

func New() Cache {
	instance := &LocalCache{
		store: map[string]CacheItem{},
	}
	return instance
}

func (lc *LocalCache) Get(key string) (value interface{}, e error) {
	cacheItem, ok := lc.store[key]
	if ok {
		value = cacheItem.value
		return
	}
	e = ErrNotFound
	return
}

func (lc *LocalCache) Set(key string, value interface{}) error {
	lc.store[key] = CacheItem{
		value: value,
		expireTimer: time.AfterFunc(expiredTime, func() {
			lc.clean(key)
		}),
	}
	return nil
}

func (lc *LocalCache) clean(key string) {
	delete(lc.store, key)
}
