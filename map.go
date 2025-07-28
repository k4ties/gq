package gq

import (
	"iter"
	"maps"
)

type Map[K comparable, V any] map[K]V

// Len tells how many items are in the map right now.
func (m Map[K, V]) Len() int {
	return len(m)
}

// Keys gives a list of all current keys in the map.
// The order of keys is random, like in regular Go maps.
func (m Map[K, V]) Keys() iter.Seq[K] {
	return maps.Keys(m)
}

// Values gives a list of all current values in the map.
// The order of values is random, like in regular Go maps.
func (m Map[K, V]) Values() iter.Seq[V] {
	return maps.Values(m)
}

// Put adds a new key-value pair only if the key doesn't exist yet.
// Returns true if the key was added, false if it already existed.
func (m Map[K, V]) Put(key K, value V) bool {
	if _, exists := m[key]; exists {
		return false
	}
	m.Set(key, value)
	return true
}

// Get looks up a key and returns its value if found.
// The second return value tells if the key was found.
func (m Map[K, V]) Get(key K) (V, bool) {
	value, found := m[key]
	return value, found
}

// Set updates or adds a key-value pair.
// This will overwrite any existing value for the key.
func (m Map[K, V]) Set(key K, value V) {
	m[key] = value
}

// Iterate goes through all items and lets you do something with each one.
// Stop early by returning false from your function.
func (m Map[K, V]) Iterate(action func(key K, value V) bool) {
	for k, v := range m {
		if !action(k, v) {
			return
		}
	}
}

// Iter gives a way to loop through items using Go's new iterators.
// It safely copies items first so you can take your time looping.
func (m Map[K, V]) Iter() iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for k, v := range m {
			if !yield(k, v) {
				return
			}
		}
	}
}

// Delete removes one item from the map by its key.
// Does nothing if the key doesn't exist.
func (m Map[K, V]) Delete(key K) {
	delete(m, key)
}

// Clear removes all items from the map at once.
func (m Map[K, V]) Clear() {
	clear(m)
}
