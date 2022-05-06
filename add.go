package alieninvasion

import (
	"errors"
	"fmt"
	"strings"
)

func (w *World) AddFromString(input string) (err error) {

	split := strings.Split(input, " ")

	// all names should be standardised to title case in case of manual input
	newName := strings.ToTitle(split[0])
	if strings.Contains(newName, kvSep) {

		return errors.New("malformed city name: may not contain '='")
	}

	if len(split) > 5 {

		return fmt.Errorf("input string \"%s\" contains too many fields", input)
	}

	neighbours := split[1:]
	if len(neighbours) > 0 {

		for i, v := range neighbours {

			if !strings.Contains(v, kvSep) {

				return fmt.Errorf("input string \"%s\" field %d contains malformed key/value pair \"%s\"",
					input, i, v)
			} else {
				strings.Count(v, kvSep)
			}

		}
	}

	// check if name has already been created by a neighbour

	w.Index[newName] = w.Length

	w.Cities = append(w.Cities, City{})

	if len(neighbours) > 0 {

		for i, v := range neighbours {

			kvs := strings.Split(v, kvSep)

			// we want to allow any variant of case and only count the first character, as it is distinctive
			key := strings.ToLower(kvs[0][:1])

			// this ensures variants of case are standardised to Title case (first character is capital)
			val := strings.ToTitle(kvs[1])

			var validDir bool
			var dir int

			for d := range dirs {

				if key == dirs[d] {
					validDir = true
					dir = d
					break
				}
			}

			if !validDir {

				return fmt.Errorf("specified neighbour of %s %d \"%s\" using invalid key \"%s\"",
					newName, i, v, key)
			}

			currIndex := w.Length
			w.Length++

			var n int
			var ok bool
			// if the key doesn't exist, create it
			if n, ok = w.Index[val]; !ok {

				err := w.AddFromString(val)
				if err != nil {

					return err
				}

				n = currIndex + 1
			}

			// either way, save the neighbour's index
			w.Cities[currIndex].Neighbor[dir] = n
		}

	}

	return nil
}
