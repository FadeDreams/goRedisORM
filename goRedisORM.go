package goRedisORM

import (
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisORM struct {
	Client *redis.Client
	Prefix string
}

func NewRedisORM(addr string, password string, db int, prefix string) *RedisORM {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
	return &RedisORM{
		Client: client,
		Prefix: prefix,
	}
}

func (orm *RedisORM) AddKeyPrefix(key string) string {
	return fmt.Sprintf("%s:%s", orm.Prefix, key)
}

func (orm *RedisORM) SetValue(key string, value string, expiration time.Duration) error {
	err := orm.Client.Set(orm.Client.Context(), orm.AddKeyPrefix(key), value, expiration).Err()
	if err != nil {
		return err
	}
	return nil
}

func (orm *RedisORM) GetValue(key string) (string, error) {
	value, err := orm.Client.Get(orm.Client.Context(), orm.AddKeyPrefix(key)).Result()
	if err != nil {
		return "", err
	}
	return value, nil
}

func (orm *RedisORM) TestClient() (string, error) {
	pong, err := orm.Client.Ping(orm.Client.Context()).Result()
	if err != nil {
		return "", err
	}
	return pong, nil
}

func (orm *RedisORM) SetList(key string, expiration time.Duration, values ...interface{}) error {
	err := orm.Client.RPush(orm.Client.Context(), orm.AddKeyPrefix(key), values...).Err()
	if err != nil {
		return err
	}
	err = orm.Client.Expire(orm.Client.Context(), orm.AddKeyPrefix(key), expiration).Err()
	if err != nil {
		return err
	}
	return nil
}

func (orm *RedisORM) GetList(key string) ([]string, error) {
	values, err := orm.Client.LRange(orm.Client.Context(), orm.AddKeyPrefix(key), 0, -1).Result()
	if err != nil {
		return nil, err
	}
	return values, nil
}

func (orm *RedisORM) SetSet(key string, expiration time.Duration, members ...interface{}) error {
	err := orm.Client.SAdd(orm.Client.Context(), orm.AddKeyPrefix(key), members...).Err()
	if err != nil {
		return err
	}
	err = orm.Client.Expire(orm.Client.Context(), orm.AddKeyPrefix(key), expiration).Err()
	if err != nil {
		return err
	}
	return nil
}

func (orm *RedisORM) GetSet(key string) ([]string, error) {
	members, err := orm.Client.SMembers(orm.Client.Context(), orm.AddKeyPrefix(key)).Result()
	if err != nil {
		return nil, err
	}
	return members, nil
}

func (orm *RedisORM) SetHash(key string, expiration time.Duration, values map[string]interface{}) error {
	err := orm.Client.HMSet(orm.Client.Context(), orm.AddKeyPrefix(key), values).Err()
	if err != nil {
		return err
	}
	err = orm.Client.Expire(orm.Client.Context(), orm.AddKeyPrefix(key), expiration).Err()
	if err != nil {
		return err
	}
	return nil
}

func (orm *RedisORM) GetHash(key string) (map[string]string, error) {
	values, err := orm.Client.HGetAll(orm.Client.Context(), orm.AddKeyPrefix(key)).Result()
	if err != nil {
		return nil, err
	}
	return values, nil
}

func (orm *RedisORM) DeleteSet(key string) error {
	err := orm.Client.Del(orm.Client.Context(), orm.AddKeyPrefix(key)).Err()
	if err != nil {
		return err
	}
	return nil
}

func (orm *RedisORM) DeleteList(key string) error {
	err := orm.Client.Del(orm.Client.Context(), orm.AddKeyPrefix(key)).Err()
	if err != nil {
		return err
	}
	return nil
}

func (orm *RedisORM) DeleteHash(key string) error {
	err := orm.Client.Del(orm.Client.Context(), orm.AddKeyPrefix(key)).Err()
	if err != nil {
		return err
	}
	return nil
}

func (orm *RedisORM) DeleteValue(key string) error {
	err := orm.Client.Del(orm.Client.Context(), orm.AddKeyPrefix(key)).Err()
	if err != nil {
		return err
	}
	return nil
}

func SetBit(client *redis.Client, key string, offset int64, value int) (int64, error) {
	bit, err := client.SetBit(client.Context(), key, offset, value).Result()
	if err != nil {
		return 0, err
	}
	return bit, nil
}

func GetBit(client *redis.Client, key string, offset int64) (int64, error) {
	bit, err := client.GetBit(client.Context(), key, offset).Result()
	if err != nil {
		return 0, err
	}
	return bit, nil
}

func DeleteBit(client *redis.Client, key string, offset int64) (int64, error) {
	bit, err := client.SetBit(client.Context(), key, offset, 0).Result()
	if err != nil {
		return 0, err
	}
	return bit, nil
}

func HllAdd(client *redis.Client, key string, values ...interface{}) (int64, error) {
	count, err := client.PFAdd(client.Context(), key, values...).Result()
	if err != nil {
		return 0, err
	}
	return count, nil
}

func HllCount(client *redis.Client, keys ...string) (int64, error) {
	count, err := client.PFCount(client.Context(), keys...).Result()
	if err != nil {
		return 0, err
	}
	return count, nil
}

func HllMerge(client *redis.Client, dest string, keys ...string) error {
	err := client.PFMerge(client.Context(), dest, keys...).Err()
	if err != nil {
		return err
	}
	return nil
}

func NewClient(addr string, password string, db int) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
	return client
}
