package flyweight

import (
	"io/ioutil"
	"path/filepath"
)

type ImageFlyweight struct {
	Path string
	Data []byte
}

// NewImageFlyweight
// p is abslote path
func NewImageFlyweight(p string) (*ImageFlyweight, error) {
	data, err := ioutil.ReadFile(p)
	if err != nil {
		return nil, err
	}
	return &ImageFlyweight{
		Path: p,
		Data: data,
	}, nil
}

type ImageFlyweightFactory struct {
	Cache map[string]*ImageFlyweight // need to init before use
}

func (i *ImageFlyweightFactory) GetImage(p string) (*ImageFlyweight, error) {
	var p_ string
	var err error
	if filepath.IsAbs(p) {
		p_, err = filepath.Abs(p)
		if err != nil {
			return nil, err
		}
	} else {
		p_ = p
	}
	if image, ok := i.Cache[p_]; ok {
		return image, nil
	}
	image, err := NewImageFlyweight(p_)
	if err != nil {
		return nil, err
	}
	i.Cache[p_] = image
	return image, err
}
