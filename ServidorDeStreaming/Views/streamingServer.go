package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"net"

	services "localServer/grpc-streamingServer/Services"
	comunicacionservidorpreferencias "localServer/grpc-streamingServer/capaComunicacionExterna/comunicacionservidorPreferencias"
	"localServer/grpc-streamingServer/capalogger"
	"localServer/grpc-streamingServer/streamingServices"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
)

type streamingServer struct {
	streamingServices.UnimplementedAudioServiceServer
	logger *capalogger.Logger
}

func NewControladorServidor(logger *capalogger.Logger) *streamingServer {
	return &streamingServer{
		logger: logger,
	}
}

// GetStreamingSong implementa el servicio gRPC para streaming de audio.
// Recibe una solicitud de canción y transmite los datos en paquetes por el stream.
// Utiliza el servicio subyacente para leer el archivo y enviar fragmentos.
func (s *streamingServer) GetStreamingSong(req *streamingServices.SongRequest, stream streamingServices.AudioService_GetStreamingSongServer) error {
	var clientAddr string
	if p, ok := peer.FromContext(stream.Context()); ok {
		clientAddr = p.Addr.String()
		log.Printf("-> CLIENT: %s | GET: %s ", clientAddr, req.GetTitle())
	}

	// NUEVO: Extraer metadata del contexto
	md, ok := metadata.FromIncomingContext(stream.Context())
	if !ok {
		log.Println("Advertencia: No se recibieron metadatos")
	}

	// Obtener y DECODIFICAR valores de los metadatos
	var userID, genre, artist, songTitle, language string // ← NUEVO: language

	if ok {
		if userIDSlice := md.Get("user-id"); len(userIDSlice) > 0 {
			userID = userIDSlice[0]
		}

		// Decodificar base64 para género
		if genreSlice := md.Get("genre"); len(genreSlice) > 0 {
			decoded, err := base64.StdEncoding.DecodeString(genreSlice[0])
			if err != nil {
				genre = genreSlice[0] // Si no es base64, usar como está
			} else {
				genre = string(decoded)
			}
		}

		// Decodificar base64 para artista
		if artistSlice := md.Get("artist"); len(artistSlice) > 0 {
			decoded, err := base64.StdEncoding.DecodeString(artistSlice[0])
			if err != nil {
				artist = artistSlice[0]
			} else {
				artist = string(decoded)
			}
		}

		// Decodificar base64 para título
		if songTitleSlice := md.Get("song-title"); len(songTitleSlice) > 0 {
			decoded, err := base64.StdEncoding.DecodeString(songTitleSlice[0])
			if err != nil {
				songTitle = songTitleSlice[0]
			} else {
				songTitle = string(decoded)
			}
		}

		// NUEVO: Decodificar base64 para idioma
		if languageSlice := md.Get("language"); len(languageSlice) > 0 {
			decoded, err := base64.StdEncoding.DecodeString(languageSlice[0])
			if err != nil {
				language = languageSlice[0]
			} else {
				language = string(decoded)
			}
		}
	}

	// Log de los datos recibidos (ahora decodificados) - ACTUALIZADO
	log.Printf("📊 Datos recibidos - Usuario: %s | Género: %s | Artista: %s | Canción: %s | Idioma: %s | IP: %s",
		userID, genre, artist, songTitle, language, clientAddr)

	// Invocación a operación asincrónica que envía datos a otro proceso
	go func() {
		err := comunicacionservidorpreferencias.RegistrarReproduccionEnTendencias(
			userID,
			genre,
			artist,
			songTitle,
			clientAddr,
			language, // ← NUEVO: pasar idioma
		)

		if err != nil {
			log.Println("Error registrando tendencia:", err)
		}
	}()

	return services.GetStreamingSong(
		req.Title,
		func(data []byte) error {
			return stream.Send(&streamingServices.SongPacket{Data: data})
		})
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Error escuchando en el puerto: %v", err)
	}

	grpcServer := grpc.NewServer()

	// Se registra el controlador que ofrece el procedimiento remoto
	streamingServices.RegisterAudioServiceServer(grpcServer, NewControladorServidor(capalogger.CrearUnicaInstanciaDelLogger()))

	fmt.Println("Servidor gRPC escuchando en :50051...")

	// Iniciar el servidor
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Error al iniciar servidor gRPC: %v", err)
	}
}
