package main

import (
	"fmt"
	"log"
	"time"

	"github.com/fadedreams/goRedisORM"
)

func main() {
	orm := goRedisORM.NewRedisORM("127.0.0.1:6379", "", 10, "prefix")

	pong, err := orm.TestClient()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(pong)

	err = orm.SetValue("username", "user100", time.Second*10)
	if err != nil {
		log.Fatal(err)
	}

	username, err := orm.GetValue("username")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Username:", username)

	// SetList example
	err = orm.SetList("mylist", time.Minute*5, "value1", "value2", "value3")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("List values set successfully.")

	// GetList example
	listValues, err := orm.GetList("mylist")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("List values:")
	for _, value := range listValues {
		fmt.Println(value)
	}

	// DeleteList example
	err = orm.DeleteList("mylist")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("List deleted successfully.")

	// Set a set with key "myset"
	setValues := []interface{}{"value1", "value2", "value3"}
	err = orm.SetSet("myset", time.Second*10, setValues...)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Set values have been set.")

	// Retrieve the set values using the key "myset"
	setResult, err := orm.GetSet("myset")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Set values:")
	for _, value := range setResult {
		fmt.Println(value)
	}

	// DeleteSet example
	err = orm.DeleteSet("myset")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Set deleted successfully.")

	// Set a hash with key "myhash"
	hashValues := map[string]interface{}{
		"field1": "value1",
		"field2": "value2",
		"field3": "value3",
	}
	err = orm.SetHash("myhash", time.Minute, hashValues)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Hash values have been set.")

	// Retrieve the hash values using the key "myhash"
	hashResult, err := orm.GetHash("myhash")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Hash values:")
	for key, value := range hashResult {
		fmt.Printf("%s: %s\n", key, value)
	}

	// DeleteHash example
	err = orm.DeleteHash("myhash")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Hash deleted successfully.")

	// SetBit example
	client := goRedisORM.NewClient("127.0.0.1:6379", "", 10)
	bit, err := goRedisORM.SetBit(client, "mykey", 0, 1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Bit value:", bit)

	// GetBit example
	bit, err = goRedisORM.GetBit(client, "mykey", 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Bit value:", bit)

	// DeleteBit example
	bit, err = goRedisORM.DeleteBit(client, "mykey", 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Bit deleted successfully.")

	// HllAdd example
	count, err := goRedisORM.HllAdd(client, "myhll", "value1", "value2", "value3")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("HyperLogLog count:", count)

	// HllCount example
	count, err = goRedisORM.HllCount(client, "myhll")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("HyperLogLog count:", count)

	// HllMerge example
	err = goRedisORM.HllMerge(client, "mergedhll", "myhll1", "myhll2", "myhll3")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("HyperLogLog merged successfully.")
}
