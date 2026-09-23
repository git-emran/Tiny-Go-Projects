package main

import "fmt"

type Store struct {
	data map[string]string
}

func (s *Store) Get(key string) (string, bool) {
	return "", false
}

func (s *Store) Set(key string, value string) {
}

func NewStore() *Store {
	return &Store{
		data: make(map[string]string),
	}
}

func main() {
	store := NewStore()
	store.Set("key1", "value")
	store.Get("key1")
	fmt.Println("Go key value store project")
}
