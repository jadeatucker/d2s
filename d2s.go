package d2s

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"regexp"
)

const h_OFFSET int64 = 767
const nameRegexp = "^[^-_][a-zA-Z]+[-_]?[a-zA-Z]+[^-_]$"

const (
	CLASS_AMAZON      = 0x00
	CLASS_SORCERESS   = 0x01
	CLASS_NECROMANCER = 0x02
	CLASS_PALADIN     = 0x03
	CLASS_BARBARIAN   = 0x04
	CLASS_DRUID       = 0x05
	CLASS_ASSASSIN    = 0x06
)

type Class struct {
	Class byte
}

type SavedGame struct {
	header saveFile
	buffer []byte
}

type saveFile struct {
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

func New(r io.Reader, size int64) (sg *SavedGame, err error) {
	if size < 0 {
		return nil, fmt.Errorf("d2s.ReadGame error: invalid size")
	} else if size < h_OFFSET {
		return nil, fmt.Errorf("d2s.ReadGame error: size too small")
	}

	sg = new(SavedGame)
	// Read static header struct
	err = binary.Read(r, binary.LittleEndian, &sg.header)
	if err != nil {
		return nil, fmt.Errorf("d2s.ReadGame failed to read header: %v", err)
	}

	// Read remaining into dynamic buffer
	sg.buffer = make([]byte, size-h_OFFSET)
	_, err = r.Read(sg.buffer)
	if err != nil {
		return nil, fmt.Errorf("d2s.ReadGame failed to read: %v", err)
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
	n += copy(p[h_OFFSET:], sg.buffer)

	return
}

func (sg *SavedGame) Checksum() uint32 {
	p := &sg.header.Checksum
	c := *p
	*p = 0

	b := make([]byte, h_OFFSET+int64(len(sg.buffer)))
	_, err := sg.Read(b)
	if err != nil {
		*p = c
		return 0
	}

	*p = checksum(b, 0)
	return *p
}

func checksum(b []byte, chk uint32) uint32 {
	for _, c := range b {
		chk = ((chk << 1) | (chk >> 31)) + uint32(c)
	}
	return chk
}

func (sg *SavedGame) Name() (str string) {
	b := sg.header.CharName
	i := bytes.IndexByte(b[:], 0)

	if i > 0 {
		str = string(b[:i])
	}
	return
}

// Sets CharName field, returns an error if name is not valid.
// Valid names must be 2-15 characters in length, and can only contain
// characters from the english alphabet with exception to one dash(-) or
// underscore(_) as long as it is not the first or last character.
func (sg *SavedGame) SetName(name string) (err error) {
	if len(name) < 2 || len(name) > 15 {
		return fmt.Errorf("d2s.SetName name must be 2-15 characters long")
	}

	match, _ := regexp.MatchString(nameRegexp, name)

	if !match {
		return fmt.Errorf("d2s.SetName invalid name: %s", name)
	}

	for i := 0; i < len(sg.header.CharName); i++ {
		if i < len(name) {
			sg.header.CharName[i] = name[i]
		} else {
			sg.header.CharName[i] = 0
		}
	}

	return
}

func (sg *SavedGame) Class() Class {
	return sg.header.CharClass
}

func (c Class) String() (str string) {
	switch c.Class {
	case CLASS_AMAZON:
		str = "Amazon"
	case CLASS_ASSASSIN:
		str = "Assasin"
	case CLASS_SORCERESS:
		str = "Sorceress"
	case CLASS_DRUID:
		str = "Druid"
	case CLASS_PALADIN:
		str = "Paladin"
	case CLASS_BARBARIAN:
		str = "Barbarian"
	case CLASS_NECROMANCER:
		str = "Necromancer"
	}
	return
}

func (sg *SavedGame) SetClass(class Class) (err error) {
	switch class.Class {
	case CLASS_AMAZON:
		break
	case CLASS_ASSASSIN:
		break
	case CLASS_SORCERESS:
		break
	case CLASS_DRUID:
		break
	case CLASS_PALADIN:
		break
	case CLASS_BARBARIAN:
		break
	case CLASS_NECROMANCER:
		break
	default:
		return fmt.Errorf("d2s.SetClass invalid value for class: %v", class)
	}
	sg.header.CharClass = class
	return
}
