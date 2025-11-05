package capafachada

import (
	"fmt"
	capaaccesoadatos "localServer/grpc-songsServer/capaAccesoADatos"
	dtos "localServer/grpc-songsServer/capaFachadaServices/DTO"
	services "localServer/grpc-songsServer/capaFachadaServices/services"
	componnteconexioncola "localServer/grpc-songsServer/componnteConexionCola"
)

type FachadaAlmacenamiento struct {
	repo         *capaaccesoadatos.RepositorioCanciones
	conexionCola *componnteconexioncola.RabbitPublisher
}

func NuevaFachadaAlmacenamiento() *FachadaAlmacenamiento {
	repo := capaaccesoadatos.GetRepositorioCanciones()
	conexionCola, err := componnteconexioncola.NewRabbitPublisher()
	if err != nil {
		fmt.Println("Error al conectar con RabbitMQ:", err)
		conexionCola = nil
	}

	return &FachadaAlmacenamiento{
		repo:         repo,
		conexionCola: conexionCola,
	}
}

func (thisF *FachadaAlmacenamiento) GuardarCancion(objCancion dtos.CancionAlmacenarDTOInput, data []byte, songsArr *[]dtos.Song, genresArr []dtos.Genre) error {
	thisF.conexionCola.PublicarNotificacion(componnteconexioncola.NotificacionCancion{
		Titulo:   objCancion.Titulo,
		Artista:  objCancion.Artista,
		Genero:   objCancion.Genero,
		Año:      objCancion.Año,
		Idioma:   objCancion.Idioma,
		Duracion: objCancion.Duracion,
		Mensaje:  "Nueva canción almacenada: " + objCancion.Titulo + " de " + objCancion.Artista,
	})

	services.SaveSongFile(data, &objCancion, songsArr, genresArr)

	return nil
}
