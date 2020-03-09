package secret

type fileVault struct {
	enc string
	path string
}

func FileVault(enc, path string) (Vault, error)  {
	return &fileVault{
		enc:  enc,
		path: path,
	}, nil
}

func (f *fileVault) Set(key, val string) error {
	panic("implement me")
}

func (f *fileVault) Get(key string) (string, error) {
	panic("implement me")
}
