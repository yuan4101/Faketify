package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"strings"
	"time"

	util "localClient/grpc-client/Utilities"
	"localClient/grpc-client/Views"
	pb "localClient/grpc-client/preferenciasGrpc"
	"localServer/grpc-songServer/songServices"
	"localServer/grpc-streamingServer/streamingServices"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

var (
	songConn           *grpc.ClientConn
	streamingConn      *grpc.ClientConn
	preferenciasConn   *grpc.ClientConn
	songClient         songServices.SongServiceClient
	streamingClient    streamingServices.AudioServiceClient
	preferenciasClient pb.PreferenciasServiceClient
)

func main() {
	initializeConnections()
	defer closeConnections()

	// Crear el controlador del menú con inyección de dependencias
	controller := Views.New(
		getGenres,
		getSongsByGenre,
		getSong,
		getStreamingSong,
		obtenerPreferencias,
	)

	controller.Start()
}

func initializeConnections() {
	var err error

	// Conexión al servidor de canciones
	songConn, err = grpc.Dial("localhost:50053", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("Error conectando a servidor de canciones: %v\n", err)
		return
	}
	songClient = songServices.NewSongServiceClient(songConn)

	// Conexión al servidor de streaming
	conexionStreaming := "localhost:50051"
	streamingConn, err = grpc.Dial(conexionStreaming, grpc.WithInsecure())
	if err != nil {
		fmt.Printf("Error conectando a servidor de streaming: %v\n", err)
		return
	}
	streamingClient = streamingServices.NewAudioServiceClient(streamingConn)
	fmt.Printf("-> Conectado a servidor de streaming en %s\n", conexionStreaming)

	// Conexión al servidor de preferencias
	conexionPreferencias := "localhost:50052"
	preferenciasConn, err = grpc.Dial(conexionPreferencias, grpc.WithInsecure())
	if err != nil {
		fmt.Printf("Error conectando a servidor de preferencias: %v\n", err)
		return
	}
	preferenciasClient = pb.NewPreferenciasServiceClient(preferenciasConn)
	fmt.Printf("-> Conectado a servidor de preferencias en %s", conexionPreferencias)
}

func closeConnections() {
	if songConn != nil {
		songConn.Close()
	}
	if streamingConn != nil {
		streamingConn.Close()
	}
	if preferenciasConn != nil {
		preferenciasConn.Close()
	}
}

func getGenres() *songServices.ResponseGenresDTO {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res, err := songClient.GetGenres(ctx, &songServices.Empty{})
	if err != nil {
		fmt.Printf("Error obteniendo géneros: %v\n", err)
		return nil
	}

	return res
}

func getSongsByGenre(prmGenre string) *songServices.ResponseSongsDTO {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	genreRequestObj := &songServices.SongsByGenreRequest{GenreName: prmGenre}
	res, err := songClient.GetSongsByGenre(ctx, genreRequestObj)
	if err != nil {
		fmt.Printf("Error obteniendo canciones: %v\n", err)
		return nil
	}

	return res
}

func getSong(prmSongTitle string) *songServices.ResponseSongDTO {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	songRequestObj := &songServices.SongRequest{Title: prmSongTitle}
	res, err := songClient.GetSong(ctx, songRequestObj)
	if err != nil {
		fmt.Printf("Error obteniendo canción: %v\n", err)
		return nil
	}

	return res
}

// ACTUALIZADO: Agregar metadato "language"
func getStreamingSong(userID string, prmSong *songServices.ResponseSongDTO) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	// Crear metadatos con el idioma
	md := metadata.Pairs(
		"user-id", userID,
		"genre", base64.StdEncoding.EncodeToString([]byte(prmSong.SongObj.Genre.GetName())),
		"artist", base64.StdEncoding.EncodeToString([]byte(prmSong.SongObj.Artist)),
		"song-title", base64.StdEncoding.EncodeToString([]byte(prmSong.SongObj.Title)),
		"language", base64.StdEncoding.EncodeToString([]byte(prmSong.SongObj.Language)), // ← NUEVO: Pasar idioma
	)

	ctx = metadata.NewOutgoingContext(ctx, md)

	fullSongTitle := prmSong.SongObj.Title + " - " + prmSong.SongObj.Artist + ".mp3"
	stream, err := streamingClient.GetStreamingSong(ctx, &streamingServices.SongRequest{Title: fullSongTitle})
	if err != nil {
		log.Fatal(err)
	}

	reader, writer := io.Pipe()
	canalSincronizacion := make(chan struct{})
	userInputChan := make(chan string, 1)
	playbackDone := make(chan bool, 1)

	go util.DecodeAndPlay(reader, canalSincronizacion)
	go func() {
		util.ReciveSong(stream, writer, canalSincronizacion)
		playbackDone <- true
	}()

	for {
		Views.ShowSongPlayMenu(prmSong, true)
		go func() {
			input := util.Read("Opción: ")
			userInputChan <- input
		}()

		select {
		case input := <-userInputChan:
			if input == "1" {
				util.ColorStringPrint("Deteniendo reproducción...\n", "yellow", false)
				cancel()
				writer.Close()
				return true
			} else {
				fmt.Println("-> ERROR: Opción no válida")
			}
		case <-playbackDone:
			util.ColorStringPrint("\nReproducción completada.\n", "yellow", false)
			return false
		}
	}
}

func obtenerPreferencias(nombreUsuario string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	response, err := preferenciasClient.ObtenerPreferenciasUsuario(ctx, &pb.PreferenciaRequest{
		NombreUsuario: nombreUsuario,
	})

	if err != nil {
		fmt.Printf("Error obteniendo preferencias: %v\n", err)
		return
	}

	mostrarPreferenciasFormato(response)
}

// ACTUALIZADO: Mostrar también idiomas
func mostrarPreferenciasFormato(preferencias *pb.PreferenciaResponse) {
	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println(" PREFERENCIAS DEL USUARIO")
	fmt.Println(strings.Repeat("=", 60))

	// Géneros favoritos
	fmt.Println("\n GÉNEROS FAVORITOS:")
	fmt.Println(strings.Repeat("-", 58))
	if len(preferencias.PreferenciasGeneros) == 0 {
		fmt.Println(" Sin datos de géneros")
	} else {
		for i, genero := range preferencias.PreferenciasGeneros {
			fmt.Printf(" %d. %-40s: %d reproducciones\n",
				i+1,
				genero.NombreGenero,
				genero.NumeroPreferencias,
			)
		}
	}

	// Artistas favoritos
	fmt.Println("\n ARTISTAS FAVORITOS:")
	fmt.Println(strings.Repeat("-", 58))
	if len(preferencias.PreferenciasArtistas) == 0 {
		fmt.Println(" Sin datos de artistas")
	} else {
		for i, artista := range preferencias.PreferenciasArtistas {
			fmt.Printf(" %d. %-40s: %d reproducciones\n",
				i+1,
				artista.NombreArtista,
				artista.NumeroPreferencias,
			)
		}
	}

	// Idiomas favoritos
	fmt.Println("\n IDIOMAS FAVORITOS:")
	fmt.Println(strings.Repeat("-", 58))
	if len(preferencias.PreferenciasIdiomas) == 0 {
		fmt.Println(" Sin datos de idiomas")
	} else {
		for i, idioma := range preferencias.PreferenciasIdiomas {
			fmt.Printf(" %d. %-40s: %d reproducciones\n",
				i+1,
				idioma.NombreIdioma,
				idioma.NumeroPreferencias,
			)
		}
	}

	fmt.Println("\n" + strings.Repeat("=", 60) + "\n")
}
