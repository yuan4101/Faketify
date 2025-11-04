package capaaccesoadatos

import (
	"fmt"
	"sync"
	. "tendencias/capaAccesoADatos/entitys"
	"time"
)

type RepositorioReproducciones struct {
	mu             sync.Mutex
	reproducciones []ReproduccionEntity
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

	reproduccion := ReproduccionEntity{
		Usuario:   usuario,
		Genero:    genero,
		Artista:   artista,
		Titulo:    titulo,
		Cliente:   cliente,
		Idioma:    idioma,
		FechaHora: time.Now().Format("2006-01-02 15:04:05"),
	}

	r.reproducciones = append(r.reproducciones, reproduccion)
	fmt.Printf("Reproducción almacenada: %+v\n", reproduccion)
	r.mostrarReproducciones()
}

func (r *RepositorioReproducciones) ListarReproducciones() []ReproduccionEntity {
	return r.reproducciones
}

func (r *RepositorioReproducciones) mostrarReproducciones() {
	fmt.Println("=== Reproducciones almacenadas ===")
	for i := 0; i < len(r.reproducciones); i++ {
		fmt.Println("Usuario: " + r.reproducciones[i].Usuario)
		fmt.Println("Genero: " + r.reproducciones[i].Genero)
		fmt.Println("Artista: " + r.reproducciones[i].Artista)
		fmt.Println("Titulo: " + r.reproducciones[i].Titulo)
		fmt.Println("Cliente: " + r.reproducciones[i].Cliente)
		fmt.Println("Idioma: " + r.reproducciones[i].Idioma)
		fmt.Println("Fecha y Hora: " + r.reproducciones[i].FechaHora)
	}
}
