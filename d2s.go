// Package d2s provides functions for interfacing with Diablo II saved game
// file formats.  Current file formats supported are v1.10 - v1.13d.
package d2s

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"regexp"
)

const hOffset int64 = 767
const nameRegexp = "^[^-_][a-zA-Z]+[-_]?[a-zA-Z]+[^-_]$"

// Top level struct for interfacing with a saved game file
type SavedGame struct {
	header saveHeader
	buffer []byte
}

type saveHeader struct {
	FileId          uint32
	FileVersion     uint32
	FileSize        uint32
	Checksum        uint32
	ActiveArms      uint32
	CharName        [16]byte
	CharStatus      byte
	CharProgression byte
	Unk0            [2]byte
	CharClass       Class
	Unk1            [2]byte
	CharLvl         byte
	Unk2            uint32
	TimeStamp       uint32
	Unk3            uint32
	Hotkeys_0       uint32
	Hotkeys_1       uint32
	Hotkeys_2       uint32
	Hotkeys_3       uint32
	Hotkeys_4       uint32
	Hotkeys_5       uint32
	Hotkeys_6       uint32
	Hotkeys_7       uint32
	Hotkeys_8       uint32
	Hotkeys_9       uint32
	Hotkeys_10      uint32
	Hotkeys_11      uint32
	Hotkeys_12      uint32
	Hotkeys_13      uint32
	Hotkeys_14      uint32
	Hotkeys_16      uint32
	L_Mouse         uint32
	R_Mouse         uint32
	L_MouseAction   uint32
	R_MouseAction   uint32
	Unk4            [32]byte
	Difficulty      [3]byte
	MapID           uint32
	Unk5            uint16
	MercDead        uint16
	MercenaryID     uint32
	MercLangIndex   uint16
	MercenaryAttr   uint16
	MercenaryExp    uint32
	Unk6            [144]byte
	Woo             [4]byte
	Unk7            [6]byte

	// Quest data for each difficulty level
	QuestsNormal    quests
	QuestsNightmare quests
	QuestsHell      quests

	// Waypoint data for each difficulty level
	WS   [2]byte
	Unk8 [6]byte

	WPSNormal    waypoints
	WPSNightmare waypoints
	WPSHell      waypoints
	Unk9         [1]byte

	// NPC Introductions
	// TODO:  Implement NPC Introductions
	W4   [2]byte
	NPCS [49]byte

	// File size is dynamic past here

	// Character Stats
	GF [2]byte

	// TODO: Character Skills
	//			 Item List
	//			 Alive or Dead (Corpse?)
	//			 Mercenary Item List
	//			 Iron Golem Item
}

// New reads from an io.Reader size bytes and returns a new SavedGame.
// This method is used to take a stream of bytes (like from an io.File),
// and convert it into a SavedGame object that can be used to modify and
// rewrite the original file.
func New(r io.Reader, size int64) (sg *SavedGame, err error) {
	if size < 0 {
		return nil, fmt.Errorf("d2s.ReadGame error: invalid size")
	} else if size < hOffset {
		return nil, fmt.Errorf("d2s.ReadGame error: size too small")
	}

	sg = new(SavedGame)
	// Read static header struct
	err = binary.Read(r, binary.LittleEndian, &sg.header)
	if err != nil {
		return nil, err
	}

	// Read remaining into buffer
	sg.buffer = make([]byte, size-hOffset)
	_, err = r.Read(sg.buffer)
	if err != nil {
		return nil, err
	}

	return
}

func (sg *SavedGame) Read(p []byte) (n int, err error) {
	b := new(bytes.Buffer)
	err = binary.Write(b, binary.LittleEndian, sg.header)
	if err != nil {
		return 0, err
	}

	n = copy(p, b.Bytes())
	n += copy(p[hOffset:], sg.buffer)

	return
}

// Checksum performs the checksum algorithim on a SavedGame.
// The current checksum value must be set to 0 then the checksum
// algorithm is performed and the new value is set.  If an error was
// encountered the original checksum value will be replaced and Checksum
// will return 0
func (sg *SavedGame) Checksum() uint32 {
	p := &sg.header.Checksum
	c := *p
	*p = 0

	b := make([]byte, hOffset+int64(len(sg.buffer)))
	_, err := sg.Read(b)
	if err != nil {
		*p = c
		return 0
	}

	*p = checksum(b, 0)
	return *p
}

// checksum takes a byte slice, an initial checksum value and returns
// a new checksum value by adding the value of each byte sequentially.
// A bitwise left-shift (accounting for overflow) is performed on the
// checksum before the next byte is added.
func checksum(b []byte, chk uint32) uint32 {
	for _, c := range b {
		chk = ((chk << 1) | (chk >> 31)) + uint32(c)
	}
	return chk
}

// Name returns the name of the character as a string.
func (sg *SavedGame) Name() (name string) {
	b := sg.header.CharName
	i := bytes.IndexByte(b[:], 0)

	if i > 0 {
		name = string(b[:i])
	}
	return
}

// Sets CharName field, returns an error if name is not valid.
// Valid names must be 2-15 characters in length, and can only contain
// characters from the english alphabet with exception to one dash(-) or
// underscore(_) as long as it is not the first or last character.
func (sg *SavedGame) SetName(name string) error {
	l := len(name)
	if l < 2 || l > 15 {
		return fmt.Errorf("d2s.SetName name must be 2-15 characters long")
	}

	match, _ := regexp.MatchString(nameRegexp, name)
	if !match {
		return fmt.Errorf("d2s.SetName invalid name: %s", name)
	}

	for i := range sg.header.CharName {
		if i < l {
			sg.header.CharName[i] = name[i]
		} else {
			sg.header.CharName[i] = 0 // padding
		}
	}

	return nil
}
