package comunicacionservidorreproducciones

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// Estructura que coincide con el DTO del servicio de tendencias
type ReproduccionDTOInput struct {
	Usuario string `json:"usuario"`
	Genero  string `json:"genero"`
	Artista string `json:"artista"`
	Titulo  string `json:"titulo"`
	Cliente string `json:"cliente"`
	Idioma  string `json:"idioma"` // ← NUEVO CAMPO
}

// Función que envía una reproducción al microservicio de tendencias
func RegistrarReproduccion(usuario, genero, artista, titulo, cliente, idioma string) error {
	url := "http://localhost:5000/playback/reproduccion" // endpoint del microservicio de tendencias

	// Crear el cuerpo JSON
	body := ReproduccionDTOInput{
		Usuario: usuario,
		Genero:  genero,
		Artista: artista,
		Titulo:  titulo,
		Cliente: cliente,
		Idioma:  idioma,
	}

	jsonData, err := json.Marshal(body)
	if err != nil {
		return fmt.Errorf("error convirtiendo a JSON: %v", err)
	}

	// Enviar solicitud POST al microservicio de tendencias
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("error al enviar POST a tendencias: %v", err)
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("el servidor de reproducciones respondió con código %d", resp.StatusCode)
	}

	//fmt.Printf("# Reproducción de '%s' registrada para '%s'.\n", titulo, usuario)
	return nil
}
