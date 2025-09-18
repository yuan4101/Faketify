package Views

import (
	"fmt"
	"localServer/grpc-songServer/songServices"
)

func ShowMenu(genres ...*songServices.ResponseGenresDTO) {
	fmt.Print("\t Faketify \t\n\n")

	if len(genres) > 0 && genres[0] != nil {
		lastIndex := 0
		for i, genre := range genres[0].GenresObjArr {
			fmt.Printf("%d. %s\n", i+1, genre.Name)
			lastIndex = i + 2
		}
		fmt.Printf("%d. Atras\n", lastIndex)
	} else {
		fmt.Println("1. Ver generos")
		fmt.Println("2. Salir")
	}
}
