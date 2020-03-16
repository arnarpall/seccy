package encrypt

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
)

type Encrypter interface {
	// Encrypt will take in plaintext and return a hex representation
	// of the encrypted value.
	// This code is based on the standard library examples at:
	//   - https://golang.org/pkg/crypto/cipher/#NewCFBEncrypter
	Encrypt(plaintext string) (string, error)

	// Encrypt will take in a writer and a new writer that will encrypt on writes.
	// This code is based on the standard library examples at:
	//   - https://golang.org/pkg/crypto/cipher/#NewCFBEncrypter
	EncryptWriter(w io.Writer) (io.Writer, error)
}

type Decrypter interface {
	// Decrypt will take in a cipherString (encrypted piece of string) and decrypt it.
	// This code is based on the standard library examples at:
	//   - https://golang.org/pkg/crypto/cipher/#NewCFBDecrypter
	Decrypt(cipherHex string) (string, error)

	// DecryptReader will take in a reader and wrap it in a reader that will decrypt on reads
	// This code is based on the standard library examples at:
	//   - https://golang.org/pkg/crypto/cipher/#NewCFBDecrypter
	DecryptReader(r io.Reader) (io.Reader, error)
}

type EncrypterDecrypter interface {
	Encrypter
	Decrypter
}

type client struct {
	block cipher.Block
}

func NewClient(key string) (EncrypterDecrypter, error) {
	block, err := newCipherBlock(key)
	if err != nil {
		return nil, err
	}

	return &client{
		block: block,
	}, nil
}

func (c *client) Encrypt(plaintext string) (string, error) {
	s := new(bytes.Buffer)
	stream, err := c.EncryptWriter(s)
	if err != nil {
		return "", err
	}

	n, err := io.Copy(stream, bytes.NewBufferString(plaintext))
	if n != int64(len(plaintext)) || err != nil {
		return "", fmt.Errorf("encrypt: unable to encrypt plaintext %s, %w", plaintext, err)
	}

	return s.String(), err
}

func (c *client) EncryptWriter(w io.Writer) (io.Writer, error) {
	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}
	stream := cipher.NewCFBEncrypter(c.block, iv)
	n, err := w.Write(iv)
	if n != len(iv) || err != nil {
		return nil, errors.New("encrypt: unable to write full iv to writer")
	}
	return &cipher.StreamWriter{S: stream, W: w}, nil
}

func (c *client) Decrypt(cipherText string) (string, error) {
	stream, err := c.DecryptReader(bytes.NewBufferString(cipherText))
	if err != nil {
		return "", err
	}

	decrypted, err := ioutil.ReadAll(stream)
	if err != nil {
		return "", fmt.Errorf("decrypt: unable to decrypt cipherText %s, %w", cipherText, err)
	}
	return string(decrypted), nil
}

func (c *client) DecryptReader(r io.Reader) (io.Reader, error) {
	iv := make([]byte, aes.BlockSize)
	n, err := r.Read(iv)
	if n < len(iv) || err != nil {
		return nil, errors.New("decrypt: unable to read the full iv")
	}
	stream := cipher.NewCFBDecrypter(c.block, iv)
	return &cipher.StreamReader{S: stream, R: r}, nil
}

func newCipherBlock(key string) (cipher.Block, error) {
	h := md5.New()
	_, err := fmt.Fprint(h, key)
	if err != nil {
		return nil, err
	}
	cipherKey := h.Sum(nil)
	return aes.NewCipher(cipherKey)
}
