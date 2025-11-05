package main

import (
	"context"
	"fmt"
	controlador "localServer/grpc-songsServer/capaControladores"
	"log"
	"net"
	"net/http"

	dto "localServer/grpc-songsServer/capaFachadaServices/DTO"
	services "localServer/grpc-songsServer/capaFachadaServices/services"
	"localServer/grpc-songsServer/songServices"

	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
)

var songsArr []dto.Song
var genresArr []dto.Genre

type songsServer struct {
	songServices.UnimplementedSongServiceServer
}

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

func (s *songsServer) SaveSong(ctx context.Context, req *songServices.SaveSongRequest) (*songServices.SaveSongResponse, error) {
	if p, ok := peer.FromContext(ctx); ok {
		log.Printf("| CLIENT: %s | SAVING: %s by %s", p.Addr.String(), req.Title, req.Artist)
	}

	cancionDTO := &dto.CancionAlmacenarDTOInput{
		Titulo:   req.Title,
		Artista:  req.Artist,
		Año:      req.Year,
		Duracion: req.Duration,
		Idioma:   req.Language,
		Genero:   req.Genre,
	}

	songID := services.SaveSongFile(req.FileContent, cancionDTO, &songsArr, genresArr)

	response := &songServices.SaveSongResponse{
		Code:    200,
		Message: "Song saved successfully",
		SongId:  songID,
	}

	return response, nil
}

func main() {
	puerto := ":50053"
	listener, err := net.Listen("tcp", puerto)
	if err != nil {
		log.Fatalf("Failed to open port %s: %v", puerto, err)
	}

	services.LoadSongsMetadata(&songsArr, &genresArr)

	go func() {
		ctrl := controlador.NuevoControladorAlmacenamientoCanciones(&songsArr, genresArr)
		http.HandleFunc("/canciones/almacenamiento", ctrl.AlmacenarAudioCancion)
		puerto := ":5001"
		fmt.Printf("-> Servicio de Almacenamiento escuchando en el puerto %s\n", puerto)
		if err := http.ListenAndServe(puerto, nil); err != nil {
			fmt.Println("Error iniciando el servidor:", err)
		}
	}()

	grpcServer := grpc.NewServer()
	songServices.RegisterSongServiceServer(grpcServer, &songsServer{})

	log.Printf("\n\t\t----- SERVIDOR DE CANCIONES (Go/gRPC) [%s] -----\n", puerto)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to start gRPC server: %v", err)
	}
}
