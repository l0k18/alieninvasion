package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	fd, err := os.Open("names.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer fd.Close()

	output, err := os.OpenFile(
		"../worldgen/names.go", os.O_CREATE|os.O_RDWR,
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
	sort.Strings(nameList)
}

var nameList = []string{`,
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
				split[1],
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
