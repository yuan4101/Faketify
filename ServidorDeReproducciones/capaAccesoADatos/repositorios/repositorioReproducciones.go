package capaaccesoadatos

import (
	"localServer/grpc-playbackServer/capaAccesoADatos/entitys"
	"log"
	"sync"
	"time"
)

type RepositorioReproducciones struct {
	mu             sync.Mutex
	reproducciones []entitys.ReproduccionEntity
}

// Instancia única del repositorio (patrón SIngleton)
var (
	instancia *RepositorioReproducciones
	once      sync.Once
)

// Crear o devolver la única instancia
func GetRepositorio() *RepositorioReproducciones {
	once.Do(func() {
		instancia = &RepositorioReproducciones{}
	})
	return instancia
}

func (r *RepositorioReproducciones) AgregarReproduccion(usuario, genero, artista, titulo, cliente, idioma string) {
	r.mu.Lock()
	defer r.mu.Unlock()

	reproduccion := entitys.ReproduccionEntity{
		Usuario:   usuario,
		Genero:    genero,
		Artista:   artista,
		Titulo:    titulo,
		Cliente:   cliente,
		Idioma:    idioma,
		FechaHora: time.Now().Format("2006-01-02 15:04:05"),
	}

	r.reproducciones = append(r.reproducciones, reproduccion)
	log.Printf("- CLIENT: %s | POST: %s - %s | USER: %s\n", cliente, titulo, artista, usuario)
}

func (r *RepositorioReproducciones) ListarReproducciones() []entitys.ReproduccionEntity {
	log.Printf("- GET: Listar Reproducciones\n")
	return r.reproducciones
}
