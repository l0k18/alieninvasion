package alieninvasion

import "fmt"

func (w *World) ToFile(filename string) {
	for i := range w.Lookup.Index {
		if i == 0 {
			continue
		}
		neighbours := [4]string{}
		for n := 0; n < 4; n++ {
			neighbours[n] = w.Lookup.Index[w.Cities[i].Neighbour[n]]
		}
		fmt.Printf(
			"%s %s=%s %s=%s %s=%s %s=%s\n",
			w.Lookup.Index[i],
			Dirs[0],
			neighbours[0],
			Dirs[1],
			neighbours[1],
			Dirs[2],
			neighbours[2],
			Dirs[3],
			neighbours[3],
		)
	}
}
