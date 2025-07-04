package gq

import (
	"iter"
	"maps"
	"slices"
)

type Set[T comparable] map[T]struct{}

func NewSet[T comparable](items ...T) Set[T] {
	s := make(Set[T])
	for _, item := range items {
		s.Add(item)
	}
	return s
}

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

// Iter lets you loop through all items in the set.
// Uses Go's new iterator pattern.
func (s Set[T]) Iter() iter.Seq[T] {
	return func(yield func(T) bool) {
		for item := range s {
			if !yield(item) {
				return
			}
		}
	}
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

// Values returns a slice of all current items in the set.
// Order is random, like in regular Go maps.
func (s Set[T]) Values() []T {
	return slices.Collect(maps.Keys(s))
}

// Len returns the number of items in the set.
func (s Set[T]) Len() int {
	return len(s)
}
