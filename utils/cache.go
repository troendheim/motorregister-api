package utils

import (
	"github.com/coocood/freecache"
	"runtime/debug"
)


var Cache *freecache.Cache

func StartCache() {
	Cache = freecache.NewCache(100 * 1024 * 1024) // 100 MiB
	debug.SetGCPercent(20)
}

func GetFromCache(key string) (string, error) {
	var cacheKey = []byte(key)
	var cacheResult, cacheErr = Cache.Get(cacheKey)

	return string(cacheResult), cacheErr
}

func SetCache(key string, data string) {
	var cacheKey = ConvertStringToByteSlice(key)
	var result = ConvertStringToByteSlice(data)

	Cache.Set(cacheKey, result, 86400)
}