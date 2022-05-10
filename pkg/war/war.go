package war

import (
	"fmt"
	"github.com/l0k18/alieninvasion/pkg/world"
	"math/rand"
	"os"
)

type Aliens map[uint32]uint32

func War(w *world.World, aliens, seed int64) {

	aliens++

	// Create a slice matching the cities, and then shuffle them
	cityCount := len(w.Cities)
	cities := make([]int32, cityCount)

	for i := range cities {

		if int64(i) < aliens {

			// add ascending numbers to represent the aliens
			cities[i] = int32(i)

		} else {

			// once aliens numbers are added, we are finished
			break
		}
	}

	// to initially place aliens we create a slice of the cities and shuffle it
	rand.Seed(seed)

	// to leave the first element alone we shorten the length by 1 and add 1 to
	// the fields to be shuffled
	rand.Shuffle(
		cityCount-1,
		func(i int, j int) {
			cities[i+1], cities[j+1] = cities[j+1], cities[i+1]
		},
	)

	alienMap := make(Aliens, aliens)

	for i := range cities {

		if cities[i] != 0 {

			alienMap[uint32(cities[i])] = uint32(i)

			// city := uint32(cities[i])
			// fmt.Println(
			//     "Alien",
			//     cities[i],
			//     w.Lookup.Index[city],
			//     world.Dirs[world.N],
			//     w.Lookup.Index[w.Cities[city].Neighbour[world.N]],
			//     world.Dirs[world.E],
			//     w.Lookup.Index[w.Cities[city].Neighbour[world.E]],
			//     world.Dirs[world.W],
			//     w.Lookup.Index[w.Cities[city].Neighbour[world.W]],
			//     world.Dirs[world.S],
			//     w.Lookup.Index[w.Cities[city].Neighbour[world.S]],
			// )
		}
	}
	w.Print(os.Stdout)
	// run for a maximum of 10000 turns
	for turn := 0; turn < 10000 && len(alienMap) > 1; turn++ {

		//fmt.Println("aliens", alienMap)
		// iterate the list of aliens and move them to a new location
		//
		// we move first then check collisions, so it is possible for up to 4
		// aliens to collide and annihilate one city
		for i := range alienMap {

			// randomly select a direction to move each alien
			moveDir := rand.Intn(5)
			//
			//if moveDir < 3 {
			//	fmt.Println(alienMap[i], "moveDir", world.Dirs[moveDir])
			//} else {
			//	fmt.Println(alienMap[i], "not moving")
			//}

			// If the random number is 4 the alien decides not to move
			if moveDir > 3 {
				continue
			}

			// get the index of the alien's current location
			city := alienMap[i]

			// change the alien's location to the new location
			newLoc := w.Cities[city].Neighbour[moveDir]

			// if the new location is index 0 the alien cannot move this
			// direction and forfeits this turn
			if newLoc != 0 {

				alienMap[i] = newLoc
			}
		}

		// we will use a map to detect location collision, when more than one
		// alien is in a city each alien's identity is appended to a list
		detector := make(map[uint32][]uint32)

		// iterate the aliens and append them to the map entries
		for i := range alienMap {

			if _, ok := detector[alienMap[i]]; !ok {

				// if the map entry is empty, add the initial alien
				detector[alienMap[i]] = []uint32{i}

			} else {

				// if there is already an entry, append the new alien
				detector[alienMap[i]] = append(detector[alienMap[i]], i)
			}
		}

		// iterate the detector and delete cities
		for city := range detector {

			// if the map entry for a city has more than one alien it is
			// destroyed
			if len(detector[city]) > 1 {

				fmt.Print("Aliens ")

				// first we will print that the aliens have collided and
				// destroyed the city
				for alien := range detector[city] {

					// if this is the last in the list, separate with and
					if alien == len(detector[city])-1 {

						fmt.Print(" and ")

						// if this is not the first, separate with a comma
					} else if alien != 0 {

						fmt.Print(", ")
					}

					fmt.Print(detector[city][alien])
					// lastly, delete the alien
					delete(alienMap, detector[city][alien])

				}
				fmt.Printf(
					" converged on %s and destroyed it!\n",
					w.Lookup.Index[city],
				)

				// remove neighbours references to the destroyed city
				for dir, neighbour := range w.Cities[city].Neighbour {

					// the neighbour in a given direction the opposite numbered
					// neighbour must be zeroed to the empty city
					w.Cities[neighbour].Neighbour[^dir&3] = 0

					// then we delete the outbound neighbour entry
					w.Cities[city].Neighbour[dir] = 0
				}
				idx := w.Lookup.Index[city]
				delete(w.Lookup.Index, city)
				delete(w.Lookup.Name, idx)

			}
		}
	}
	w.Print(os.Stdout)

}
