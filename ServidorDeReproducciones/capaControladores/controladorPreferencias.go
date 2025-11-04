package controlador

import (
	"encoding/json"
	"fmt"
	"net/http"
	dtos "tendencias/capaFachadaServices/DTOs"
	capafachada "tendencias/capaFachadaServices/fachada"
)

type ControladorTendencias struct {
	fachada *capafachada.FachadaTendencias
}

// Constructor del Controlador
func NuevoControladorTendencias() *ControladorTendencias {
	return &ControladorTendencias{
		fachada: capafachada.NuevaFachadaTendencias(),
	}
}

// Servicio REST POST que recibe una reproducción en formato JSON
func (c *ControladorTendencias) RegistrarReproduccionHandler(w http.ResponseWriter, r *http.Request) {
	var dto dtos.ReproduccionDTOInput

	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		http.Error(w, "Error al leer el cuerpo de la petición", http.StatusBadRequest)
		return
	}

	c.fachada.RegistrarReproduccion(dto.Usuario, dto.Genero, dto.Artista, dto.Titulo, dto.Cliente, dto.Idioma)

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Reproducción registrada correctamente")
}

// SErvicio REST GET que devuelve todas las reproducciones en formato JSON
func (c *ControladorTendencias) ListarReproduccionesHandler(w http.ResponseWriter, r *http.Request) {
	repros := c.fachada.ObtenerReproducciones()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(repros)
}
