package main

import (
	"fmt"

	"./movies"
)

func main() {
	m := movies.NewMovie("thing", movie.CHILDRENS)
	fmt.Printf("%s\n", m)
}
