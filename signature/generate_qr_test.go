package signature

import (
	"os"
	"testing"
)

func TestGenerateQr(t *testing.T) {
	qr, err := GenerateQr("glenn_steven.jpg")
	if err != nil {
		t.Fatalf("Expected: %v, Actual: %v", nil, err)
	}
	if len(qr) == 0 {
		t.Fatalf("Expected not zero value")
	}

	logo, err := GenerateLogo("logo.png")
	if err != nil {
		t.Fatalf("Expected: %v, Actual: %v", nil, err)
	}
	if len(logo) == 0 {
		t.Fatalf("Expected not zero value")
	}

	withLogo, err := QrCodeWithLogo(qr, logo)
	if err != nil {
		t.Fatalf("Expected: %v, Actual: %v", nil, err)
	}

	if len(withLogo) == 0 {
		t.Fatalf("New Qrcode Expected not zero value")
	}

	createNewQr, err := os.Create("RESULTQR.png")
	if err != nil {
		t.Fatalf("Expected: %v, Actual: %v", nil, err)
	}
	_, err = createNewQr.Write(withLogo)
	if err != nil {
		t.Fatalf("Expected: %v, Actual: %v", nil, err)
	}

}
