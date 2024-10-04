package storage

import (
	"maps"

	"github.com/DistilledP/projects/simple_redis/internal/types"
)

type InMemory struct {
	store map[string]types.Value
}

func (m *InMemory) Add(key, value string) {
	matches := m.Find(key)
	if len(matches) == 1 {
		match := matches[key]
		match.Update(value)
		m.store[key] = match
	} else {
		m.store[key] = types.NewValue(value)
	}
}

func (m *InMemory) Del(keys ...string) int {
	matches := m.Find(keys...)
	for k := range matches {
		delete(m.store, k)
	}

	return len(matches)
}

func (m *InMemory) Find(keys ...string) map[string]types.Value {
	matches := map[string]types.Value{}
	for _, key := range keys {
		// Need some logic to handle wild card matches
		if v, found := m.store[key]; found {
			matches[key] = v
		}
	}

	return matches
}

func (m *InMemory) Indexes(filter string) []string {
	indexes := []string{}
	for k := range maps.Keys(m.store) {
		// Need some logic to wild card matches
		indexes = append(indexes, k)
	}

	return indexes
}

func NewInMemory() *InMemory {
	return &InMemory{
		store: make(map[string]types.Value),
	}
}
