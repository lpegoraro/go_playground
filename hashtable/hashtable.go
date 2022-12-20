package hashtable

import "fmt"

type HashEntry struct {
	key    []byte
	bucket []struct {
		data interface{}
	}
}

type HashTable interface {
	hash(key string) []byte
	New(size int) *hashtable
	Insert(key string, value interface{}) (added bool, err error)
	Get(key string) (values []interface{}, found bool)
	Delete(key string) (deleted bool, err error)
	BucketSize(key string) int
}

type hashtable struct {
	table []HashEntry
}

//Simple method
func (h *hashtable) hash(key string) []byte {
	var hashed []byte
	for i := 0; i <= len(key); i++ {
		hashed[i] = key[i]
	}
	fmt.Printf("hashed %s : %s", key, hashed)
	return hashed
}

func (h *hashtable) Insert(key string, value interface{}) {
	h.hash(key)

}

func (h *hashtable) Get(key string) (values []interface{}) {
	//TODO implement me
	panic("implement me")
}

func (h *hashtable) Delete(key string) error {
	//TODO implement me
	panic("implement me")
}

func (h *hashtable) BucketSize(key string) int {
	//TODO implement me
	panic("implement me")
}
