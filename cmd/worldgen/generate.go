package main

import (
	"errors"
	"fmt"
	. "github.com/l0k18/alieninvasion/pkg/world"
	"math/rand"
	"os"
	"strconv"
)

// Line is a 1d array of names
type Line []string

// Grid is a 2d array of names
type Grid []Line

func result(err error) {

	if err != nil {
		fmt.Println("alieninvasion world map grid generator")
		fmt.Printf("usage: %s <h> <v> <seed> <filename>\n", os.Args[0])
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

func main() {

	if len(os.Args) != 5 {
		result(errors.New("incorrect command line parameters"))
	}

	var h, v, seed int64
	var err error

	h, err = strconv.ParseInt(os.Args[1], 10, 64)
	result(err)

	v, err = strconv.ParseInt(os.Args[2], 10, 64)
	result(err)

	if int(h*v) > len(NameList) {
		err = fmt.Errorf(
			"sorry, there isn't enough names in our list "+
				"for that many human settlements, "+
				"there is %d on earth over 1000 population", len(NameList),
		)
		result(err)
	}

	seed, err = strconv.ParseInt(os.Args[3], 10, 64)
	result(err)

	w := GenerateWorld(h, v, seed)
	w.ToFile(os.Args[4])

}

func GenerateWorld(h, v int64, seed int64) (w *World) {

	rand.Seed(seed)
	rand.Shuffle(
		len(NameList),
		func(i, j int) { NameList[i], NameList[j] = NameList[j], NameList[i] },
	)

	w = NewWorld()

	// We are going to generate a uniform grid of h*v cities from a 2d slice
	// of random generated names

	grid := make(Grid, v)
	for i := range grid {
		grid[i] = make(Line, h)
	}

	var name int
	for lat := range grid {

		for long := range grid[lat] {

			grid[lat][long] = NameList[name]
			name++
		}
	}

	latMax := int(v - 1)
	longMax := int(h - 1)

	for lat := range grid {

		for long := range grid[lat] {

			name := grid[lat][long]

			// wrap the map at the edges
			latN := lat - 1
			if latN < 0 {
				latN = latMax
			}
			latS := lat + 1
			if latS > latMax {
				latS = 0
			}
			longE := long - 1
			if longE < 0 {
				longE = longMax
			}
			longW := long + 1
			if longW > longMax {
				longW = 0
			}

			nN := grid[latN][long]
			nE := grid[lat][longE]
			nW := grid[lat][longW]
			nS := grid[latS][long]

			lineString := fmt.Sprintf("%s %s=%s %s=%s %s=%s %s=%s",
				name, Dirs[N], nN, Dirs[E], nE, Dirs[W], nW, Dirs[S], nS)

			err := w.AddFromString(lineString)
			if err != nil {
				fmt.Println("error adding", lat, long, name, err)
				os.Exit(1)
			}
		}
	}

	return
}
