package capaaccesoadatos

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"
)

type RepositorioCanciones struct {
	mu sync.Mutex
}

var (
	instancia *RepositorioCanciones
	once      sync.Once
)

// GetRepositorioCanciones aplica patrón Singleton
func GetRepositorioCanciones() *RepositorioCanciones {
	once.Do(func() {
		instancia = &RepositorioCanciones{}
	})
	return instancia
}

func (r *RepositorioCanciones) GuardarCancion(titulo string, genero string, artista string, año string, idioma string, duracion string, data []byte) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	repoPath := "../../RepositorioCanciones"

	err := os.MkdirAll(repoPath, os.ModePerm)
	if err != nil {
		return fmt.Errorf("error creando carpeta: %v", err)
	}

	fileName := fmt.Sprintf("%s_%s_%s.mp3", titulo, genero, artista)
	filePath := filepath.Join(repoPath, fileName)

	err = os.WriteFile(filePath, data, 0644)
	if err != nil {
		return fmt.Errorf("error guardando archivo: %v", err)
	}

	log.Printf("- SAVE: '%s - %s' en '%s'", titulo, artista, filePath)
	return nil
}
