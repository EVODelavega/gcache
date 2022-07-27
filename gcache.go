package gcache

import (
	"sync"
	"time"
)

type cmp interface {
	comparable
}

type val[T any] struct {
	val   T
	ttl   time.Duration
	valid time.Time
}

type cache[K comparable, V any] struct {
	rmu  *sync.RWMutex
	data map[K]val[V]
}

func New[K comparable, V any]() *cache[K, V] {
	return NewWithVals[K, V](nil)
}

func NewWithVals[K comparable, V any](vals map[K]V) *cache[K, V] {
	data := make(map[K]val[V], len(vals))
	for k, v := range vals {
		data[k] = val[V]{
			val: v,
		}
	}
	return &cache[K, V]{
		rmu:  &sync.RWMutex{},
		data: data,
	}
}

func (c *cache[K, V]) Has(k K) bool {
	c.rmu.RLock()
	defer c.rmu.RUnlock()
	_, ok := c.data[k]
	return ok
}

func (c *cache[K, V]) Get(k K) (V, bool) {
	c.rmu.RLock()
	defer c.rmu.RUnlock()
	v, ok := c.data[k]
	return v.val, ok
}

func (c *cache[K, V]) Set(k K, v V) {
	c.rmu.Lock()
	defer c.rmu.Unlock()
	c.data[k] = val[V]{
		val: v,
	}
}
