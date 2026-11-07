package kvstore

import (
	"errors"
)

var ErrNotFound = errors.New("not found")

type Store struct {
	m map[string]string
}

func NewStore() *Store {
	return &Store{
		m: map[string]string{},
	}
}

func (s *Store) Get(key string) (string, error) {
	val, ok := s.m[key]
	if !ok {
		return "", ErrNotFound
	}

	return val, nil
}

func (s *Store) Set(key, value string) {
	s.m[key] = value
}

func (s *Store) Delete(key string) error {
	if _, ok := s.m[key]; !ok {
		return ErrNotFound
	}

	delete(s.m, key)

	return nil
}
