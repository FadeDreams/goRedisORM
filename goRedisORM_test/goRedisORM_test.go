package goRedisORM_test

import (
	"fmt"
	"log"
	"sort"
	"testing"
	"time"

	"github.com/fadedreams/goRedisORM"
	"github.com/go-redis/redis/v8"
)

func TestSetValueGetValue(t *testing.T) {
	// Create a Redis client for testing
	client := NewClient("localhost:6379", "", 0)
	orm := &goRedisORM.RedisORM{
		Client: client,
		Prefix: "test",
	}

	// Set a value
	key := "key1"
	value := "value1"
	expiration := time.Hour
	err := orm.SetValue(key, value, expiration)
	if err != nil {
		t.Errorf("Error setting value: %v", err)
	}

	// Get the value and check if it matches the expected value
	gotValue, err := orm.GetValue(key)
	if err != nil {
		t.Errorf("Error getting value: %v", err)
	}
	if gotValue != value {
		t.Errorf("Expected value %s, got %s", value, gotValue)
	}

	// Cleanup: Delete the key
	err = orm.DeleteValue(orm.AddKeyPrefix(key))
	if err != nil {
		t.Errorf("Error deleting value: %v", err)
	}
}

func NewClient(addr string, password string, db int) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
	return client
}

func TestSetGetList(t *testing.T) {
	// Create a Redis client for testing
	client := NewClient("localhost:6379", "", 0)
	orm := &goRedisORM.RedisORM{
		Client: client,
		Prefix: "test",
	}

	// Set a list of values
	key := "list1"
	values := []interface{}{"value1", "value2", "value3"}
	expiration := time.Hour
	err := orm.SetList(key, expiration, values...)
	if err != nil {
		t.Errorf("Error setting list: %v", err)
	}

	// Get the list and check if it matches the expected values
	gotValues, err := orm.GetList(key)
	if err != nil {
		t.Errorf("Error getting list: %v", err)
	}

	if len(gotValues) != len(values) {
		t.Errorf("Expected list length %d, got %d", len(values), len(gotValues))
	}

	for i := range values {
		if values[i] != gotValues[i] {
			t.Errorf("Expected value %s, got %s", values[i], gotValues[i])
		}
	}
	// Cleanup: Delete the key
	err = orm.DeleteValue(orm.AddKeyPrefix(key))
	if err != nil {
		t.Errorf("Error deleting list: %v", err)
	}
}

func TestSetGetSet(t *testing.T) {
	// Create a Redis client for testing
	client := NewClient("localhost:6379", "", 0)
	orm := &goRedisORM.RedisORM{
		Client: client,
		Prefix: "test",
	}

	// Set a set of members
	key := "set1"
	members := []interface{}{"member1", "member2", "member3"}
	expiration := time.Hour
	err := orm.SetSet(key, expiration, members...)
	if err != nil {
		t.Errorf("Error setting set: %v", err)
	}

	// Get the set and check if it matches the expected members
	gotMembers, err := orm.GetSet(key)
	if err != nil {
		t.Errorf("Error getting set: %v", err)
	}

	// Sort the members for comparison
	sort.Strings(gotMembers)
	expectedMembers := []string{"member1", "member2", "member3"}
	sort.Strings(expectedMembers)

	if len(gotMembers) != len(expectedMembers) {
		t.Errorf("Expected set length %d, got %d", len(expectedMembers), len(gotMembers))
	}

	for i := range expectedMembers {
		if expectedMembers[i] != gotMembers[i] {
			t.Errorf("Expected member %s, got %s", expectedMembers[i], gotMembers[i])
		}
	}

	// Cleanup: Delete the key
	err = orm.DeleteValue(orm.AddKeyPrefix(key))
	if err != nil {
		t.Errorf("Error deleting list: %v", err)
	}
}

func TestSetGetBit(t *testing.T) {
	// Create a Redis client for testing
	client := NewClient("localhost:6379", "", 0)

	// Set a bit
	key := "bit1"
	offset := int64(10)
	value := 1
	_, err := goRedisORM.SetBit(client, key, offset, value)
	if err != nil {
		t.Errorf("Error setting bit: %v", err)
	}

	// Get the bit and check if it matches the expected value
	gotBit, err := goRedisORM.GetBit(client, key, offset)
	if err != nil {
		t.Errorf("Error getting bit: %v", err)
	}

	if gotBit != int64(value) {
		t.Errorf("Expected bit value %d, got %d", value, gotBit)
	}

	// Cleanup: Delete the bit
	_, err = goRedisORM.DeleteBit(client, key, offset)
	if err != nil {
		t.Errorf("Error deleting bit: %v", err)
	}
}

func TestHyperLogLogOperations(t *testing.T) {
	// Create a Redis client
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	// Flush any previous data in Redis
	err := client.FlushDB(client.Context()).Err()
	if err != nil {
		log.Fatalf("Failed to flush Redis DB: %v", err)
	}

	// Test HllAdd function
	count, err := goRedisORM.HllAdd(client, "hll_key", "value1", "value2", "value3")
	if err != nil {
		t.Errorf("HllAdd failed: %v", err)
	}
	fmt.Printf("Added %d items to HyperLogLog\n", count)

	// Test HllCount function
	count, err = goRedisORM.HllCount(client, "hll_key")
	if err != nil {
		t.Errorf("HllCount failed: %v", err)
	}
	fmt.Printf("HyperLogLog count: %d\n", count)

	// Test HllMerge function
	err = goRedisORM.HllMerge(client, "hll_dest", "hll_key")
	if err != nil {
		t.Errorf("HllMerge failed: %v", err)
	}
	fmt.Println("Merged HyperLogLogs")

	// Test HllCount after merge
	count, err = goRedisORM.HllCount(client, "hll_dest")
	if err != nil {
		t.Errorf("HllCount failed after merge: %v", err)
	}
	fmt.Printf("HyperLogLog count after merge: %d\n", count)
}
