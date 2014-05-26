package d2s

import (
	"os"
	"testing"
)

func TestReadGame(t *testing.T) {
	var sg SavedGame
	r, err := os.Open("testdata/Sillynecro.d2s")

	if err != nil {
		t.Fatalf("Unable to open file for reading: %v", err)
	}

	ReadGame(&sg, r)

	if sg.FileId != 0xAA55AA55 {
		t.Errorf("Bad value for File Identifier: 0x%X", sg.FileId)
	}

	if sg.FileVersion != 0x60 {
		t.Errorf("Bad value for File Version: 0x%X", sg.FileVersion)
	}

	nameStr := string(sg.CharName[:])
	expectedBuff := []byte{'S', 'i', 'l', 'l', 'y', 'n', 'e', 'c', 'r', 'o', 0, 0, 0, 0, 0, 0}
	expectedStr := string(expectedBuff[:])
	if nameStr != expectedStr {
		t.Errorf("Bad value for Character Name: %s", sg.CharName)
	}
}
