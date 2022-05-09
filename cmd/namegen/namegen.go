package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
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

	names := make([]string, 139661)
	scanner := bufio.NewScanner(fd)
	var counter int
	for scanner.Scan() {
		if counter == 0 {
			counter++
			continue
		}
		split := strings.Split(scanner.Text(), ";")
		names[counter] =
			strings.ReplaceAll(
				strings.ReplaceAll(
					strings.ReplaceAll(
						split[2],
						" ", "_",
					),
					"'", "",
				),
				"\"", "",
			)
		counter++
	}

	sort.Strings(names)

	fmt.Fprintln(
		output,
		`package main
	
// Run 'go generate ./...' at the root of the repo on this package to update
`+`//go:generate go run ../namegen/.
		
var NameList = []string{
	"",`,
	)
	prev := names[0]
	for i := 1; i < counter; i++ {
		if names[i] == prev || names[i] == "" {
			continue
		}
		fmt.Fprintf(output, "\t\"%s\",\n", names[i])
		prev = names[i]
	}
	fmt.Fprintln(
		output,
		`}`,
	)

	output.Close()
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
