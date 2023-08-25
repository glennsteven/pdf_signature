package signature

import (
	"os"
	"testing"
)

func TestConvertFileImage(t *testing.T) {
	file, err := LoadFile("fotofahmi.png")
	if err != nil {
		t.Fatalf("Expected: %v, but got %v", nil, err)
	}

	mimetype, err := CheckMimetype(file)
	if err != nil {
		t.Fatalf("Expected: %v, but got %v", nil, err)
	}

	//if mimetype != "image/jpeg" {
	//	t.Fatalf("Expected: %s, Actual: %s", "image/jpeg", mimetype)
	//}
	//
	//jpg, err := ConvertFileJpg(file)
	//if err != nil {
	//	t.Fatalf("Expected: %v, but got %v", nil, err)
	//}
	//
	//if len(jpg) == 0 {
	//	t.Fatalf("Expected not zero")
	//}
	//
	//create, err := os.Create("output.pdf")
	//if err != nil {
	//	t.Fatalf("Expected: %v, but got %v", nil, err)
	//}
	//
	//_, err = create.Write(jpg)
	//if err != nil {
	//	t.Fatalf("Expected: %v, but got %v", nil, err)
	//}

	if mimetype != "image/png" {
		t.Fatalf("Expected: %s, Actual: %s", "image/png", mimetype)
	}

	png, err := ConvertFilePng(file)
	if err != nil {
		t.Fatalf("Expected: %v, but got %v", nil, err)
	}

	if len(png) == 0 {
		t.Fatalf("Expected not zero")
	}

	create, err := os.Create("output_png.pdf")
	if err != nil {
		t.Fatalf("Expected: %v, but got %v", nil, err)
	}

	_, err = create.Write(png)
	if err != nil {
		t.Fatalf("Expected: %v, but got %v", nil, err)
	}
}
