package data

import (
	types "EDirac/src/types"
	fmt "fmt"
)

func Get() types.Data {
	// Name
	var name string
	fmt.Println("Wpisz nazwę książki (np. Biologia)")
	fmt.Scanln(&name)

	// var codename string
	// fmt.Println("Wpisz zakodowaną  (np. Biologia)")
	// fmt.Scanln(&codename)

	var id int
	fmt.Println("Wpisz id")
	fmt.Scanln(&id)

	var length int
	fmt.Println("Wpisz ilość tomów")
	fmt.Scanln(&length)

	// var author string
	// fmt.Println("Wpisz autora (autorów) ksiązki")
	// fmt.Scanln(&author)

	// var color string
	// fmt.Println("Wpisz autora (autorów) ksiązki")
	// fmt.Scanln(&color)

	return types.Data{Name: string(name), Length: int(length), ID: int(id)}
}
