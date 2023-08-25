package signature

import (
	"bytes"
	"io"
	"os"
)

func LoadFile(input string) ([]byte, error) {
	result := new(bytes.Buffer)
	open, err := os.Open(input)
	if err != nil {
		return nil, err
	}

	_, err = io.Copy(result, open)
	if err != nil {
		return nil, err
	}

	return result.Bytes(), nil
}
