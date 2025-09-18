package main

import (
	"context"
	"log"
	"net"

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

func (s *songsServer) GetSong(ctx context.Context, req *songServices.SongRequest) (*songServices.ResponseSongDTO, error) {
	title := req.GetTitle()

	resp := services.GetSong(title, songsArr)

	var response songServices.ResponseSongDTO
	response.Code = resp.CODE
	response.Message = resp.MESSAGE
	if p, ok := peer.FromContext(ctx); ok {
		log.Printf("-> Client(%s) | GET: %s | %d | %s", p.Addr.String(), title, response.Code, response.Message)
	}

	if resp.CODE == 200 {
		response.SongObj = new(songServices.Song)
		response.SongObj.Id = resp.SONG_OBJ.ID
		response.SongObj.Title = resp.SONG_OBJ.TITLE
		response.SongObj.Artist = resp.SONG_OBJ.ARTIST
		response.SongObj.Year = resp.SONG_OBJ.YEAR
		response.SongObj.Duration = resp.SONG_OBJ.DURATION

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
		log.Printf("-> Client(%s) | GET: Genres | %d | %s", p.Addr.String(), response.Code, response.Message)
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

func main() {
	listener, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatalf("Failed to open port 50053: %v", err)
	}

	// Crear servidor gRPC
	grpcServer := grpc.NewServer()

	// Rrgistrar el servicio
	songServices.RegisterSongServiceServer(grpcServer, &songsServer{})

	// Cargar metadatos de canciones
	services.LoadSongsMetadata(&songsArr, &genresArr)

	// Iniciar el servidor
	log.Println("Songs gRPC server listening on port 50053")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to start gRPC server: %v", err)
	}
}
