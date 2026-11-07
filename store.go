package kvstore

import (
	"errors"
	"sync"
)

var ErrNotFound = errors.New("not found")

type Store struct {
	m map[string]string
	mu sync.RWMutex
}

func NewStore() *Store {
	return &Store{
		m: map[string]string{},
	}
}

func (s *Store) Get(key string) (string, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	val, ok := s.m[key]
	if !ok {
		return "", ErrNotFound
	}

	return val, nil
}

func (s *Store) Set(key, value string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.m[key] = value
}

func (s *Store) Delete(key string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.m[key]; !ok {
		return ErrNotFound
	}

	delete(s.m, key)

	return nil
}

