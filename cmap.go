package cmap

import (
	"hash/fnv"
	"sync"
)

var BUCKET_COUNT uint32 = 16
var bucket_mask = uint32(BUCKET_COUNT - 1)

type Bucket struct {
	sync.RWMutex
	lookup map[string]interface{}
}

type CMap []*Bucket

func New() CMap {
	m := make(CMap, 16)
	for i := uint32(0); i < BUCKET_COUNT; i++ {
		m[i] = &Bucket{lookup: make(map[string]interface{})}
	}
	return m
}

func (c CMap) Get(key string) (interface{}, bool) {
	bucket := c.bucket(key)
	bucket.RLock()
	v, b := bucket.lookup[key]
	bucket.RUnlock()
	return v, b
}

func (c CMap) Set(key string, value interface{}) {
	bucket := c.bucket(key)
	bucket.Lock()
	bucket.lookup[key] = value
	bucket.Unlock()
}

func (c CMap) Delete(key string) {
	bucket := c.bucket(key)
	bucket.Lock()
	delete(bucket.lookup, key)
	bucket.Unlock()
}

func (c CMap) bucket(key string) *Bucket {
	h := fnv.New32a()
	h.Write([]byte(key))
	return c[h.Sum32()&bucket_mask]
}

func (c CMap) Len() int {
	l := 0
	for i := uint32(0); i < BUCKET_COUNT; i++ {
		bucket := c[i]
		bucket.RLock()
		l += len(bucket.lookup)
		bucket.RUnlock()
	}
	return l
}
