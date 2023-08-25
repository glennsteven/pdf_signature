package signature

import "testing"

func TestLoadFile(t *testing.T) {
	file, err := LoadFile("file.txt")
	if err != nil {
		t.Fatalf("Expected: %v, but got %v", nil, err)
	}
	if string(file) != "Glenn Steven" {
		t.Fatalf("Expected: %s, Actual: %s", "Glenn Steven", string(file))
	}
}
