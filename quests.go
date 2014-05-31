package d2s

type quests struct {

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

type waypoints struct {
	Unk15    uint16
	BitField [5]byte
	Unk16    [17]byte
}
