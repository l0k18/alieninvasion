package world

import (
	"fmt"
	"io"
	"log"
	"os"
)

// ToFile - Print the world description to a file
func (w *World) ToFile(filename string) {

	output, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0755)
	if err != nil {

		log.Fatal(err)
	}

	defer output.Close()

	w.Print(output)
}

// Print the plain text document that encodes the current in memory data structure to an io.Writer
func (w *World) Print(output io.Writer) {

	for i := range w.Lookup.Index {

		// skip the first, empty city in the index
		if i == 0 {

			continue
		}

		output.Write([]byte(w.Lookup.Index[i]))

		for n := 0; n < 4; n++ {

			neighbour := w.Cities[i].Neighbour[n]

			// only append the item if it doesn't refer to the empty city
			if neighbour != 0 {

				fmt.Fprintf(output, " %s=%s", Dirs[n], w.Lookup.Index[neighbour])
			}
		}

		fmt.Fprint(output, "\n")
	}

}
