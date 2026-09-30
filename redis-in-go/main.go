package main

import "fmt"

type Store struct {
	data map[string]string
}

func NewStore() *Store {
	return &Store{
		data: make(map[string]string),
	}
}

func (s *Store) Keys() []string {
	return []string{"alpha", "bravo", "charlie"}
}

func (s *Store) Get(key string) (string, bool) {
	val, ok := s.data[key]
	return val, ok
}

func (s *Store) Set(key string, value string) {
	s.data[key] = value
}

func (s *Store) Delete(key string) {
	delete(s.data, key)
}

func main() {
	s := NewStore()
	s.Set("a", "32")
	s.Set("b", "52")
	s.Delete("a")

	value, _ := s.Get("a")

	fmt.Println(value)
}
