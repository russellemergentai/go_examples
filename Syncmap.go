package main

import (
	"fmt"
	"sync"
)

func main() {
	// Create a new sync.Map
	m := sync.Map{}

	// Add key-value pairs to the map
	m.Store("key1", "value1")
	m.Store("key2", "value2")

	// Retrieve a value by key
	val, found := m.Load("key1")
	if found {
		fmt.Printf("Key1: %v\n", val)
	}

	// Delete a key-value pair
	m.Delete("key2")

	// Check if a key exists
	_, found = m.Load("key2")
	if !found {
		fmt.Println("Key2 does not exist")
	}

	// Range over the map
	m.Range(func(key, value interface{}) bool {
		fmt.Printf("Key: %v, Value: %v\n", key, value)
		return true
	})
}
