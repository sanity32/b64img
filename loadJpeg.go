package b64img

import (
	"encoding/base64"
	"os"
)

func Load(filename string) (r Image, err error) {
	b, err := os.ReadFile(filename)
	return Image(base64.StdEncoding.EncodeToString(b)), err
}
