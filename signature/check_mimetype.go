package signature

import "net/http"

func CheckMimetype(input []byte) (string, error) {
	resMimeType := http.DetectContentType(input)
	return resMimeType, nil
}
