package signature

import "testing"

func TestCheckMimetype(t *testing.T) {
	t.Run("load file txt", func(t *testing.T) {
		file, err := LoadFile("file.txt")
		if err != nil {
			t.Fatalf("Expected: %v, but got %v", nil, err)
		}
		if string(file) != "Glenn Steven" {
			t.Fatalf("Expected: %s, Actual: %s", "Glenn Steven", string(file))
		}

		mimetype, err := CheckMimetype(file)
		if err != nil {
			t.Fatalf("Expected: %v, but got %v", nil, err)
		}

		if mimetype != "text/plain; charset=utf-8" {
			t.Fatalf("Expected: %s, Actual: %s", "text/plain", mimetype)
		}
	})

	t.Run("load file pdf", func(t *testing.T) {
		file, err := LoadFile("testing_seal.pdf")
		if err != nil {
			t.Fatalf("Expected: %v, but got %v", nil, err)
		}

		mimetype, err := CheckMimetype(file)
		if err != nil {
			t.Fatalf("Expected: %v, but got %v", nil, err)
		}

		if mimetype != "application/pdf" {
			t.Fatalf("Expected: %s, Actual: %s", "application/pdf", mimetype)
		}
	})

	t.Run("load file image", func(t *testing.T) {
		file, err := LoadFile("glenn_steven.jpg")
		if err != nil {
			t.Fatalf("Expected: %v, but got %v", nil, err)
		}

		mimetype, err := CheckMimetype(file)
		if err != nil {
			t.Fatalf("Expected: %v, but got %v", nil, err)
		}

		if mimetype != "image/jpeg" {
			t.Fatalf("Expected: %s, Actual: %s", "image/jpeg", mimetype)
		}
	})
}
