package capalogger

import (
	"fmt"
	"os"
	"sync"
	"time"
)

type Logger struct {
	archivo *os.File
	mu      sync.Mutex
}

var (
	objInstancia *Logger
	once         sync.Once
)

func abrirArchivoDeLog(nombre string) *os.File {
	file, err := os.OpenFile(nombre, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(fmt.Sprintf("error creando archivo de log %s: %v", nombre, err))
	}
	fmt.Printf("Abriendo archivo de log...")
	return file
}

func inicializadorLogger() {
	file := abrirArchivoDeLog("canciones.log")
	objInstancia = &Logger{archivo: file}
}

func CrearUnicaInstanciaDelLogger() *Logger {
	fmt.Println("Creando instancia unica del Logger...")
	once.Do(inicializadorLogger)
	return objInstancia
}

func (thisL *Logger) AlmacenarSolicitud(titulo, cliente string) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	msg := fmt.Sprintf("{\"titulo\":\"%s\", \"cliente\":\"%s\", \"fecha\":\"%s\"}", titulo, cliente, timestamp)

	thisL.mu.Lock()
	simularEspera()
	fmt.Println("Almacenando en archivo de log: ", msg)
	defer thisL.mu.Unlock()
	thisL.archivo.WriteString(msg + "\n")
}

func simularEspera() {
	time.Sleep(5 * time.Second)
}

func (l *Logger) CerrarArchivo() {
	fmt.Println("Cerrando archivo de log...")
	l.archivo.Close()
}
