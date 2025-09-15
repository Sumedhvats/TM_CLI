package db

import (
	"encoding/binary"
	"time"

	"github.com/boltdb/bolt"
)

var taskBucket = []byte("tasks")
var db *bolt.DB

type Tasks struct {
	Key   int
	Value string
}

func Init(dbPath string) error {
	var err error
	db, err = bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return err

	}
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
		return b.Put(itob(id64), []byte(task))
	})
	if err != nil {
		return -1, err
	}

	return id, nil
}
func AllTask() ([]Tasks, error) {
	var tasks []Tasks
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			tasks = append(tasks, Tasks{
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
func DeleteTask(key int)error{
	return db.Update(func(tx *bolt.Tx) error {
		b:=tx.Bucket(taskBucket)
		return b.Delete(itob(uint64(key)))

	})

}
func itob(i uint64) (buff []byte) {
	buff = make([]byte, 8)
	binary.BigEndian.PutUint64(buff, i)
	return
}
func btoi(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}
