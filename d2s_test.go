package d2s

import (
	"os"
	"testing"
)

func TestSaveGame(t *testing.T) {
	var sg SavedGame
	r, err := os.Open("testdata/Sillynecro.d2s")

	if err != nil {
		t.Fatalf("Unable to open file for reading: %v", err)
	}

	ReadGame(&sg, r)

	if sg.FileId != 65493 {
		t.Fatalf("Bad value for FieldId: %v", sg.FileId)
	}
}
