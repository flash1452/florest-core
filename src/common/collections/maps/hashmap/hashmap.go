// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package hashmap implements a map backed by a hash table.
//
// Elements are unordered in the map.
//
// Structure is not thread safe.
//
// Reference: http://en.wikipedia.org/wiki/Associative_array
package hashmap

import (
	"fmt"
)

// Map holds the elements in go's native map
type Map struct {
	m map[interface{}]interface{}
}

// New instantiates a hash map.
func New() *Map {
	return &Map{m: make(map[interface{}]interface{})}
}

// Put inserts element into the map.
func (m *Map) Put(key interface{}, value interface{}) {
	m.m[key] = value
}

// Get searches the element in the map by key and returns its value or nil if key is not found in map.
// Second return parameter is true if key was found, otherwise false.
func (m *Map) Get(key interface{}) (value interface{}, found bool) {
	value, found = m.m[key]
	return
}

// Remove removes the element from the map by key.
func (m *Map) Remove(key interface{}) {
	delete(m.m, key)
}

// Empty returns true if map does not contain any elements
func (m *Map) IsEmpty() bool {
	return m.Size() == 0
}

// Size returns number of elements in the map.
func (m *Map) Size() int {
	return len(m.m)
}

// Keys returns all keys (random order).
func (m *Map) Keys() []interface{} {
	keys := make([]interface{}, m.Size())
	for key, _ := range m.m {
		keys = append(keys, key)
	}
	return keys
}

// Values returns all values (random order).
func (m *Map) Values() []interface{} {
	values := make([]interface{}, m.Size())
	for _, value := range m.m {
		values = append(values, value)
	}
	return values
}

func (m *Map) Contains(keys ...interface{}) bool {
	for _, key := range keys {
		_, found := m.m[key]
		if !found {
			return false
		}
	}
	return true
}

// Clear removes all elements from the map.
func (m *Map) Clear() {
	m.m = make(map[interface{}]interface{})
}

// String returns a string representation of container
func (m *Map) String() string {
	str := "HashMap\n"
	str += fmt.Sprintf("%v", m.m)
	return str
}