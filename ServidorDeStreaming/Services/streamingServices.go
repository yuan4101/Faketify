package services

import (
	"fmt"
	"io"
	"os"
)

func GetStreamingSong(prmTitle string, sendFragmentFunction func([]byte) error) error {
	file, err := os.Open("../Songs/" + prmTitle)
	if err != nil {
		return fmt.Errorf("no se pudo abrir el archivo: %w", err)
	}
	defer file.Close()

	buffer := make([]byte, 64*1024) // 64 KB se envian por fragmento
	fragmento := 0

	for {
		n, err := file.Read(buffer)
		if err == io.EOF {
			//log.Printf("Cancion enviada completamente desde la fachada.")
			break
		}
		if err != nil {
			return fmt.Errorf("error leyendo el archivo: %w", err)
		}
		fragmento++
		//log.Printf("Fragmento #%d leido (%d bytes) y enviando", fragmento, n)

		// Ejecutamos la funcion para enviar el fragmento al cliente
		err = sendFragmentFunction(buffer[:n])
		if err != nil {
			return fmt.Errorf("error enviando fragmento #%d: %w", fragmento, err)
		}
	}

	return nil
}
