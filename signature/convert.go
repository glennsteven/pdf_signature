package signature

import (
	"bytes"
	"github.com/pdfcpu/pdfcpu/pkg/api"
	"io"
)

func ConvertFileJpg(input []byte) ([]byte, error) {
	resPdf := new(bytes.Buffer)
	err := api.ImportImages(nil, resPdf, []io.Reader{bytes.NewBuffer(input)}, nil, nil)
	if err != nil {
		return nil, err
	}

	return resPdf.Bytes(), nil
}

func ConvertFilePng(input []byte) ([]byte, error) {
	resPdf := new(bytes.Buffer)
	err := api.ImportImages(nil, resPdf, []io.Reader{bytes.NewBuffer(input)}, nil, nil)
	if err != nil {
		return nil, err
	}
	return resPdf.Bytes(), nil
}

func Convert(input []byte, mimetype string) ([]byte, error) {
	switch mimetype {
	case "image/jpeg":
		return ConvertFileJpg(input)
	case "image/png":
		return ConvertFilePng(input)
	case "application/vnd.openxmlformats-officedocument.wordprocessingml.document":
		return nil, nil
	default:
		return input, nil
	}
}
