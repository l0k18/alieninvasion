package main

import (
	"errors"
	"fmt"
	"github.com/l0k18/alieninvasion/pkg/cities"
	"github.com/l0k18/alieninvasion/pkg/world"
	"os"
	"strconv"
)

func check(err error) {

	if err != nil {
		fmt.Println("alieninvasion world map grid generator")
		fmt.Printf("usage: %s <h> <v> <seed> <filename>\n", os.Args[0])
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

func main() {

	if len(os.Args) != 5 {
		check(errors.New("incorrect command line parameters"))
	}

	var h, v, seed int64
	var err error

	h, err = strconv.ParseInt(os.Args[1], 10, 64)
	check(err)

	v, err = strconv.ParseInt(os.Args[2], 10, 64)
	check(err)

	if int(h*v) > len(cities.NameList) {
		err = fmt.Errorf(
			"sorry, there isn't enough names in our list "+
				"for that many human settlements, "+
				"there is %d on earth over 1000 population",
			len(cities.NameList),
		)
		check(err)
	}

	seed, err = strconv.ParseInt(os.Args[3], 10, 64)
	check(err)

	w := world.Generate(h, v, seed)
	w.ToFile(os.Args[4])

}
