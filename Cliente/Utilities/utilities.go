package utilities

import (
	"fmt"
	"io"
	"log"
	"time"

	"localServer/grpc-streamingServer/streamingServices"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

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

func cutePrint(prmInt int, prmString string, prmColor string, prmBold bool) {
	if prmInt == -1 {
		fmt.Printf("\033[%s%sm%s\033[0m", bold(prmBold), color(prmColor), prmString)
	} else {
		fmt.Printf("\033[%s%sm%d\033[0m", bold(prmBold), color(prmColor), prmInt)
	}
}

func bold(prmBold bool) string {
	if prmBold {
		return "1;"
	}
	return ""
}

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

func ColorIntPrint(prmInt int, prmColor string, prmBold bool) {
	cutePrint(prmInt, "", prmColor, prmBold)
}

func ColorStringPrint(prmString string, prmColor string, prmBold bool) {
	cutePrint(-1, prmString, prmColor, prmBold)
}
