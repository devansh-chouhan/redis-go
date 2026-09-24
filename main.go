package main

import (
	"fmt"
)

func main() {
	store := NewStore(20)
	store.Set("a", "42")
	store.Set("b", "7")

	fmt.Printf("GoKV - Go key value store project")
}
