package d2s

import (
	"encoding/binary"
	"io"
)

type SavedGame struct {
	FileId          uint16
	FileVersion     uint16
	FileSize        uint16
	Checksum        uint16
	ActiveArms      uint16
	CharName        [16]string
	CharStatus      byte
	CharProgression byte
	unknown0        [2]byte
	CharClass       byte
	unknown1        [2]byte
	CharLvl         byte
	unknown2        uint16
	TimeStamp       uint16
	unknown3        uint16
	Hotkeys_0       uint16
	Hotkeys_1       uint16
	Hotkeys_2       uint16
	Hotkeys_3       uint16
	Hotkeys_4       uint16
	Hotkeys_5       uint16
	Hotkeys_6       uint16
	Hotkeys_7       uint16
	Hotkeys_8       uint16
	Hotkeys_9       uint16
	Hotkeys_10      uint16
	Hotkeys_11      uint16
	Hotkeys_12      uint16
	Hotkeys_13      uint16
	Hotkeys_14      uint16
	Hotkeys_16      uint16
	L_Mouse         uint16
	R_Mouse         uint16
	L_MouseAction   uint16
	R_MouseAction   uint16
	unknown4        [32]byte
	Difficulty      [3]byte
	MapID           uint16
	unknown5        uint8
	MercDead        uint8
	MercenaryID     uint16
	MercLangIndex   uint8
	MercenaryAttr   uint8
	MercenaryExp    uint16
	unknown6        [144]byte
}

func ReadGame(sg *SavedGame, r io.Reader) {
	buff := make([]byte, 4)
	r.Read(buff)
	x, _ := binary.Varint(buff)
	sg.FileId = uint16(x)
}
