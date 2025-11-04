package dtos

type ReproduccionDTOInput struct {
	Usuario string `json:"usuario"`
	Genero  string `json:"genero"`
	Artista string `json:"artista"`
	Titulo  string `json:"titulo"`
	Cliente string `json:"cliente"`
	Idioma  string `json:"idioma"`
}
