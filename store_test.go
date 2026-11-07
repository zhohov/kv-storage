package kvstore

import (
	"errors"
	"testing"
)

func TestGetMissingKey(t *testing.T) {
	store := NewStore()

	_, err := store.Get("test")
	if !errors.Is(err, ErrNotFound) {
		t.Fatalf("expected %v, got %v", ErrNotFound, err)
	}
}

func TestSetKey(t *testing.T) {
	store := NewStore()
	k, v := "key", "value"

	store.Set(k, v)

	got, err := store.Get(k)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if got != v {
		t.Errorf("expected %s, got %s", v, got)
	}
}

func TestOverwriteKey(t *testing.T) {
	store := NewStore()
	k, v1, v2 := "key", "value 1", "value 2"

	store.Set(k, v1)
	store.Set(k, v2)

	got, err := store.Get(k)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if got != v2 {
		t.Errorf("expected %s, got %s", v2, got)
	}
}

func TestDeleteKey(t *testing.T) {
	store := NewStore()
	k, v := "key", "value"

	store.Set(k, v)

	if err := store.Delete(k); err != nil {
		t.Fatalf("unexpected delete error: %v", err)
	}

	_, err := store.Get(k)
	if !errors.Is(err, ErrNotFound) {
		t.Fatalf("expected %v, got %v", ErrNotFound, err)
	}
}

func TestDeleteMissingKey(t *testing.T) {
	store := NewStore()

	err := store.Delete("key")
	if !errors.Is(err, ErrNotFound) {
		t.Fatalf("expected %v, got %v", ErrNotFound, err)
	}
}
