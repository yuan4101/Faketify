package utilities

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"localServer/grpc-streamingServer/streamingServices"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

// DecodeAndPlay decodifica y reproduce un stream MP3 desde un reader.
// Inicializa el speaker con la tasa de muestreo del archivo y reproduce el audio.
// Al finalizar la reproducción, cierra el canal de sincronización para señalizar.
func DecodeAndPlay(reader io.Reader, canalSincronizacion chan struct{}) {
	streamer, format, err := mp3.Decode(io.NopCloser(reader))
	if err != nil {
		log.Fatalf("error decodificando MP3: %v", err)
	}
	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/2))

	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		close(canalSincronizacion)
	})))
}

// ReciveSong recibe fragmentos de audio via streaming y los escribe en un pipe.
// Lee secuencialmente del stream gRPC y escribe los datos en el writer.
// Espera en el canal de sincronización hasta que termine la reproducción.
func ReciveSong(stream streamingServices.AudioService_GetStreamingSongClient, writer *io.PipeWriter, canalSincronizacion chan struct{}) {
	noFragmento := 0
	for {
		fragmento, err := stream.Recv()
		if err == io.EOF {
			//fmt.Println("Cancion recibida completa.")
			writer.Close()
			break
		}
		if err != nil {
			log.Fatalf("Error recibiendo chunk: %v", err)
		}
		noFragmento++
		//fmt.Printf("\n Fragmento #%d recibido (%d bytes) reproduciendo ...", noFragmento, len(fragmento.Data))

		if _, err := writer.Write(fragmento.Data); err != nil {
			//log.Printf("Error escribiendo en pipe: %v", err)
			break
		}
	}

	// Esperar a que la reproducción termine
	<-canalSincronizacion
	//fmt.Println("Reproducción terminada.")
}

// Read muestra un mensaje y lee una entrada del usuario desde la consola.
// Elimina espacios y saltos de línea alrededor de la entrada.
// Retorna el texto ingresado por el usuario como string.
func Read(prmMessage string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("\n%s", prmMessage)
	varReaded, _ := reader.ReadString('\n')
	varReaded = strings.TrimSpace(varReaded)
	return varReaded
}

// cutePrint imprime texto o números con colores y formato ANSI.
// Parámetros: valor a imprimir, color y opción de negrita.
// Uso interno para las funciones públicas de impresión.
func cutePrint(prmInt int, prmString string, prmColor string, prmBold bool) {
	if prmInt == -1 {
		fmt.Printf("\033[%s%sm%s\033[0m", bold(prmBold), color(prmColor), prmString)
	} else {
		fmt.Printf("\033[%s%sm%d\033[0m", bold(prmBold), color(prmColor), prmInt)
	}
}

// bold retorna el código ANSI para formato negrita.
// Retorna "1;" si es true, cadena vacía si es false.
func bold(prmBold bool) string {
	if prmBold {
		return "1;"
	}
	return ""
}

// color retorna el código ANSI para el color especificado.
// Soporta: white(1), yellow(33), red(31), green(32), blue(34).
func color(prmColor string) string {
	switch prmColor {
	case "white":
		return "1"
	case "yellow":
		return "33"
	case "red":
		return "31"
	case "green":
		return "32"
	case "blue":
		return "34"
	default:
		return "1"
	}
}

// ColorIntPrint imprime un número entero con color y formato.
// Ejemplo: ColorIntPrint(42, "green", true) → número 42 en verde negrita.
func ColorIntPrint(prmInt int, prmColor string, prmBold bool) {
	cutePrint(prmInt, "", prmColor, prmBold)
}

// ColorStringPrint imprime un string con color y formato.
// Ejemplo: ColorStringPrint("Éxito", "green", false) → texto en verde.
func ColorStringPrint(prmString string, prmColor string, prmBold bool) {
	cutePrint(-1, prmString, prmColor, prmBold)
}
