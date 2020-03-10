package encrypt

type noOpClient struct {
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
