package capafachada

import (
	"tendencias/capaAccesoADatos/entitys"
	capaaccesoadatos "tendencias/capaAccesoADatos/repositorios"
)

type FachadaTendencias struct {
	repo *capaaccesoadatos.RepositorioReproducciones
}

// Constructor de la fachada
func NuevaFachadaTendencias() *FachadaTendencias {
	return &FachadaTendencias{
		repo: capaaccesoadatos.GetRepositorio(),
	}
}

func (f *FachadaTendencias) RegistrarReproduccion(usuario, genero, artista, titulo, cliente, idioma string) {
	f.repo.AgregarReproduccion(usuario, genero, artista, titulo, cliente, idioma)
}

func (f *FachadaTendencias) ObtenerReproducciones() []entitys.ReproduccionEntity {
	return f.repo.ListarReproducciones()
}
