package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	fd, err := os.Open("./cmd/namegen/names.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer fd.Close()

	output, err := os.OpenFile(
		"./cmd/worldgen/names.go", os.O_CREATE|os.O_RDWR,
		0755,
	)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(fd)

	fmt.Fprintln(
		output,
		`package main

import "sort"

func init() {
	NameLen = len(nameList)
	sort.Strings(nameList)
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
					split[1],
					" ", "_",
				),
				"\"", "",
			),
		)
	}

	fmt.Fprintln(
		output,
		`}
`,
	)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
