package main

import (
	"fmt"
	"log"
	"net"

	services "localServer/grpc-streamingServer/Services"
	"localServer/grpc-streamingServer/streamingServices"

	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
)

type streamingServer struct {
	streamingServices.UnimplementedAudioServiceServer
}

// Implementación del procedimiento remoto
func (s *streamingServer) GetStreamingSong(req *streamingServices.SongRequest, stream streamingServices.AudioService_GetStreamingSongServer) error {

	if p, ok := peer.FromContext(stream.Context()); ok {
		log.Printf("-> CLIENT: %s | GET: %s ", p.Addr.String(), req.GetTitle())
	}
	// Usamos la fachada directamente
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
	streamingServices.RegisterAudioServiceServer(grpcServer, &streamingServer{})

	fmt.Println("Servidor gRPC escuchando en :50051...")

	// Iniciar el servidor
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Error al iniciar servidor gRPC: %v", err)
	}
}
