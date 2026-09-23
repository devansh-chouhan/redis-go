package main

import (
	"fmt"
)

func main() {
	store := NewStore()
	store.Set("key1", "value1")
	store.Delete("key1")
	val, err := store.Get("key1")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(val)

	fmt.Printf("GoKV - Go key value store project")
}
