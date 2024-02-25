package b64img

import (
	"errors"
	"os"
	"path"
)

var ErrDirHasNoHashImg = errors.New("dir has no hashed image to recover")

type HashDir string

func (hd HashDir) Exist() bool {
	stat, err := os.Stat(string(hd))
	return err == nil && stat.IsDir()
}

func (hd HashDir) Create() error {
	return os.MkdirAll(string(hd), 0644)
}

func (hd HashDir) Read(h Hash) (im Image, err error) {
	if !hd.Exist() {
		return im, ErrDirHasNoHashImg
	}
	f := path.Join(string(hd), string(h))
	bb, err := os.ReadFile(f)
	return Image(bb), err
}

func (hd HashDir) Write(im Image) error {
	if !hd.Exist() {
		return ErrDirHasNoHashImg
	}
	f := path.Join(string(hd), string(im.Hash()))
	return os.WriteFile(f, []byte(im), 0644)
}
