package d2s

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
)

const H_OFFSET int64 = 767

type SavedGame struct {
	FileHeader SaveFile
	FileBuff   []byte
}

type SaveFile struct {
	FileId          uint32
	FileVersion     uint32
	FileSize        uint32
	Checksum        uint32
	ActiveArms      uint32
	CharName        [16]byte
	CharStatus      byte
	CharProgression byte
	Unk0            [2]byte
	CharClass       byte
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
	QuestsNormal    Quests
	QuestsNightmare Quests
	QuestsHell      Quests

	// Waypoint data for each difficulty level
	WS   [2]byte
	Unk8 [6]byte

	WPSNormal    Waypoints
	WPSNightmare Waypoints
	WPSHell      Waypoints
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

type Quests struct {

	// ACT I
	Warriv                uint16
	DenOfEvil             uint16
	SistersBurialGrounds  uint16
	ToolsoftheTrade       uint16
	TheSearchforCain      uint16
	TheForgottenTower     uint16
	SisterstotheSlaughter uint16
	Unk10                 [2]byte

	// ACT II
	Jerhyn           uint16
	RadamentsLair    uint16
	TheHoradricStaff uint16
	TaintedSun       uint16
	ArcaneSanctuary  uint16
	TheSummoner      uint16
	TheSevenTombs    uint16
	Unk11            [2]byte

	// ACT III
	Hratli                uint16
	LamEsensTome          uint16
	KhalimsWill           uint16
	BladeoftheOldReligion uint16
	TheGoldenBird         uint16
	TheBlackenedTemple    uint16
	TheGuardian           uint16
	Unk12                 [2]byte

	// ACT IV
	Act_IV         uint16
	TheFallenAngel uint16
	TerrorsEnd     uint16
	HellsForge     uint16
	Unk13          [14]byte

	// ACT V
	SiegeonHarrogath    uint16
	RescueonMountArreat uint16
	PrisonofIce         uint16
	BetrayalofHarrogath uint16
	RiteofPassage       uint16
	EveofDestruction    uint16
	Unk14               [14]byte
}

type Waypoints struct {
	Unk15    uint16
	BitField [5]byte
	Unk16    [17]byte
}

func ReadGame(sg *SavedGame, r io.ReadSeeker, size int64) (err error) {
	if size < 0 {
		return fmt.Errorf("d2s.ReadGame error: invalid size")
	} else if size < H_OFFSET {
		return fmt.Errorf("d2s.ReadGame error: size too small")
	}

	// Read static header struct
	err = binary.Read(r, binary.LittleEndian, &sg.FileHeader)
	if err != nil {
		return fmt.Errorf("d2s.ReadGame failed to read header: %v", err)
	}

	// Read remaining into dynamic buffer
	sg.FileBuff = make([]byte, size-H_OFFSET)
	_, err = r.Read(sg.FileBuff)
	if err != nil {
		return fmt.Errorf("d2s.ReadGame failed to read: %v", err)
	}

	return nil
}

func (sg *SavedGame) Write(p []byte) (n int, err error) {
	//sg.Checksum()
	// perform algorithm
	// write checksum
	// write struct to buffer
	// write dynamic data to buffer
	return 0, nil
}

func (sg *SavedGame) Checksum() error {
	sh := &sg.FileHeader
	var sum uint32 = 0
	sh.Checksum = 0

	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, sg.FileHeader)
	if err != nil {
		return fmt.Errorf("binary.Write failed:  %v", err)
	}

	c, err := buf.ReadByte()
	for err != io.EOF {
		sum = ((sum << 1) | (sum >> 31)) + uint32(c)
		c, err = buf.ReadByte()
	}

	for _, c := range sg.FileBuff {
		sum = ((sum << 1) | (sum >> 31)) + uint32(c)
	}

	sh.Checksum = sum
	return nil
}
