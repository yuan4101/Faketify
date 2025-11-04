package capafachada

import (
	"fmt"
	dtos "localServer/grpc-songsServer/Models"
	services "localServer/grpc-songsServer/Services"
	capaaccesoadatos "localServer/grpc-songsServer/capaAccesoADatos"
	componnteconexioncola "localServer/grpc-songsServer/componnteConexionCola"
)

type FachadaAlmacenamiento struct {
	repo         *capaaccesoadatos.RepositorioCanciones
	conexionCola *componnteconexioncola.RabbitPublisher
}

// Constructor de la fachada
func NuevaFachadaAlmacenamiento() *FachadaAlmacenamiento {
	fmt.Println("🔧 Inicializando fachada de almacenamiento...")
	repo := capaaccesoadatos.GetRepositorioCanciones()
	conexionCola, err := componnteconexioncola.NewRabbitPublisher()
	if err != nil {
		fmt.Println("❌ Error al conectar con RabbitMQ:", err)
		conexionCola = nil
	}

	return &FachadaAlmacenamiento{
		repo:         repo,
		conexionCola: conexionCola,
	}
}

func (thisF *FachadaAlmacenamiento) GuardarCancion(objCancion dtos.CancionAlmacenarDTOInput, data []byte, songsArr *[]dtos.Song, genresArr []dtos.Genre) error {
	// Publicar notificación a RabbitMQ
	thisF.conexionCola.PublicarNotificacion(componnteconexioncola.NotificacionCancion{
		Titulo:   objCancion.Titulo,
		Artista:  objCancion.Artista,
		Genero:   objCancion.Genero,
		Año:      objCancion.Año,
		Idioma:   objCancion.Idioma,
		Duracion: objCancion.Duracion,
		Mensaje:  "Nueva canción almacenada: " + objCancion.Titulo + " de " + objCancion.Artista,
	})

	// ✅ CAMBIO: Usar SaveSongFile en lugar de solo guardar el archivo
	// Esta función guarda el archivo Y agrega la canción al array en memoria
	services.SaveSongFile(data, &objCancion, songsArr, genresArr)

	return nil
}
