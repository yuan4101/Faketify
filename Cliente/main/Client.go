package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"localClient/grpc-client/Views"
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

	Views.ShowMenu()

	fmt.Printf("\nOpcion: ")
	reader := bufio.NewReader(os.Stdin)
	varOpt, _ := reader.ReadString('\n')
	varOpt = strings.TrimSpace(varOpt)

	if varOpt == "0" {
		//Se captura el título de la canción a buscar
		fmt.Printf("\nSong title: ")
		reader := bufio.NewReader(os.Stdin)
		readedTitle, _ := reader.ReadString('\n')
		readedTitle = strings.TrimSpace(readedTitle)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()

		//Se crea un objeto de tipo DTO que contiene el título de la canción a buscar
		songRequestObj := &songServices.SongRequest{Title: readedTitle}

		//Llamada al procedimiento remoto buscarCancion
		res, err := c.GetSong(ctx, songRequestObj)
		if err != nil {
			fmt.Printf("Error calling gRPC: %v", err)
			return
		}
		Views.DisplaySongMetadataResponse(res)
	}

	if varOpt == "1" {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()

		res, err := c.GetGenres(ctx, &songServices.Empty{})
		if err != nil {
			fmt.Printf("Error calling gRPC: %v", err)
			return
		}
		Views.ShowMenu(res)
	}
}
