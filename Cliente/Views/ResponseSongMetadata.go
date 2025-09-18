package Views

import (
	"fmt"
	"localServer/grpc-songServer/songServices"
)

func DisplaySongMetadataResponse(res *songServices.ResponseSongDTO) {
	//Impresión de la respuesta
	fmt.Printf("\nMessage: %s", res.Message)
	fmt.Printf("\nCode: %d\n", res.Code)

	if res.Code == 200 {
		fmt.Printf("\nSong found:\n")
		fmt.Printf("Title: %s\n", res.SongObj.Title)
		fmt.Printf("Artist: %s\n", res.SongObj.Artist)
		fmt.Printf("Year: %d\n", res.SongObj.Year)
		fmt.Printf("Duration: %s minutes\n", res.SongObj.Duration)
		fmt.Printf("Genre: %s\n\n", res.SongObj.Genre.Name)
	}
}
