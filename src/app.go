package app

import (
	Generate "EDirac/src/generate"
	fmt "fmt"
)

func Init() {
	fmt.Println("eDirac Book Generator")

	Generate.Run()
}
