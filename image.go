package b64img

import (
	"bytes"
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"strings"
)

type Image string

func (b Image) Hash() Hash {
	b.Clean()
	return Hash(fmt.Sprintf("%x", md5.Sum([]byte(b))))
}

func (b Image) Match(h Hash) bool {
	return b.Hash() == h
}

func (b Image) String() string {
	return string(b)
}

func (b Image) hasPrefix(p string) bool {
	return strings.HasPrefix(b.String(), p)
}

func (b Image) RemovePrefix(p string) Image {
	if b.hasPrefix(p) {
		b = Image(strings.TrimPrefix(b.String(), p))
	}
	return b
}

func (b *Image) Clean() {
	*b = b.RemovePrefix(PREFIX_B64_JPG)
	*b = b.RemovePrefix(PREFIX_B64_PNG)
}

func (b Image) ToPng() (image.Image, error) {
	data, err := base64.StdEncoding.DecodeString(b.String())
	if err != nil {
		return nil, err
	}
	bb := bytes.NewReader(data)
	return png.Decode(bb)
}

func (b Image) SavePng(filename string) error {
	b.Clean()
	img, err := b.ToPng()
	if err != nil {
		return err
	}

	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	return png.Encode(f, img)
}

func (b Image) ToJpg() (image.Image, error) {
	data, err := base64.StdEncoding.DecodeString(b.String())
	if err != nil {
		return nil, err
	}
	bb := bytes.NewReader(data)
	return jpeg.Decode(bb)
}

func (b Image) SaveJpeg(filename string, q ...int) error {
	b.Clean()
	img, err := b.ToJpg()
	if err != nil {
		return err
	}
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	var opts *jpeg.Options
	if len(q) > 0 {
		opts = &jpeg.Options{Quality: q[0]}
	}

	return jpeg.Encode(f, img, opts)
}

func (p Image) WithJpgPrefix() Image {
	if p.hasPrefix(PREFIX_B64_JPG) {
		return p
	}
	return PREFIX_B64_JPG + p
}
