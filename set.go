package gq

import (
	"iter"
	"maps"
)

type Set[T comparable] map[T]struct{}

// Add puts a new item in the set.
// If item already exists, does nothing.
func (s Set[T]) Add(item T) {
	s[item] = struct{}{}
}

// Contains checks if an item exists in the set.
func (s Set[T]) Contains(item T) bool {
	_, exists := s[item]
	return exists
}

// Delete removes an item from the set.
// Does nothing if item doesn't exist.
func (s Set[T]) Delete(item T) {
	delete(s, item)
}

func (s Set[T]) Iterate(f func(T) bool) {
	for item := range s {
		if !f(item) {
			return
		}
	}
}

// Clear removes all items from the set at once.
func (s Set[T]) Clear() {
	clear(s)
}

// Values returns an iterator of all current items in the set.
// Order is random, like in regular Go maps.
func (s Set[T]) Values() iter.Seq[T] {
	return maps.Keys(s)
}

// Len returns the number of items in the set.
func (s Set[T]) Len() int {
	return len(s)
}
