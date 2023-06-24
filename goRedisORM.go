package goRedisORM

import (
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
)

func SetValue(client *redis.Client, key string, value string) error {
	err := client.Set(client.Context(), key, value, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

func GetValue(client *redis.Client, key string) (string, error) {
	value, err := client.Get(client.Context(), key).Result()
	if err != nil {
		return "", err
	}
	return value, nil
}

func TestClient(client *redis.Client) (string, error) {
	pong, err := client.Ping(client.Context()).Result()
	if err != nil {
		return "", err
	}
	return pong, nil
}

func SetList(client *redis.Client, key string, values ...interface{}) error {
	err := client.RPush(client.Context(), key, values...).Err()
	if err != nil {
		return err
	}
	return nil
}

func GetList(client *redis.Client, key string) ([]string, error) {
	values, err := client.LRange(client.Context(), key, 0, -1).Result()
	if err != nil {
		return nil, err
	}
	return values, nil
}

func SetSet(client *redis.Client, key string, members ...interface{}) error {
	err := client.SAdd(client.Context(), key, members...).Err()
	if err != nil {
		return err
	}
	return nil
}

func GetSet(client *redis.Client, key string) ([]string, error) {
	members, err := client.SMembers(client.Context(), key).Result()
	if err != nil {
		return nil, err
	}
	return members, nil
}

func SetHash(client *redis.Client, key string, values map[string]interface{}) error {
	err := client.HMSet(client.Context(), key, values).Err()
	if err != nil {
		return err
	}
	return nil
}

func GetHash(client *redis.Client, key string) (map[string]string, error) {
	values, err := client.HGetAll(client.Context(), key).Result()
	if err != nil {
		return nil, err
	}
	return values, nil
}

func DeleteSet(client *redis.Client, key string) error {
	err := client.Del(client.Context(), key).Err()
	if err != nil {
		return err
	}
	return nil
}

func DeleteList(client *redis.Client, key string) error {
	err := client.Del(client.Context(), key).Err()
	if err != nil {
		return err
	}
	return nil
}

func DeleteHash(client *redis.Client, key string) error {
	err := client.Del(client.Context(), key).Err()
	if err != nil {
		return err
	}
	return nil
}

func DeleteValue(client *redis.Client, key string) error {
	err := client.Del(client.Context(), key).Err()
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

func ExampleUsage() {
	client := NewClient("127.0.0.1:6379", "", 10)

	pong, err := TestClient(client)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(pong)

	err = SetValue(client, "username", "user100")
	if err != nil {
		log.Fatal(err)
	}

	username, err := GetValue(client, "username")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Username:", username)

	err = SetList(client, "mylist", "value1", "value2", "value3")
	if err != nil {
		log.Fatal(err)
	}

	listValues, err := GetList(client, "mylist")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("List Values:", listValues)

	hashValues := map[string]interface{}{
		"key1": "value1",
		"key2": "value2",
		"key3": "value3",
	}

	err = SetHash(client, "myhash", hashValues)
	if err != nil {
		log.Fatal(err)
	}

	hashResult, err := GetHash(client, "myhash")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Hash Result:", hashResult)
	for key, value := range hashResult {
		fmt.Println("Key:", key, "Value:", value)
	}

	err = SetSet(client, "myset", "member1", "member2", "member3")
	if err != nil {
		log.Fatal(err)
	}

	setMembers, err := GetSet(client, "myset")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Set Members:", setMembers)

	err = DeleteSet(client, "myset")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Set deleted.")

	err = DeleteList(client, "mylist")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("List deleted.")

	err = DeleteHash(client, "myhash")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Hash deleted.")

	err = DeleteValue(client, "username")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Value deleted.")

	bit, err := SetBit(client, "mykey", 0, 1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Bit set:", bit)

	bit, err = GetBit(client, "mykey", 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Bit:", bit)

	bit, err = DeleteBit(client, "mykey", 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Bit deleted:", bit)

	count, err := HllAdd(client, "hllkey", "value1", "value2", "value3")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("HLL Add Count:", count)

	hllCount, err := HllCount(client, "hllkey")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("HLL Count:", hllCount)

	err = HllMerge(client, "hllmerged", "hllkey", "hllkey2")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("HLL merged.")
}
