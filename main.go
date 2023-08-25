package main

import (
	"fmt"
	"github.com/google/uuid"
	"os"
	"pdf_signature/signature"
)

func main() {
	stamp, err := PdfStamp("./signature/testing_seal.pdf", 7)
	if err != nil {
		panic(err)
	}

	create, err := os.Create("result_main.pdf")
	if err != nil {
		panic(err)
	}

	_, err = create.Write(stamp)
	if err != nil {
		panic(err)
	}
}

func PdfStamp(input string, pos int) ([]byte, error) {
	uid := uuid.NewString()
	qrCodeValue := fmt.Sprintf("https://privy.id/verify/%s", uid)
	file, err := signature.LoadFile(input)
	if err != nil {
		return nil, err
	}

	mimetype, err := signature.CheckMimetype(file)
	if err != nil {
		return nil, err
	}

	convert, err := signature.Convert(file, mimetype)
	if err != nil {
		return nil, err
	}

	qr, err := signature.GenerateQr(qrCodeValue)
	if err != nil {
		return nil, err
	}

	logo := "logo.png"
	logoQr, err := signature.GenerateLogo(logo)
	if err != nil {
		return nil, err
	}

	fmt.Println("LOGO", logo)

	qrWithLogo, err := signature.QrCodeWithLogo(qr, logoQr)
	if err != nil {
		return nil, err
	}

	stamp, err := signature.Stamp(convert, qrWithLogo, pos)
	if err != nil {
		return nil, err
	}

	return stamp, nil
}
