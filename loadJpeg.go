package b64img

import (
	"encoding/base64"
	"os"
)

func Load(filename string) (r Image, err error) {
	bb, err := os.ReadFile(filename)
	return Image(base64.StdEncoding.EncodeToString(bb)), err
}
