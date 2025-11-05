package main

import (
	"fmt"
	controlador "localServer/grpc-playbackServer/capaControladores"
	"net/http"
)

func main() {
	ctrl := controlador.NuevoControladorReproducciones()

	http.HandleFunc("/playback/reproduccion", ctrl.RegistrarReproduccionHandler)
	http.HandleFunc("/playback/listarReproducciones", ctrl.ListarReproduccionesHandler)

	puerto := ":5000"
	fmt.Printf("\n\t\t----- SERVIDOR DE REPRODUCCIONES (Go/REST) [%s] -----\n", puerto)
	if err := http.ListenAndServe(puerto, nil); err != nil {
		fmt.Println("Error iniciando el servidor", err)
	}
}
