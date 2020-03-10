package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
)

type Encrypter interface {
	// Encrypt will take in a key and plaintext and return a hex representation
	// of the encrypted value.
	// This code is based on the standard library examples at:
	//   - https://golang.org/pkg/crypto/cipher/#NewCFBEncrypter
	Encrypt(plaintext string) (string, error)
}

type Decrypter interface {
	// Decrypt will take in a key and a cipherHex (hex representation of
	// the ciphertext) and decrypt it.
	// This code is based on the standard library examples at:
	//   - https://golang.org/pkg/crypto/cipher/#NewCFBDecrypter
	Decrypt(cipherHex string) (string, error)
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
	cipherText := make([]byte, aes.BlockSize+len(plaintext))
	iv := cipherText[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	stream := cipher.NewCFBEncrypter(c.block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], []byte(plaintext))

	return fmt.Sprintf("%x", cipherText), nil
}

func (c *client) Decrypt(cipherHex string) (string, error) {
	cipherText, err := hex.DecodeString(cipherHex)
	if err != nil {
		return "", err
	}

	if len(cipherText) < aes.BlockSize {
		return "", errors.New("encrypt: cipher too short")
	}
	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(c.block, iv)

	// XORKeyStream can work in-place if the two arguments are the same.
	stream.XORKeyStream(cipherText, cipherText)
	return string(cipherText), nil
}

func newCipherBlock(key string) (cipher.Block, error) {
	hasher := md5.New()
	fmt.Fprint(hasher, key)
	cipherKey := hasher.Sum(nil)
	return aes.NewCipher(cipherKey)
}
