package alieninvasion

import (
	"errors"
	"fmt"
	"strings"
)

func (w *World) AddFromString(input string) (err error) {

	// City name is first and then space separated up to 4 neighbours one for
	// each direction
	split := strings.Split(input, " ")

	// All names should be standardised to title case in case of manual input
	newName := strings.ToTitle(split[0])

	// The first string should be a city name and not a key value pair
	if strings.Contains(newName, kvSep) {

		return errors.New("malformed city name: may not contain '='")
	}

	// There cannot be more than 4 directions and a city so the line is
	// malformed if it is more than 5 items (with 4 spaces)
	if len(split) > 5 {

		return fmt.Errorf("input string \"%s\" contains too many fields", input)
	}

	neighbours := split[1:]
	if len(neighbours) > 0 {

		// Validate that all specified neighbours are `dir=Name`
		for i, v := range neighbours {

			if !strings.Contains(v, kvSep) {

				return fmt.Errorf(
					"input string \"%s\" field %d contains malformed key/value pair \"%s\"",
					input, i, v,
				)
			} else {

				if strings.Count(v, kvSep) > 1 {

					return fmt.Errorf(
						"input string \"%s\" contains multiple key"+
							"/value separators on element %d \"%s\"",
						input, i, v,
					)
				}
			}

		}
	}

	// Add new index to lookup table
	city := w.Length
	err = w.Lookup.Add(newName, city)
	if err != nil {

		// This means we are trying to add the same city on a new index
		return err
	}

	// Append new empty city entry with matching index
	w.Cities = append(w.Cities, City{})

	// The Cities slice is now one element longer,
	// next addition must be on the next index
	w.Length++

	if len(neighbours) > 0 {

		for i, v := range neighbours {

			// Split the key and value
			kvs := strings.Split(v, kvSep)

			// We want to allow any variant of case and only count the first
			// character, as it is distinctive
			key := strings.ToLower(kvs[0][:1])

			// this ensures variants of case are standardised to Title case
			// (first character is capital, the rest lower case)
			newNeighbour := strings.ToTitle(kvs[1])

			var validDir bool
			var dir int

			// Check that the key is a valid direction string.
			// Note we only compare the first character.
			for d := range dirs {
				if key == dirs[d][:1] {

					validDir = true
					dir = d
					break
				}
			}
			if !validDir {

				return fmt.Errorf(
					"specified neighbour of %s %d \"%s\" using invalid key \"%s\"",
					newName, i, v, key,
				)
			}

			var n int
			var ok bool
			// if the key doesn't exist, create it
			if n, ok = w.Lookup.Name[newNeighbour]; !ok {

				// automatically set current to be opposite neighbour
				newCity :=
					fmt.Sprintf("%s %s=%s", newNeighbour, dirs[^dir&3], newName)

				err := w.AddFromString(newCity)
				if err != nil {

					return err
				}

				// If the neighbour in the specified direction already exists
				// and isn't the same neighbour as was given the input is
				// invalid
			} else if w.Cities[city].Neighbor[dir] != n {

				return fmt.Errorf(
					"city already exists in direction %s"+
						" from city %s with name %s, cannot overwrite with %s",
					dirs[dir], newName, w.Lookup.Index[city], newNeighbour,
				)
			}

		}

	}

	return nil
}
