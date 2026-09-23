package main

import (
	"reflect"
	"testing"
)

func TestKeys_ReturnsAllKeysSorted(t *testing.T) {
	store := NewStore()
	store.Set("charlie", "3")
	store.Set("alpha", "1")
	store.Set("bravo", "2")

	got := store.Keys()
	want := []string{"alpha", "bravo", "charlie"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Keys() = %v, want %v", got, want)
	}
}

func TestKeys_EmptyStore(t *testing.T) {
	store := NewStore()

	got := store.Keys()
	if len(got) != 0 {
		t.Errorf("Keys() on empty store = %v , want empty slice", got)
	}
}

func TestSetGet_RoundTrip(t *testing.T) {
	store := NewStore()
	store.Set("hello", "world")

	if val, ok := store.Get("hello"); !ok || val != "world" {
		t.Errorf("Get() failed")
	}

	if val, ok := store.Get("missing"); ok || val != "" {
		t.Errorf("Get() failed")
	}
}
