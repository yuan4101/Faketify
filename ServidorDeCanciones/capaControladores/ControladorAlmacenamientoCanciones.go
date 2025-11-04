package controlador

import (
	"fmt"
	"io"
	dtos "localServer/grpc-songsServer/Models"
	capafachada "localServer/grpc-songsServer/capaFachadaServices/fachada"
	"net/http"
)

type ControladorAlmacenamientoCanciones struct {
	fachada   *capafachada.FachadaAlmacenamiento
	songsArr  *[]dtos.Song
	genresArr []dtos.Genre
}

// Constructor del Controlador
func NuevoControladorAlmacenamientoCanciones(songsArr *[]dtos.Song, genresArr []dtos.Genre) *ControladorAlmacenamientoCanciones {
	return &ControladorAlmacenamientoCanciones{
		fachada:   capafachada.NuevaFachadaAlmacenamiento(),
		songsArr:  songsArr,
		genresArr: genresArr,
	}
}

func (thisC *ControladorAlmacenamientoCanciones) AlmacenarAudioCancion(w http.ResponseWriter, r *http.Request) {
	fmt.Print("📝 Almacenando canción...\n")

	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	r.ParseMultipartForm(50 << 20)
	file, _, err := r.FormFile("archivo")
	if err != nil {
		http.Error(w, "Error leyendo el archivo", http.StatusBadRequest)
		return
	}
	defer file.Close()

	data, _ := io.ReadAll(file)

	// Leer los campos del DTO
	dto := dtos.CancionAlmacenarDTOInput{
		Titulo:   r.FormValue("titulo"),
		Artista:  r.FormValue("artista"),
		Genero:   r.FormValue("genero"),
		Año:      r.FormValue("año"),
		Idioma:   r.FormValue("idioma"),
		Duracion: r.FormValue("duracion"),
	}

	// ✅ Pasar los arrays a la fachada
	thisC.fachada.GuardarCancion(dto, data, thisC.songsArr, thisC.genresArr)

	// Respuesta exitosa
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"mensaje":"Canción guardada exitosamente","titulo":"%s","artista":"%s"}`, dto.Titulo, dto.Artista)
}
