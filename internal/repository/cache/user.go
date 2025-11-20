package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/ILk8S/basc-go/internal/domain"
	"github.com/redis/go-redis/v9"
)

var ErrKeyNotExist = redis.Nil

type UserCache struct {
	cmd redis.Cmdable
	//过期时间
	expiration time.Duration
}

func NewUserCache(cmd redis.Cmdable) *UserCache {
	return &UserCache{
		cmd:        cmd,
		expiration: time.Minute * 15,
	}
}

func (cache *UserCache) Get(ctx context.Context, uid int64) (domain.User, error) {
	key := cache.key(uid)
	data, err := cache.cmd.Get(ctx, key).Result()
	if err != nil {
		return domain.User{}, err
	}
	// 把查出来的value反序列化成user对象
	var u domain.User
	err = json.Unmarshal([]byte(data), &u)
	return u, err
}

func (cache *UserCache) Set(ctx context.Context, u domain.User) error {
	data, err := json.Marshal(u)
	if err != nil {
		return err
	}
	key := cache.key(u.Id)
	return cache.cmd.Set(ctx, key, data, cache.expiration).Err()
}
func (cache *UserCache) key(uid int64) string {
	return fmt.Sprintf("user:info:%d", uid)
}
