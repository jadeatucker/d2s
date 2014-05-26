package d2s

import (
	"encoding/binary"
	"fmt"
	"io"
)

type SavedGame struct {
	FileId          uint32
	FileVersion     uint32
	FileSize        uint32
	Checksum        uint32
	ActiveArms      uint32
	CharName        [16]byte
	CharStatus      byte
	CharProgression byte
	_               [2]byte
	CharClass       byte
	_               [2]byte
	CharLvl         byte
	_               uint32
	TimeStamp       uint32
	_               uint32
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
	_               [32]byte
	Difficulty      [3]byte
	MapID           uint32
	_               uint16
	MercDead        uint16
	MercenaryID     uint32
	MercLangIndex   uint16
	MercenaryAttr   uint16
	MercenaryExp    uint32
	_               [144]byte
	Woo             [4]byte
}

func ReadGame(sg *SavedGame, r io.Reader) {
	err := binary.Read(r, binary.LittleEndian, sg)

	if err != nil {
		fmt.Println("binary.Read failed:", err)
	}
}
