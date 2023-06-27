## goRedisORM

goRedisORM is a Go package that provides a simple and convenient way to interact with Redis using an Object-Relational Mapping (ORM) approach. It offers functions for setting and getting values, working with lists, sets, hashes, and more.

### Installation

To use goRedisORM in your Go project, you can install it using the following command:

```go
go get github.com/your-username/goRedisORM
```

### Usage

Import the package in your Go code:

```go
import (
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
	"github.com/your-username/goRedisORM"
)
```

### Creating a Redis client

You can create a new Redis client using the NewClient function provided by goRedisORM:

```go
orm := goRedisORM.NewRedisORM("127.0.0.1:6379", "", 10, "prefix")

```

This function takes the Redis server address, password (if any), a prefix for redis keys ,and database number as parameters. Modify the values accordingly.

### Testing the client connection

You can test the client connection to Redis using the TestClient function:

````go
``	pong, err := orm.TestClient()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(pong)
`

### Setting and getting values

You can use the SetValue and GetValue functions to set and get string values in Redis:

```go
err = orm.SetValue("username", "user100", time.Second*10)
if err != nil {
    log.Fatal(err)
}

username, err := orm.GetValue("username")
if err != nil {
    log.Fatal(err)
}
fmt.Println("Username:", username)
````

### Working with lists

To work with lists in Redis, you can use the SetList and GetList functions:

```go
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
```

### Working with sets

You can use the SetSet and GetSet functions to work with sets in Redis:

```go
err := goRedisORM.SetSet(client, "myset", "member1", "member2", "member3")
if err != nil {
	log.Fatal(err)
}

setMembers, err := goRedisORM.GetSet(client, "myset")
if err != nil {
	log.Fatal(err)
}
fmt.Println("Set Members:", setMembers)
```

### Working with hashes

To work with hashes in Redis, you can use the SetHash and GetHash functions:

```go
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
```

### Deleting keys

To delete keys from Redis, you can use the DeleteSet, DeleteList, DeleteHash, and DeleteValue functions:

```go
// DeleteList example
err = orm.DeleteList("mylist")
if err != nil {
    log.Fatal(err)
}
fmt.Println("List deleted successfully.")

// DeleteSet example
err = orm.DeleteSet("myset")
if err != nil {
    log.Fatal(err)
}
fmt.Println("Set deleted successfully.")

// DeleteHash example
err = orm.DeleteHash("myhash")
if err != nil {
    log.Fatal(err)
}
fmt.Println("Hash deleted successfully.")
```

### Working with bits

You can use the SetBit, GetBit, and DeleteBit functions to work with individual bits in Redis:

```go
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
```

### Working with HyperLogLog

You can use the HllAdd, HllCount, and HllMerge functions to work with HyperLogLog data structure in Redis:

```go
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
```

### Complete Example

You can find a complete example of goRedisORM usage in the ExampleUsage function provided. Modify the Redis server address, password, and database number according to your setup.

```go
goRedisORM.ExampleUsage()
```

### Contributing

Contributions to goRedisORM are welcome! If you find any issues or have suggestions for improvement, please open an issue or submit a pull request on the GitHub repository.
