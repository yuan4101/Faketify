package main

import (
	"fmt"
	"net/http"
	controlador "tendencias/capaControladores"
)

func main() {
	ctrl := controlador.NuevoControladorTendencias()

	http.HandleFunc("/tendencias/reproduccion", ctrl.RegistrarReproduccionHandler)
	http.HandleFunc("/tendencias/listarReproducciones", ctrl.ListarReproduccionesHandler)

	fmt.Println("Servicio de Tendencias escuchando en el puerto 5000...")
	if err := http.ListenAndServe(":5000", nil); err != nil {
		fmt.Println("Error iniciando el servidor", err)
	}
}
