package memory

import (
	"fmt"
	"sync"

	"github.com/arnarpall/seccy/internal/store"
)

type memoryStore struct {
	store map[string]string
	mu    *sync.Mutex
}

func (m *memoryStore) ListKeys() ([]string, error) {
	panic("implement me")
}

func (m *memoryStore) Set(key, val string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.store[key] = val
	return nil
}

func (m *memoryStore) Get(key string) (string, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	v, ok := m.store[key]
	if !ok {
		return "", fmt.Errorf("value for key %s does not exist", key)
	}

	return v, nil
}

func InMemory() store.Store {
	return &memoryStore{
		store: make(map[string]string),
		mu:    new(sync.Mutex),
	}
}
