package main

import (
	"encoding/base64"
	"errors"
	"fmt"
	"sort"
)

var ErrKeyDoesNotExist = errors.New("key does not exist")
var ErrEmptyKey = errors.New("key is mandatory")
var ErrStoreFull = errors.New("store is full")

type Store struct {
	data    map[string]string
	maxSize int
}

func NewStore(maxSize int) *Store {
	return &Store{
		data:    make(map[string]string),
		maxSize: maxSize,
	}
}

func (s *Store) SetKeyWithEncryption(key, val string) (string, error) {
	encoded := base64.StdEncoding.EncodeToString([]byte(val))
	if err := s.Set(key, encoded); err != nil {
		return "", err
	}
	return s.Get(key)
}

func (s *Store) Keys() []string {
	keys := make([]string, 0, len(s.data))
	for key := range s.data {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	return keys
}

func (s *Store) Get(key string) (string, error) {
	if key == "" {
		return "", ErrEmptyKey
	}

	val, ok := s.data[key]
	if !ok {
		return "", ErrKeyDoesNotExist
	}
	return val, nil
}

func (s *Store) Set(key, value string) error {
	if key == "" {
		return ErrEmptyKey
	}

	_, exists := s.data[key]
	if s.maxSize > 0 && s.Len() >= s.maxSize && !exists {
		return fmt.Errorf("Set(%q): %w", key, ErrStoreFull)
	}

	s.data[key] = value
	return nil
}

func (s *Store) Delete(key string) {
	delete(s.data, key)
}

func (s *Store) Len() int {
	return len(s.data)
}
