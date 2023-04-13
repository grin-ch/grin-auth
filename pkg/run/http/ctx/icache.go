package ctx

import "sync"

type ICache[K comparable, V any] interface {
	Set(K, V)
	Get(K) (V, bool)
	Del(K)
}

func newCtxCache[K comparable, V any]() *ctxCache[K, V] {
	return &ctxCache[K, V]{
		cmap: map[K]V{},
	}
}

type ctxCache[K comparable, V any] struct {
	cmap map[K]V
	sync.Mutex
}

func (cache *ctxCache[K, V]) Set(k K, v V) {
	cache.Lock()
	cache.cmap[k] = v
	cache.Unlock()
}

func (cache *ctxCache[K, V]) Get(k K) (V, bool) {
	cache.Lock()
	v, has := cache.cmap[k]
	cache.Unlock()
	return v, has
}

func (cache *ctxCache[K, V]) Del(k K) {
	cache.Lock()
	delete(cache.cmap, k)
	cache.Unlock()
}
