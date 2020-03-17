package file

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"

	"github.com/arnarpall/seccy/internal/encrypt"
	"github.com/arnarpall/seccy/internal/store"
)

type fileStore struct {
	enc  encrypt.EncrypterDecrypter
	path string
	mu   sync.Mutex
}

func NewFileStore(enc encrypt.EncrypterDecrypter, path string) (store.Store, error) {
	return &fileStore{
		path: path,
		enc:  enc,
	}, nil
}

func (fs *fileStore) Set(key, val string) error {
	fs.mu.Lock()
	defer fs.mu.Unlock()

	entries, err := fs.load()
	if err != nil {
		return err
	}

	entries[key] = val
	if err := fs.saveEntries(entries); err != nil {
		return err
	}

	return nil
}

func (fs *fileStore) Get(key string) (string, error) {
	fs.mu.Lock()
	defer fs.mu.Unlock()

	entries, err := fs.load()
	if err != nil {
		return "", err
	}

	entry, ok := entries[key]
	if !ok {
		return "", fmt.Errorf("no entry for key %s exists, %w", key, store.ErrKeyNotFound)
	}

	return entry, nil
}

func (fs *fileStore) ListKeys() ([]string, error) {
	entries, err := fs.load()
	if err != nil {
		return []string{}, err
	}

	keys := make([]string, 0, len(entries))
	for k := range entries {
		keys = append(keys, k)
	}

	return keys, nil
}

func (fs *fileStore) load() (map[string]string, error) {
	file, err := os.Open(fs.path)
	entries := make(map[string]string)
	if err != nil {
		return entries, nil
	}
	defer file.Close()

	r, err := fs.enc.DecryptReader(file)
	if err != nil {
		return nil, err
	}

	dec := json.NewDecoder(r)
	err = dec.Decode(&entries)
	if err != nil {
		return entries, err
	}

	return entries, nil
}

func (fs *fileStore) saveEntries(entries map[string]string) error {
	f, err := os.OpenFile(fs.path, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	defer f.Close()

	w, err := fs.enc.EncryptWriter(f)

	if err != nil {
		return err
	}

	enc := json.NewEncoder(w)
	return enc.Encode(entries)
}
