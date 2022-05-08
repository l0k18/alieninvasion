package alieninvasion

import "fmt"

const (
	N = iota
	E
	W
	S
	kvSep = "="
)

// directions are ordered so that `^dir & 3` is the opposite direction:
// n=00 e=01 w=10 s=11
// `& 3` masks the higher bits out
// `^` is bitwise NOT
var dirs = [4]string{"north", "east", "west", "south"}

type FromName map[string]int
type FromNumber map[int]string

type Lookup struct {
	Name  FromName
	Index FromNumber
}

func NewLookup() *Lookup {
	return &Lookup{
		Name:  make(FromName),
		Index: make(FromNumber),
	}
}

func (l *Lookup) Add(name string, index int) (err error) {

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
	Neighbor [4]int
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
		Lookup: NewLookup(), Cities: NewCities(),
	}
}
