package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	utilities "localClient/grpc-client/Utilities"
	"localClient/grpc-client/Views"
	"localServer/grpc-songServer/songServices"
	"localServer/grpc-streamingServer/streamingServices"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50053", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("Error connecting to gRPC server: %v", err)
	}
	defer conn.Close()

	connection := songServices.NewSongServiceClient(conn)

	conn2, err2 := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err2 != nil {
		log.Fatal(err2)
	}
	defer conn2.Close()

	client := streamingServices.NewAudioServiceClient(conn2)

	exitOption := "0"
	varOpt := ""
	songOption := ""
	genreOption := ""
	var varSong *songServices.ResponseSongDTO

	for exitOption != "-1" {

		if exitOption == "0" {
			Views.ShowMainMenu()
			varOpt = read("Opcion: ")
			switch varOpt {
			case "2":
				varOpt = "5"
			case "3", "4":
				varOpt = "0"
			}
		}

		switch varOpt {
		case "1":
			genres := getGenres(connection)
			if len(genres.GenresObjArr) == 0 {
				fmt.Println("No hay generos disponibles")
				exitOption = "0"
				continue
			}
			exitOption = Views.ShowGenresMenu(genres)
			genreOption = read("Opcion: ")
			genreIndex, _ := strconv.Atoi(genreOption)
			exitIndex, _ := strconv.Atoi(exitOption)

			if genreOption == exitOption {
				exitOption = "0"
				continue
			} else if genreIndex > 0 && genreIndex < exitIndex {
				exitOption = "1"
				varOpt = "2"
				genreOption = genres.GenresObjArr[genreIndex-1].GetName()
				continue
			}
			fmt.Println("-> ERROR: Opcion no valida")

		case "2":
			songs := getSongsByGenre(connection, genreOption)
			if len(songs.SongsObjArr) == 0 {
				fmt.Printf("No hay canciones disponibles para el genero: %s", genreOption)
				exitOption = "1"
				varOpt = "1"
				continue
			}
			exitOption = Views.ShowSongsMenu(genreOption, songs)
			songOption = read("Opcion: ")
			songIndex, _ := strconv.Atoi(songOption)
			exitIndex, _ := strconv.Atoi(exitOption)

			if songOption == exitOption {
				exitOption = "1"
				varOpt = "1"
				continue
			} else if songIndex > 0 && songIndex < exitIndex {
				exitOption = "2"
				varOpt = "3"
				songOption = songs.SongsObjArr[songIndex-1].GetTitle()
				continue
			}
			fmt.Println("-> ERROR: Opcion no valida")

		case "3":
			varSong = getSong(connection, songOption)
			if varSong == nil {
				fmt.Printf("Cancion: %s no disponible", songOption)
				exitOption = "2"
				varOpt = "2"
				continue
			}
			exitOption = Views.ShowSongMenu(varSong)
			songAction := read("Opcion: ")
			actionIndex, _ := strconv.Atoi(songAction)
			exitIndex, _ := strconv.Atoi(exitOption)

			if songAction == "2" {
				exitOption = "2"
				varOpt = "2"
				continue
			} else if actionIndex > 0 && actionIndex < exitIndex {
				exitOption = "3"
				varOpt = "4"
				continue
			}
			fmt.Println("-> ERROR: Opcion no valida")

		case "4":
			getBack := getStreamingSong(client, varSong)
			if getBack {
				exitOption = "3"
				varOpt = "3"
			} else {
				exitOption = "4"
				varOpt = "3"
			}

		case "5":
			Views.ColorStringPrint("\nSaliendo...\n", "yellow", false)
			exitOption = "-1"

		default:
			fmt.Println("-> ERROR: Opcion no valida")
		}
	}
}

func read(prmMessage string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("\n%s", prmMessage)
	varReaded, _ := reader.ReadString('\n')
	varReaded = strings.TrimSpace(varReaded)
	return varReaded
}

func getGenres(connection songServices.SongServiceClient) *songServices.ResponseGenresDTO {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res, err := connection.GetGenres(ctx, &songServices.Empty{})
	if err != nil {
		fmt.Printf("Error calling gRPC: %v", err)
		return nil
	}
	return res
}

func getSongsByGenre(connection songServices.SongServiceClient, prmGenre string) *songServices.ResponseSongsDTO {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	genreRequestObj := &songServices.SongsByGenreRequest{GenreName: prmGenre}

	res, err := connection.GetSongsByGenre(ctx, genreRequestObj)
	if err != nil {
		fmt.Printf("Error calling gRPC: %v", err)
		return nil
	}
	return res
}

func getSong(connection songServices.SongServiceClient, prmSongTitle string) *songServices.ResponseSongDTO {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	songRequestObj := &songServices.SongRequest{Title: prmSongTitle}

	res, err := connection.GetSong(ctx, songRequestObj)
	if err != nil {
		fmt.Printf("Error calling gRPC: %v", err)
		return nil
	}
	return res
}

func getStreamingSong(client streamingServices.AudioServiceClient, prmSong *songServices.ResponseSongDTO) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	songTitle := prmSong.SongObj.Title + " - " + prmSong.SongObj.Artist + ".mp3"

	stream, err := client.GetStreamingSong(ctx, &streamingServices.SongRequest{Title: songTitle})
	if err != nil {
		log.Fatal(err)
	}

	reader, writer := io.Pipe()
	canalSincronizacion := make(chan struct{})
	userInputChan := make(chan string, 1)
	playbackDone := make(chan bool, 1)

	// Iniciar reproducción en goroutines
	go utilities.DecodeAndPlay(reader, canalSincronizacion)
	go func() {
		utilities.ReciveSong(stream, writer, canalSincronizacion)
		playbackDone <- true
	}()

	// Mostrar menú y esperar entrada del usuario
	for {
		Views.ShowSongPlayMenu(prmSong, true)

		// Goroutine para leer entrada del usuario
		go func() {
			input := read("Opción: ")
			userInputChan <- input
		}()

		select {
		case input := <-userInputChan:
			if input == "1" {
				Views.ColorStringPrint("Deteniendo reproducción...\n", "yellow", false)
				cancel()
				writer.Close()
				return true
			} else {
				fmt.Println("-> ERROR: Opcion no valida")
			}
		case <-playbackDone:
			fmt.Println("\nReproducción completada.")
			return false
		}
	}
}
