package encrypt

import "io"

type noOpClient struct {
}

func (n noOpClient) DecryptReader(r io.Reader) (io.Reader, error) {
	return r, nil
}

func (n noOpClient) EncryptWriter(w io.Writer) (io.Writer, error) {
	return w, nil
}

func (n noOpClient) Encrypt(plaintext string) (string, error) {
	return plaintext, nil
}

func (n noOpClient) Decrypt(cipherHex string) (string, error) {
	return cipherHex, nil
}

func NoOp(key string) (EncrypterDecrypter, error) {
	return new(noOpClient), nil
}
