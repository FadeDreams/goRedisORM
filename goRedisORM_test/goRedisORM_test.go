package goRedisORM_test

import (
	"testing"

	"github.com/fadedreams/goRedisORM" // Replace "your-username" with your actual username
	"github.com/go-redis/redis/v8"
)

func TestSetValue(t *testing.T) {
	client := getTestClient()

	err := goRedisORM.SetValue(client, "test-key", "test-value")
	if err != nil {
		t.Errorf("SetValue returned an error: %s", err.Error())
	}

	value, err := goRedisORM.GetValue(client, "test-key")
	if err != nil {
		t.Errorf("GetValue returned an error: %s", err.Error())
	}

	if value != "test-value" {
		t.Errorf("Expected value: %s, got: %s", "test-value", value)
	}
}

func TestSetList(t *testing.T) {
	client := getTestClient()

	key := "mylist"
	expectedValues := []interface{}{"value1", "value2", "value3"}

	// Delete the list if it already exists
	err := goRedisORM.DeleteList(client, key)
	if err != nil {
		t.Fatal(err)
	}

	err = goRedisORM.SetList(client, key, expectedValues...)
	if err != nil {
		t.Fatal(err)
	}

	actualValues, err := goRedisORM.GetList(client, key)
	if err != nil {
		t.Fatal(err)
	}

	if len(actualValues) != len(expectedValues) {
		t.Fatalf("Expected list length: %d, got: %d", len(expectedValues), len(actualValues))
	}

	for i := range expectedValues {
		if expectedValues[i] != actualValues[i] {
			t.Errorf("Mismatch at index %d. Expected: %v, got: %v", i, expectedValues[i], actualValues[i])
		}
	}
}

func TestSetHash(t *testing.T) {
	client := getTestClient()

	hashValues := map[string]interface{}{
		"key1": "value1",
		"key2": "value2",
		"key3": "value3",
	}

	err := goRedisORM.SetHash(client, "test-hash", hashValues)
	if err != nil {
		t.Errorf("SetHash returned an error: %s", err.Error())
	}

	hashResult, err := goRedisORM.GetHash(client, "test-hash")
	if err != nil {
		t.Errorf("GetHash returned an error: %s", err.Error())
	}

	expected := map[string]string{
		"key1": "value1",
		"key2": "value2",
		"key3": "value3",
	}
	if !equalStringStringMaps(hashResult, expected) {
		t.Errorf("Expected hash result: %v, got: %v", expected, hashResult)
	}
}

func TestSetSet(t *testing.T) {
	client := getTestClient()

	err := goRedisORM.SetSet(client, "test-set", "member1", "member2", "member3")
	if err != nil {
		t.Errorf("SetSet returned an error: %s", err.Error())
	}

	setMembers, err := goRedisORM.GetSet(client, "test-set")
	if err != nil {
		t.Errorf("GetSet returned an error: %s", err.Error())
	}

	expected := []string{"member3", "member2", "member1"}
	if !equalStringSlices(setMembers, expected) {
		t.Errorf("Expected set members: %v, got: %v", expected, setMembers)
	}
}

func TestDeleteSet(t *testing.T) {
	client := getTestClient()

	err := goRedisORM.SetSet(client, "test-set", "member1", "member2", "member3")
	if err != nil {
		t.Errorf("SetSet returned an error: %s", err.Error())
	}

	err = goRedisORM.DeleteSet(client, "test-set")
	if err != nil {
		t.Errorf("DeleteSet returned an error: %s", err.Error())
	}

	setMembers, err := goRedisORM.GetSet(client, "test-set")
	if err != nil {
		t.Errorf("GetSet returned an error: %s", err.Error())
	}

	if len(setMembers) != 0 {
		t.Errorf("Expected no set members, got: %v", setMembers)
	}
}

// Add more test cases for other functions as needed

// Helper function to create a test Redis client
func getTestClient() *redis.Client {
	return goRedisORM.NewClient("127.0.0.1:6379", "", 10)
}

// Helper function to compare two string slices
func equalStringSlices(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}

// Helper function to compare two string-string maps
func equalStringStringMaps(a, b map[string]string) bool {
	if len(a) != len(b) {
		return false
	}

	for k, v := range a {
		if bVal, ok := b[k]; !ok || v != bVal {
			return false
		}
	}

	return true
}
