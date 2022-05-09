package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	fd, err := os.Open("../namegen/names.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer fd.Close()

	output, err := os.OpenFile(
		"names.go", os.O_CREATE|os.O_RDWR|os.O_TRUNC,
		0755,
	)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(fd)

	fmt.Fprintln(
		output,
		`package main

// Run 'go generate' at the root of the repo on this package to update this file
`+`//go:generate go run ../namegen/.

import "sort"

func init() {
	NameLen = len(NameList)
	sort.Strings(NameList)
}

var NameLen int
var NameList = []string{`,
	)

	first := true
	for scanner.Scan() {
		if first {
			first = false
			continue
		}
		split := strings.Split(scanner.Text(), ";")
		fmt.Fprintf(
			output, "\t\"%s\",\n",
			strings.ReplaceAll(
				strings.ReplaceAll(
					split[2],
					" ", "_",
				),
				"\"", "",
			),
		)
	}

	fmt.Fprintln(
		output,
		`}`,
	)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
