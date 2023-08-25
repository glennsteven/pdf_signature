package signature

import (
	"os"
	"testing"
)

func TestStamp(t *testing.T) {
	file, err := LoadFile("testing_seal.pdf")
	if err != nil {
		t.Fatalf("Expected: %v, but got %v", nil, err)
	}
	qr, err := GenerateQr("glenn_steven.png")
	if err != nil {
		t.Fatalf("Expected: %v, but got %v", nil, err)
	}
	stamp, err := Stamp(file, qr, 7)
	if err != nil {
		t.Fatalf("Expected: %v, but got %v", nil, err)
	}
	if len(stamp) == 0 {
		t.Fatalf("Expected not zero value")
	}

	create, err := os.Create("test_results.pdf")
	if err != nil {
		t.Fatalf("Expected: %v, but got %v", nil, err)
	}

	_, err = create.Write(stamp)
	if err != nil {
		t.Fatalf("Expected: %v, but got %v", nil, err)
	}
}
