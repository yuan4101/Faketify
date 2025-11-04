package models

type CancionAlmacenarDTOInput struct {
	Titulo   string `json:"titulo"`
	Artista  string `json:"artista"`
	Genero   string `json:"genero"`
	Año      string `json:"año"`
	Duracion string `json:"duracion"`
	Idioma   string `json:"idioma"`
}
