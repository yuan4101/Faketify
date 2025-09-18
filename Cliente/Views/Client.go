package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"localServer/grpc-songServer/songServices"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50053", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("Error connecting to gRPC server: %v", err)
	}
	defer conn.Close()

	c := songServices.NewSongServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	//Se captura el título de la canción a buscar
	fmt.Printf("\nSong title: ")
	reader := bufio.NewReader(os.Stdin)
	readedTitle, _ := reader.ReadString('\n')
	readedTitle = strings.TrimSpace(readedTitle)

	//Se crea un objeto de tipo DTO que contiene el título de la canción a buscar
	songRequestObj := &songServices.SongRequest{Title: readedTitle}

	//Llamada al procedimiento remoto buscarCancion
	res, err := c.GetSong(ctx, songRequestObj)
	if err != nil {
		fmt.Printf("Error calling gRPC: %v", err)
		return
	}

	//Impresión de la respuesta
	fmt.Printf("\nMessage: %s", res.Message)
	fmt.Printf("\nCode: %d\n", res.Code)

	if res.Code == 200 {
		fmt.Printf("\nSong found:\n")
		fmt.Printf("Title: %s\n", res.SongObj.Title)
		fmt.Printf("Artist: %s\n", res.SongObj.Artist)
		fmt.Printf("Year: %d\n", res.SongObj.Year)
		fmt.Printf("Duration: %s seconds\n", res.SongObj.Duration)
		fmt.Printf("Genre: %s\n\n", res.SongObj.Genre.Name)
	}
}
