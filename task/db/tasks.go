package db

import (
	"encoding/binary"
	"time"

	"github.com/boltdb/bolt"
)

var taskBucket = []byte("tasks")

var db *bolt.DB

type Task struct {
	Key   int
	Value string
}

func Init(dbPath string) error {
	// There, we declare error explicitly instead of using "db. err  := "
	// That's why, by doing this, db will have a local scope, but we want to refer to the global "db"
	// variable declared above
	var err error
	db, err = bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return err
	}

	// Create a bucket
	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(taskBucket)
		return err
	})
}

func CreateTask(task string) (int, error) {
	var id int
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		id64, _ := b.NextSequence()
		id = int(id64)
		key := itob(id)
		return b.Put(key, []byte(task))
	})
	if err != nil {
		return -1, err
	}
	return id, nil
}

// Regarding the read operation, we can either read individually or all of them
func ReadAll() ([]Task, error) {
	var tasks []Task
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			tasks = append(tasks, Task{
				Key:   btoi(k),
				Value: string(v),
			})
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func DeleteTask(key int) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		return b.Delete(itob(key))
	})
}

// Take an integer and return the corresponding byte slice
func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

func btoi(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}

// import (
// 	"log"
// 	"time"

// 	"github.com/boltdb/bolt"
// )

// func main() {
// 	// Connect to the database
// 	db, err := bolt.Open("my.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer db.Close()
// }

// /*
// <UserBucket>
// key		|	value
// ID123	| 	some bucket

// <someBucket>
// key		|	value
// name	|	"John Calhoun"
// Email	|	"jon@calhoun.io"

// You can also use Storm package in order to encode a simple Go struct into Bolt stuff
// */
