package secret

import (
	"fmt"
	"sync"
)

type memoryVault struct {
	store map[string]string
	mu    *sync.Mutex
}

func (m *memoryVault) Set(key, val string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.store[key] = val
	return nil
}

func (m *memoryVault) Get(key string) (string, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	v, ok := m.store[key]
	if !ok {
		return "", fmt.Errorf("vaule for key %s does not exist", key)
	}

	return v, nil
}

func InMemory() Vault {
	return &memoryVault{
		store: make(map[string]string),
		mu:    new(sync.Mutex),
	}
}
