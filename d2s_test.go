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

	fi, err := r.Stat()
	if err != nil {
		t.Fatalf("Unable to read file stats: %v", err)
	}

	err = ReadGame(&sg, r, fi.Size())
	if err != nil {
		t.Errorf("Error reading saved game: %v", err)
	}

	data := &sg.FileHeader
	if data.FileId != 0xAA55AA55 {
		t.Errorf("Bad value for File Identifier: 0x%X", data.FileId)
	}

	if data.FileVersion != 0x60 {
		t.Errorf("Bad value for File Version: 0x%X", data.FileVersion)
	}

	nameStr := string(data.CharName[:])
	expectedBuff := []byte{'S', 'i', 'l', 'l', 'y', 'n', 'e', 'c', 'r', 'o', 0, 0, 0, 0, 0, 0}
	expectedStr := string(expectedBuff[:])
	if nameStr != expectedStr {
		t.Errorf("Bad value for Character Name: %s", data.CharName)
	}

	chksum := data.Checksum
	err = sg.Checksum()
	if err != nil {
		t.Errorf("Error calculating checksum: %v", err)
	}

	if chksum != data.Checksum {
		t.Errorf("Bad value for Checksum: %v", data.Checksum)
	}
}
