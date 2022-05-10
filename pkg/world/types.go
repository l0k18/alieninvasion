package world

import "fmt"

const (
	N = iota
	E
	W
	S
	kvSep = "="
)

// Dirs are ordered so that `^dir & 3` is the opposite direction:
// n=00 e=01 w=10 s=11
// `& 3` masks the higher bits out
// `^` is bitwise NOT
var Dirs = [4]string{"north", "east", "west", "south"}

type FromName map[string]uint32
type FromNumber map[uint32]string

type Lookup struct {
	Name  FromName
	Index FromNumber
}

func NewLookup() *Lookup {
	return &Lookup{
		Name:  FromName{"": 0},
		Index: FromNumber{0: ""},
	}
}

func (l *Lookup) Add(name string, index uint32) (err error) {

	// Check for existing entries
	if n, ok := l.Name[name]; ok {
		return fmt.Errorf(
			"name conflict: "+
				"%s already exists with different index %d from submitted %d",
			name, n, index,
		)
	}
	// Consuming code only adds new entries as it grows the Cities slice so
	// this error cannot occur, leaving this here for hypothetical
	// if n, ok := l.Index[index]; ok {
	// 	return fmt.Errorf(
	// 		"index conflict: "+
	// 			"%d already exists with different name %s from submitted %s",
	// 		index, n, name,
	// 	)
	// }

	l.Name[name] = index
	l.Index[index] = name
	return
}

// City name is not stored in the structure as the Index is the proper source
type City struct {
	Neighbour [4]uint32
}

type Cities []City

func NewCities() Cities { return Cities{City{}} }

type World struct {
	Length int
	*Lookup
	Cities
}

func NewWorld() *World {
	return &World{
		Length: 1, Lookup: NewLookup(), Cities: NewCities(),
	}
}
