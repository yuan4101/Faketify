package main

import (
	"context"
	"fmt"
	controlador "localServer/grpc-songsServer/capaControladores"
	"log"
	"net"
	"net/http"

	models "localServer/grpc-songsServer/Models"
	services "localServer/grpc-songsServer/Services"
	"localServer/grpc-songsServer/songServices"

	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
)

var songsArr []models.Song
var genresArr []models.Genre

type songsServer struct {
	songServices.UnimplementedSongServiceServer
}

// GetSong implementa el servicio gRPC para buscar una canción por título.
// Recibe una solicitud con el título y retorna los metadatos de la canción.
// Loggea la consulta del cliente y convierte la respuesta al formato gRPC.
func (s *songsServer) GetSong(ctx context.Context, req *songServices.SongRequest) (*songServices.ResponseSongDTO, error) {
	title := req.GetTitle()

	resp := services.GetSong(title, songsArr)

	var response songServices.ResponseSongDTO
	response.Code = resp.CODE
	response.Message = resp.MESSAGE
	if p, ok := peer.FromContext(ctx); ok {
		log.Printf("| CLIENT: %s | GET: %s | CODE: %d | %s", p.Addr.String(), title, response.Code, response.Message)
	}

	if resp.CODE == 200 {
		response.SongObj = new(songServices.Song)
		response.SongObj.Id = resp.SONG_OBJ.ID
		response.SongObj.Title = resp.SONG_OBJ.TITLE
		response.SongObj.Artist = resp.SONG_OBJ.ARTIST
		response.SongObj.Year = resp.SONG_OBJ.YEAR
		response.SongObj.Duration = resp.SONG_OBJ.DURATION
		response.SongObj.Language = resp.SONG_OBJ.LANGUAGE

		response.SongObj.Genre = new(songServices.Genre)
		response.SongObj.Genre.Id = resp.SONG_OBJ.GENRE.ID
		response.SongObj.Genre.Name = resp.SONG_OBJ.GENRE.NAME
	}

	return &response, nil
}

// GetGenres implementa el servicio gRPC para obtener todos los géneros musicales.
// Retorna la lista completa de géneros en formato protobuf.
// Loggea la consulta del cliente y convierte los datos al formato gRPC.
func (s *songsServer) GetGenres(ctx context.Context, req *songServices.Empty) (*songServices.ResponseGenresDTO, error) {
	resp := services.GetGenres(genresArr)

	var response songServices.ResponseGenresDTO
	response.Code = resp.CODE
	response.Message = resp.MESSAGE
	if p, ok := peer.FromContext(ctx); ok {
		log.Printf("| CLIENT: %s | GET: Genres | CODE: %d | %s", p.Addr.String(), response.Code, response.Message)
	}

	if resp.CODE == 200 {
		for _, genre := range genresArr {
			protoGenre := &songServices.Genre{
				Id:   genre.ID,
				Name: genre.NAME,
			}
			response.GenresObjArr = append(response.GenresObjArr, protoGenre)
		}
	}

	return &response, nil
}

// GetSongsByGenre implementa el servicio gRPC para obtener canciones por género.
// Filtra el catálogo por nombre de género y retorna las canciones correspondientes.
// Convierte los resultados al formato protobuf y loggea la consulta del cliente.
func (s *songsServer) GetSongsByGenre(ctx context.Context, req *songServices.SongsByGenreRequest) (*songServices.ResponseSongsDTO, error) {
	genre := req.GetGenreName()

	resp := services.GetSongsByGenre(genre, songsArr)

	var response songServices.ResponseSongsDTO
	response.Code = resp.CODE
	response.Message = resp.MESSAGE
	if p, ok := peer.FromContext(ctx); ok {
		log.Printf("| CLIENT: %s | GET: %s songs | CODE: %d | %s", p.Addr.String(), genre, response.Code, response.Message)
	}

	if resp.CODE == 200 {
		for _, song := range resp.SONGS_ARR {
			protoSong := &songServices.Song{
				Id:       song.ID,
				Title:    song.TITLE,
				Artist:   song.ARTIST,
				Year:     song.YEAR,
				Duration: song.DURATION,
				Genre: &songServices.Genre{
					Id:   song.GENRE.ID,
					Name: song.GENRE.NAME,
				},
			}
			response.SongsObjArr = append(response.SongsObjArr, protoSong)
		}
	}

	return &response, nil
}

// SaveSong implementa el servicio gRPC para guardar una canción
func (s *songsServer) SaveSong(ctx context.Context, req *songServices.SaveSongRequest) (*songServices.SaveSongResponse, error) {
	if p, ok := peer.FromContext(ctx); ok {
		log.Printf("| CLIENT: %s | SAVING: %s by %s", p.Addr.String(), req.Title, req.Artist)
	}

	// Crear DTO de entrada
	cancionDTO := &models.CancionAlmacenarDTOInput{
		Titulo:   req.Title,
		Artista:  req.Artist,
		Año:      req.Year,
		Duracion: req.Duration,
		Idioma:   req.Language,
		Genero:   req.Genre,
	}

	// Guardar archivo Y agregar a songsArr
	songID := services.SaveSongFile(req.FileContent, cancionDTO, &songsArr, genresArr)

	response := &songServices.SaveSongResponse{
		Code:    200,
		Message: "Song saved successfully",
		SongId:  songID,
	}

	return response, nil
}

func main() {
	listener, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatalf("Failed to open port 50053: %v", err)
	}

	// ✅ Cargar metadatos de canciones PRIMERO
	services.LoadSongsMetadata(&songsArr, &genresArr)

	go func() {
		// ✅ Pasar los arrays al controlador
		ctrl := controlador.NuevoControladorAlmacenamientoCanciones(&songsArr, genresArr)
		http.HandleFunc("/canciones/almacenamiento", ctrl.AlmacenarAudioCancion)
		fmt.Println("✅ Servicio de Almacenamiento escuchando en el puerto 5001...")
		if err := http.ListenAndServe(":5001", nil); err != nil {
			fmt.Println("Error iniciando el servidor:", err)
		}
	}()

	// Crear servidor gRPC
	grpcServer := grpc.NewServer()
	songServices.RegisterSongServiceServer(grpcServer, &songsServer{})

	// Iniciar el servidor
	log.Println("Songs gRPC server listening on port 50053")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to start gRPC server: %v", err)
	}
}
