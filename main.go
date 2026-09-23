package main

import (
	"fmt"
	"sort"
)

type Store struct {
	data map[string]string
}

func (s *Store) Keys() []string {
	keys := make([]string, 0, len(s.data))
	for key := range s.data {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	return keys
}

func (s *Store) Get(key string) (string, bool) {
	val, ok := s.data[key]
	return val, ok
}

func (s *Store) Set(key, value string) {
	s.data[key] = value
}

func (s *Store) Delete(key string) {
	delete(s.data, key)
}

func NewStore() *Store {
	return &Store{
		data: make(map[string]string),
	}
}

func main() {
	store := NewStore()
	store.Set("key1", "value1")
	store.Get("key1")
	store.Delete("key1")

	fmt.Printf("GoKV - Go key value store project")
}
