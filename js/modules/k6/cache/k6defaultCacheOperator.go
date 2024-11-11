package k6cache

func (k6cache *K6Cache) Store(key string, value string) {
	k6cache.defaultCache.Store(key, value)
}

func (k6cache *K6Cache) Load(key string) interface{} {
	stringValue, _ := k6cache.defaultCache.Load(key)
	return stringValue
}

func (k6cache *K6Cache) Delete(key string) {
	k6cache.defaultCache.Delete(key)
}
