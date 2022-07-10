package decorator

import (
	"encoding/base64"
	"io/ioutil"
)

type Source interface {
	WriteDate(p string, data []byte) error
	ReadData(p string) ([]byte, error)
}

type FileSource struct {
}

func (s *FileSource) WriteDate(p string, data []byte) error {
	return ioutil.WriteFile(p, data, 0666)
}

func (s *FileSource) ReadData(p string) ([]byte, error) {
	return ioutil.ReadFile(p)
}

type EncryptionDecorator struct {
	Wrapper Source
}

func (s *EncryptionDecorator) WriteDate(p string, data []byte) error {
	var out string
	out = base64.StdEncoding.EncodeToString(data)
	return s.Wrapper.WriteDate(p, []byte(out))
}

func (s *EncryptionDecorator) ReadData(p string) ([]byte, error) {
	data, err := s.Wrapper.ReadData(p)
	if err != nil {
		return nil, err
	}
	var out []byte
	_, err = base64.StdEncoding.DecodeString(string(data))
	return out, err
}
