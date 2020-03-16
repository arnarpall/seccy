package encrypt

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEncryptText(t *testing.T) {
	enc, err := NewClient("test-key")
	require.NoError(t, err)

	greeting := "hello"
	cipher, err := enc.Encrypt(greeting)
	require.NoError(t, err)
	assert.NotEmpty(t, cipher)

	decrypt, err := enc.Decrypt(cipher)
	require.NoError(t, err)
	assert.Equal(t, greeting, decrypt)
}

func TestEncryptWriter(t *testing.T) {
	enc, err := NewClient("test-key")
	require.NoError(t, err)

	encryptBuf := new(bytes.Buffer)
	w, err := enc.EncryptWriter(encryptBuf)
	require.NoError(t, err)
	assert.NotNil(t, w)

	greeting := "hello there"
	n, err := w.Write([]byte(greeting))
	assert.NoError(t, err)
	assert.Equal(t, len(greeting), n)

	r, err := enc.DecryptReader(encryptBuf)
	require.NoError(t, err)
	assert.NotNil(t, r)

	decrypted, err := ioutil.ReadAll(r)
	require.NoError(t, err)
	assert.Equal(t, greeting, string(decrypted))
}
