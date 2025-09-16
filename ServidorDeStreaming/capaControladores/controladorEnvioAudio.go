package capacontroladores

import (
	capafachadaservices "servidorDeStreaming.local/grpc-servidorDeStreaming/capaFachadaServices"
	pb "servidorDeStreaming.local/grpc-servidorDeStreaming/serviciosCancion"
)

type ControladorServidor struct {
	pb.UnimplementedAudioServiceServer
}

// Implementación del procedimiento remoto
func (s *ControladorServidor) EnviarCancionMedianteStream(req *pb.PeticionDTO, stream pb.AudioService_EnviarCancionMedianteStreamServer) error {

	// Usamos la fachada directamente
	return capafachadaservices.StreamAudioFile(
		req.Titulo,
		func(data []byte) error {
			return stream.Send(&pb.FragmentoCancion{Data: data})
		})
}
