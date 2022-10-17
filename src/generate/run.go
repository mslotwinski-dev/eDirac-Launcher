package generate

import (
	Data "EDirac/src/data"
	Types "EDirac/src/types"

	os "os"
	filepath "path/filepath"
)

func Run() {

	var D Types.Data = Data.Get()

	newpath := filepath.Join(".", "/"+D.Name)
	os.MkdirAll(newpath, os.ModePerm)

	for i := D.ID; i < D.ID+D.Length; i++ {
		Generate(newpath, i)
	}

}
