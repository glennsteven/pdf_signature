package signature

import (
	"bytes"
	"github.com/nfnt/resize"
	"github.com/skip2/go-qrcode"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"log"
	"os"
)

func GenerateQr(input string) ([]byte, error) {
	encode, err := qrcode.Encode(input, qrcode.Medium, 125)
	if err != nil {
		return nil, err
	}

	return encode, nil
}

func GenerateLogo(input string) ([]byte, error) {
	logoFile, err := os.ReadFile(input)
	if err != nil {
		return nil, err
	}

	return logoFile, nil
}

func QrCodeWithLogo(qrCodeFile, logoFile []byte) ([]byte, error) {
	result := new(bytes.Buffer)
	readerLogo := bytes.NewReader(logoFile)

	decodeLogo, err := png.Decode(readerLogo)
	if err != nil {
		log.Fatalf("failed to decode: %s", err)
	}
	newLogo := resize.Resize(50, 50, decodeLogo, resize.Lanczos2)

	readerQr := bytes.NewReader(qrCodeFile)

	decodeQrCode, err := png.Decode(readerQr)
	if err != nil {
		log.Fatalf("failed to decode: %s", err)
	}
	newQrcode := resize.Resize(215, 215, decodeQrCode, resize.Lanczos2)

	offset := image.Pt(85, 85)
	sizeQrcode := newQrcode.Bounds()
	newQRCode := image.NewRGBA(sizeQrcode)
	draw.Draw(newQRCode, sizeQrcode, newQrcode, image.ZP, draw.Src)
	draw.Draw(newQRCode, newQRCode.Bounds().Add(offset), newLogo, image.ZP, draw.Over)

	err = jpeg.Encode(result, newQRCode, &jpeg.Options{jpeg.DefaultQuality})
	if err != nil {
		return nil, err
	}

	return result.Bytes(), nil
}
