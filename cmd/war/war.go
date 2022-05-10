package main

import (
	"fmt"
	"github.com/l0k18/alieninvasion/pkg/war"
	"github.com/l0k18/alieninvasion/pkg/world"
	"os"
	"strconv"
)

func result(err error) {
	res := 0
	if err != nil {
		fmt.Println("Error:", err)
		res = 1
	}
	fmt.Println("alieninvasion war simulator")
	fmt.Printf("usage: %s <aliencount> <seed> <filename>\n", os.Args[0])
	os.Exit(res)
}

func main() {

	if len(os.Args) != 4 {
		result(nil)
	}

	var aliens, seed int64
	var err error

	aliens, err = strconv.ParseInt(os.Args[1], 10, 64)
	if err != nil {
		result(err)
	}

	seed, err = strconv.ParseInt(os.Args[2], 10, 64)
	if err != nil {
		result(err)
	}

	w := world.New()
	w.AddFromFile(os.Args[3])

	war.War(w, aliens, seed)
	// w.Print(os.Stdout)
}
