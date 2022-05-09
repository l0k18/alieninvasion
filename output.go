package alieninvasion

import "github.com/davecgh/go-spew/spew"

func (w *World) ToFile(filename string) {
	spew.Dump(w)
}
