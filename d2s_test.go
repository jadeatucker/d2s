package d2s

import (
	"os"
	"testing"
)

var (
	sg     *SavedGame
	f_size int64
)

const (
	e_checksum uint32 = 1436947527
	e_name            = "Sillynecro"
)

func TestReadGame(t *testing.T) {
	r, err := os.Open("testdata/Sillynecro.d2s")
	if err != nil {
		t.Fatalf("Unable to open file for reading: %v\n", err)
	}

	fi, err := r.Stat()
	if err != nil {
		t.Fatalf("Unable to read file stats: %v\n", err)
	} else {
		f_size = fi.Size()
	}

	sg, err = New(r, f_size)
	if err != nil {
		t.Fatalf("Error reading saved game: %v\n", err)
	}
}

func TestRead(t *testing.T) {
	b := make([]byte, f_size)

	n, err := sg.Read(b)
	if err != nil {
		t.Fatal(err)
	} else if int64(n) != f_size {
		t.Fatalf("Unexpected number of bytes read: %d\n", n)
	}
}

func TestChecksum(t *testing.T) {
	c := sg.Checksum()

	if c != e_checksum {
		t.Fatalf("Bad value for checksum: %v\n", c)
	}
}

func TestName(t *testing.T) {
	name := sg.Name()
	if name != e_name {
		t.Fatalf("Bad value for name: %v\n", name)
	}
}

func TestSetName(t *testing.T) {
	var err error

	validNames := []string{"Testname", "Test-Name", "Test_name"}
	for _, n := range validNames {
		err = sg.SetName(n)
		if err != nil {
			t.Error(err)
		}
	}

	badNames := []string{"Test-_Name", "Testname-", "_Testname", "", "a"}
	for _, n := range badNames {
		err = sg.SetName(n)
		if err == nil {
			t.Error(err)
		}
	}

	// Character name should be set to last valid name
	name := sg.Name()
	if name != validNames[len(validNames)-1] {
		t.Fatalf("Bad value for name: %v\n", name)
	}
}
