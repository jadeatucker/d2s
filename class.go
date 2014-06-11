package d2s

import "fmt"

const (
	ClassAmazon      = 0x00
	ClassSorceress   = 0x01
	ClassNecromancer = 0x02
	ClassPaladin     = 0x03
	ClassBarbarian   = 0x04
	ClassDruid       = 0x05
	ClassAssasin     = 0x06
)

type Class struct {
	Class byte
}

// String returns the name of the class as a string.
func (c Class) String() (str string) {
	switch c.Class {
	case ClassAmazon:
		str = "Amazon"
	case ClassAssasin:
		str = "Assasin"
	case ClassSorceress:
		str = "Sorceress"
	case ClassDruid:
		str = "Druid"
	case ClassPaladin:
		str = "Paladin"
	case ClassBarbarian:
		str = "Barbarian"
	case ClassNecromancer:
		str = "Necromancer"
	}
	return
}

// Class returns the Class of the SavedGame character.
func (sg *SavedGame) Class() Class {
	return sg.header.CharClass
}

// SetClass takes a Class and assigns it to the SavedGame.
// A class.Class must equal one of ClassAmazon, ClassAssasin, ClassSorceress,
// ClassDruid, ClassPaladin, ClassBarbarian of ClassNecromancer.  Any other
// values will return an error and the class will not be changed.
func (sg *SavedGame) SetClass(class Class) (err error) {
	switch class.Class {
	case ClassAmazon, ClassAssasin, ClassSorceress,
		ClassDruid, ClassPaladin, ClassBarbarian,
		ClassNecromancer:
	default:
		return fmt.Errorf("d2s.SetClass invalid value for class: %v", class.Class)
	}
	sg.header.CharClass = class
	return
}
