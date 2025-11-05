package capafachada

import (
	"localServer/grpc-playbackServer/capaAccesoADatos/entitys"
	capaaccesoadatos "localServer/grpc-playbackServer/capaAccesoADatos/repositorios"
)

type FachadaReproducciones struct {
	repo *capaaccesoadatos.RepositorioReproducciones
}

// Constructor de la fachada
func NuevaFachadaReproducciones() *FachadaReproducciones {
	return &FachadaReproducciones{
		repo: capaaccesoadatos.GetRepositorio(),
	}
}

func (f *FachadaReproducciones) RegistrarReproduccion(usuario, genero, artista, titulo, cliente, idioma string) {
	f.repo.AgregarReproduccion(usuario, genero, artista, titulo, cliente, idioma)
}

func (f *FachadaReproducciones) ObtenerReproducciones() []entitys.ReproduccionEntity {
	return f.repo.ListarReproducciones()
}
