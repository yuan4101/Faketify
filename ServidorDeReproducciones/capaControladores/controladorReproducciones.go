package controlador

import (
	"encoding/json"
	"fmt"
	dtos "localServer/grpc-playbackServer/capaFachadaServices/DTOs"
	capafachada "localServer/grpc-playbackServer/capaFachadaServices/fachada"
	"net/http"
)

type ControladorReproducciones struct {
	fachada *capafachada.FachadaReproducciones
}

// Constructor del Controlador
func NuevoControladorReproducciones() *ControladorReproducciones {
	return &ControladorReproducciones{
		fachada: capafachada.NuevaFachadaReproducciones(),
	}
}

// Servicio REST POST que recibe una reproducción en formato JSON
func (c *ControladorReproducciones) RegistrarReproduccionHandler(w http.ResponseWriter, r *http.Request) {
	var dto dtos.ReproduccionDTOInput

	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		http.Error(w, "Error al leer el cuerpo de la petición", http.StatusBadRequest)
		return
	}

	c.fachada.RegistrarReproduccion(dto.Usuario, dto.Genero, dto.Artista, dto.Titulo, dto.Cliente, dto.Idioma)

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Reproducción registrada correctamente")
}

// Servicio REST GET que devuelve todas las reproducciones en formato JSON
func (c *ControladorReproducciones) ListarReproduccionesHandler(w http.ResponseWriter, r *http.Request) {
	repros := c.fachada.ObtenerReproducciones()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(repros)
}
