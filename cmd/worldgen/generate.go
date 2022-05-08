package main

import (
	"errors"
	"fmt"
	"github.com/l0k18/alieninvasion"
	"math/rand"
	"os"
	"strconv"
)

// Grid is a 2d array of names
type Grid [][]string

func fail(err error) {
	fmt.Println(err)
	os.Exit(1)
}

func main() {

	var h, v, seed int64
	var err error

	h, err = strconv.ParseInt(os.Args[1], 10, 64)
	if err != nil {
		fail(err)
	}

	v, err = strconv.ParseInt(os.Args[2], 10, 64)
	if err != nil {
		fail(err)
	}

	if int(h*v) > nameLen {
		err = errors.New(
			"sorry, there isn't enough names in our list" +
				"for that many human settlements",
		)
		fail(err)
	}

	seed, err = strconv.ParseInt(os.Args[3], 10, 64)
	if err != nil {
		fail(err)
	}

	w := GenerateWorld(h, v, seed)
	w.ToFile("filename")

}

func GenerateWorld(h, v int64, seed int64) (w *alieninvasion.World) {

	rand.Seed(seed)

	w = alieninvasion.NewWorld()

	// We are going to generate a uniform grid of h*v cities from a 2d slice
	// of random generated names

	grid := make([][]string, h)
	for i := range grid {
		grid[i] = make([]string, v)
	}

	rand.Seed(seed)
	total := int(h * v)

	for lat := range grid {

		for long := range grid[lat] {

			name := rand.Intn(total)
			grid[lat][long] = nameList[name]
		}
	}

	return
}
