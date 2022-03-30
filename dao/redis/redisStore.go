package redis

import (
	"fmt"
	"time"
)

/*
	实现captcha的Store接口，将验证码保存至redis
*/

const CAPTCHA = "ziweiShop:admin:captcha_id:"

type RedisStore struct {
}

func (r RedisStore) Set(id string, value string) (err error) {
	key := CAPTCHA + id
	err = rdb.Set(ctx, key, value, time.Minute*2).Err()
	return
}

func (r RedisStore) Get(id string, clear bool) string {
	key := CAPTCHA + id
	value, err := rdb.Get(ctx, key).Result()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	if clear {
		if err := rdb.Del(ctx, key).Err(); err != nil {
			fmt.Println(err)
			return ""
		}
	}
	return value
}

func (r RedisStore) Verify(id, answer string, clear bool) bool {
	return answer == r.Get(id, clear)
}
