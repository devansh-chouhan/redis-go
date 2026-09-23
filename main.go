package main

import (
	"fmt"
)

func main() {
	store := NewStore()
	store.Set("key1", "value1")
	store.Get("key1")

	fmt.Printf("GoKV - Go key value store project")
}
