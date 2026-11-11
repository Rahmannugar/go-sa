package main

import (
	"fmt"
)

const (
	maxLoadFactor = 0.75
	fnvOffset     = uint64(14695981039346656037)
	fnvPrime      = uint64(1099511628211)
)

type Entry struct {
	key   string
	value any
}

type HashMap struct {
	buckets [][]Entry
	size    int
}

// NewHashMap creates a hash map with the given number of buckets.
// Time complexity: O(capacity), Space complexity: O(capacity)
func NewHashMap(capacity int) *HashMap {
	if capacity <= 0 {
		panic("hashmap capacity must be greater than zero")
	}

	return &HashMap{
		buckets: make([][]Entry, capacity),
	}
}

// hash generates a 64-bit hash from a string key.
// Time complexity: O(k), Space complexity: O(1)
func (h *HashMap) hash(key string) uint64 {
	hash := fnvOffset

	for i := 0; i < len(key); i++ {
		hash ^= uint64(key[i])
		hash *= fnvPrime
	}

	return hash
}

// keyToIndex maps a string key to a bucket index.
// Time complexity: O(k), Space complexity: O(1)
func (h *HashMap) keyToIndex(key string) int {
	return int(h.hash(key) % uint64(len(h.buckets)))
}

// LoadFactor returns the number of entries divided by the number of buckets.
// Time complexity: O(1), Space complexity: O(1)
func (h *HashMap) LoadFactor() float64 {
	return float64(h.size) / float64(len(h.buckets))
}

// Len returns the number of entries.
// Time complexity: O(1), Space complexity: O(1)
func (h *HashMap) Len() int {
	return h.size
}

// Capacity returns the number of buckets.
// Time complexity: O(1), Space complexity: O(1)
func (h *HashMap) Capacity() int {
	return len(h.buckets)
}

// resize doubles the bucket count and redistributes existing entries.
// Time complexity: O(nk), Space complexity: O(n)
func (h *HashMap) resize() {
	oldBuckets := h.buckets

	newCapacity := len(oldBuckets) * 2
	h.buckets = make([][]Entry, newCapacity)

	for _, bucket := range oldBuckets {
		for _, entry := range bucket {
			index := h.keyToIndex(entry.key)

			h.buckets[index] = append(
				h.buckets[index],
				entry,
			)
		}
	}
}

// Insert adds a key-value pair or updates an existing key.
// Time complexity: O(k) average amortized, Space complexity: O(n) on resize
func (h *HashMap) Insert(key string, value any) {
	index := h.keyToIndex(key)

	// Update an existing entry if the key already exists.
	for i, entry := range h.buckets[index] {
		if entry.key == key {
			h.buckets[index][i].value = value
			return
		}
	}

	// Resize before inserting if the load factor would exceed 0.75.
	if float64(h.size+1)/float64(len(h.buckets)) > maxLoadFactor {
		h.resize()
		index = h.keyToIndex(key)
	}

	entry := Entry{
		key:   key,
		value: value,
	}

	h.buckets[index] = append(h.buckets[index], entry)
	h.size++
}

// Get returns the value for a key, or false if the key is absent.
// Time complexity: O(k) average, Space complexity: O(1)
func (h *HashMap) Get(key string) (any, bool) {
	index := h.keyToIndex(key)

	for _, entry := range h.buckets[index] {
		if entry.key == key {
			return entry.value, true
		}
	}

	return nil, false
}

// Delete removes a key and reports whether it was present.
// Time complexity: O(k) average, Space complexity: O(1)
func (h *HashMap) Delete(key string) bool {
	index := h.keyToIndex(key)

	bucket := h.buckets[index]

	for i, entry := range bucket {
		if entry.key == key {
			copy(bucket[i:], bucket[i+1:])

			bucket[len(bucket)-1] = Entry{}
			h.buckets[index] = bucket[:len(bucket)-1]

			h.size--
			return true
		}
	}

	return false
}

// BucketSize returns the number of entries in a key's bucket.
// Time complexity: O(k), Space complexity: O(1)
func (h *HashMap) BucketSize(key string) int {
	index := h.keyToIndex(key)
	return len(h.buckets[index])
}

type User struct {
	ID   int
	Name string
}

func main() {
	hashmap := NewHashMap(4)

	hashmap.Insert("cat", User{ID: 1, Name: "John"})
	hashmap.Insert("tac", User{ID: 2, Name: "Jane"})
	hashmap.Insert("dog", User{ID: 3, Name: "Mike"})
	hashmap.Insert("john", User{ID: 4, Name: "David"})

	fmt.Println("Entries:", hashmap.Len())
	fmt.Println("Capacity:", hashmap.Capacity())
	fmt.Println("Load factor:", hashmap.LoadFactor())

	// Retrieve a user.
	value, ok := hashmap.Get("cat")
	if ok {
		fmt.Println("cat:", value)
	}

	// Update an existing key.
	hashmap.Insert("cat", User{ID: 5, Name: "Updated John"})

	value, ok = hashmap.Get("cat")
	if ok {
		fmt.Println("Updated cat:", value)
	}

	// Inspect bucket occupancy.
	fmt.Println("cat bucket size:", hashmap.BucketSize("cat"))
	fmt.Println("tac bucket size:", hashmap.BucketSize("tac"))

	// Delete an entry.
	fmt.Println("Deleted tac:", hashmap.Delete("tac"))

	fmt.Println("Entries after deletion:", hashmap.Len())
	fmt.Println("Capacity after deletion:", hashmap.Capacity())
}
