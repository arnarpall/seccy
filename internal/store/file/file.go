package file

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"

	"github.com/arnarpall/seccy/internal"
	"github.com/arnarpall/seccy/internal/encrypt"
)

type fileStore struct {
	enc  encrypt.EncrypterDecrypter
	path string
	mu   sync.Mutex
}

func NewFileStore(enc encrypt.EncrypterDecrypter, path string) (internal.Store, error) {
	return &fileStore{
		path: path,
		enc:  enc,
	}, nil
}

func (f *fileStore) Set(key, val string) error {
	f.mu.Lock()
	defer  f.mu.Unlock()

	entries, err := f.load()
	if err != nil {
		return err
	}

	entries[key] = val
	if err := f.saveEntries(entries); err != nil {
		return err
	}

	return nil
}

func (f *fileStore) Get(key string) (string, error) {
	f.mu.Lock()
	defer  f.mu.Unlock()

	entries, err := f.load()
	if err != nil {
		return "", err
	}

	entry, ok := entries[key]
	if !ok {
		return "", fmt.Errorf("no entry for key %s exists", key)
	}

	return entry, nil
}

func (f *fileStore) load() (map[string]string, error) {
	file, err := os.Open(f.path)
	entries := make(map[string]string)
	if err != nil {
		return entries, nil
	}
	defer file.Close()

	var sb strings.Builder
	_, err = io.Copy(&sb, file)
	if err != nil {
		return entries, err
	}

	decryptedJSON, err := f.enc.Decrypt(sb.String())
	if err != nil {
		return entries, err
	}

	r := strings.NewReader(decryptedJSON)
	dec := json.NewDecoder(r)
	err = dec.Decode(&entries)
	if err != nil {
		return entries, err
	}

	return entries, nil
}

func (f *fileStore) saveEntries(entries map[string]string) error {
	var sb strings.Builder
	enc := json.NewEncoder(&sb)

	err := enc.Encode(entries)
	if err != nil {
		return err
	}

	encryptedJSON, err := f.enc.Encrypt(sb.String())
	if err != nil  {
		return err
	}

	file, err := os.OpenFile(f.path, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil  {
		return err
	}
	defer file.Close()

	_, err = fmt.Fprint(file, encryptedJSON)
	if err != nil  {
		return err
	}

	return nil
}
