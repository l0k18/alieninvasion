package world

import (
	"fmt"
	"github.com/l0k18/alieninvasion/pkg/cities"
	"math/rand"
	"os"
)

// Line is a 1d array of names
type Line []string

// Grid is a 2d array of names
type Grid []Line

func Generate(h, v int64, seed int64) (w *World) {

	rand.Seed(seed)
	rand.Shuffle(
		len(cities.NameList),
		func(i, j int) { cities.NameList[i], cities.NameList[j] = cities.NameList[j], cities.NameList[i] },
	)

	w = New()

	// We are going to generate a uniform grid of h*v cities from a 2d slice
	// of random generated names

	grid := make(Grid, v)
	for i := range grid {
		grid[i] = make(Line, h)
	}

	var name int
	for lat := range grid {

		for long := range grid[lat] {

			grid[lat][long] = cities.NameList[name]
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
