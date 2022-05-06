package redis

import (
	"encoding/json"
	"time"

	"go.uber.org/zap"
)

var CacheDB = &cacheDB{}

type cacheDB struct {
}

func (c cacheDB) Set(key string, value interface{}, expiration int) {
	v, err := json.Marshal(value)
	if err != nil {
		zap.L().Error("[pkg: redis] [func: (c CacheDB) Set(key string, value interface{}, expiration int)] [json.Marshal(value)] failed, ", zap.Error(err))
	} else {
		rdb.Set(ctx, key, string(v), time.Second*time.Duration(expiration))
	}
}

func (c cacheDB) Get(key string, obj interface{}) bool {
	v, err := rdb.Get(ctx, key).Result()
	if err == nil && v != "" {
		err1 := json.Unmarshal([]byte(v), obj)
		return err1 == nil
	}
	return false
}

func (c cacheDB) FlushAll() {
	rdb.FlushAll(ctx)
}
