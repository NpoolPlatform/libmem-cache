package cache

import (
	"sync"
	"time"
)

type Any interface{}

type entry struct {
	expireAt time.Time
	value    Any
}

var (
	entries = map[string]*entry{}
	mutex   sync.Mutex
)

func AddEntry(key string, value Any, expire time.Duration) {
	mutex.Lock()
	defer mutex.Unlock()
	entries[key] = &entry{
		expireAt: time.Now().Add(expire),
		value:    value,
	}
}

func DelEntry(key string) {
	mutex.Lock()
	defer mutex.Unlock()
	delete(entries, key)
}

func GetEntry(key string) Any {
	mutex.Lock()
	defer mutex.Unlock()
	if ent, ok := entries[key]; ok {
		if time.Now().Before(ent.expireAt) {
			return ent.value
		}
	}
	return nil
}
