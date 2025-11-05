package Views

import (
	"fmt"
	util "localClient/grpc-client/Utilities"
	"localServer/grpc-songServer/songServices"
	"strconv"
)

type MenuState string

const (
	StateLogin       MenuState = "LOGIN"
	StateMainMenu    MenuState = "MAIN"
	StateGenres      MenuState = "GENRES"
	StateSongs       MenuState = "SONGS"
	StateSongDetail  MenuState = "SONG_DETAIL"
	StatePlaying     MenuState = "PLAYING"
	StatePreferences MenuState = "PREFERENCES"
	StateExit        MenuState = "EXIT"
)

// Estructura de usuario con contraseña
type User struct {
	Username string
	Password string
}

// Usuarios válidos en memoria con contraseñas
var ValidUsers = []User{
	{Username: "juan", Password: "123"},
	{Username: "camila", Password: "456"},
	{Username: "ramon", Password: "789"},
}

type Controller struct {
	state   MenuState
	handler *ActionHandler

	// Inyección de dependencias - funciones de los servicios
	GetGenresFn        func() *songServices.ResponseGenresDTO
	GetSongsByGenreFn  func(string) *songServices.ResponseSongsDTO
	GetSongFn          func(string) *songServices.ResponseSongDTO
	GetStreamingSongFn func(userID string, song *songServices.ResponseSongDTO) bool
	GetPreferencesFn   func(string)
}

type ActionHandler struct {
	UserID       string
	CurrentGenre string
	CurrentSong  *songServices.ResponseSongDTO
}

func New(
	getGenresFn func() *songServices.ResponseGenresDTO,
	getSongsByGenreFn func(string) *songServices.ResponseSongsDTO,
	getSongFn func(string) *songServices.ResponseSongDTO,
	getStreamingSongFn func(userID string, song *songServices.ResponseSongDTO) bool,
	getPreferencesFn func(string),
) *Controller {
	return &Controller{
		state: StateLogin,
		handler: &ActionHandler{
			UserID: "",
		},
		GetGenresFn:        getGenresFn,
		GetSongsByGenreFn:  getSongsByGenreFn,
		GetSongFn:          getSongFn,
		GetStreamingSongFn: getStreamingSongFn,
		GetPreferencesFn:   getPreferencesFn,
	}
}

func (mc *Controller) Start() {
	for mc.state != StateExit {
		switch mc.state {
		case StateLogin:
			mc.handleLogin()
		case StateMainMenu:
			mc.handleMainMenu()
		case StateGenres:
			mc.handleGenres()
		case StateSongs:
			mc.handleSongs()
		case StateSongDetail:
			mc.handleSongDetail()
		case StatePlaying:
			mc.handlePlaying()
		case StatePreferences:
			mc.handlePreferences()
		}
	}
}

// ========== LOGIN CON USUARIO Y CONTRASEÑA ==========

func (mc *Controller) handleLogin() {
	util.ColorStringPrint("\n\t Faketify \t\n", "green", true)
	util.ColorStringPrint("\n--- INICIO DE SESIÓN ---\n", "blue", true)

	// Pedir usuario
	username := util.Read("Usuario: ")
	if username == "" {
		util.ColorStringPrint("-> Usuario vacío\n", "red", false)
		return
	}

	// Pedir contraseña
	password := util.Read("Contraseña: ")
	if password == "" {
		util.ColorStringPrint("-> Contraseña vacía\n", "red", false)
		return
	}

	// Validar credenciales
	if mc.validateCredentials(username, password) {
		mc.handler.UserID = username
		util.ColorStringPrint(fmt.Sprintf("\n-> Bienvenido, %s\n", mc.handler.UserID), "green", true)
		mc.state = StateMainMenu
	} else {
		util.ColorStringPrint("-> Usuario o contraseña incorrectos\n", "red", false)
	}
}

// Función para validar credenciales
func (mc *Controller) validateCredentials(username, password string) bool {
	for _, user := range ValidUsers {
		if user.Username == username && user.Password == password {
			return true
		}
	}
	return false
}

// ========== MENÚ PRINCIPAL ==========

func (mc *Controller) handleMainMenu() {
	util.ColorStringPrint(fmt.Sprintf("\n--- MENÚ PRINCIPAL (Usuario: %s) ---\n", mc.handler.UserID), "blue", true)

	util.ColorStringPrint("1. ", "yellow", false)
	fmt.Println("Ver géneros disponibles")

	util.ColorStringPrint("2. ", "yellow", false)
	fmt.Println("Ver preferencias")

	util.ColorStringPrint("3. ", "yellow", false)
	fmt.Println("Cerrar sesión")

	option := util.Read("Opción: ")

	switch option {
	case "1":
		mc.state = StateGenres
	case "2":
		mc.state = StatePreferences
	case "3":
		util.ColorStringPrint(fmt.Sprintf("\n-> Sesión cerrada para %s\n", mc.handler.UserID), "yellow", false)
		mc.handler.UserID = ""
		mc.state = StateLogin
	default:
		util.ColorStringPrint("-> Opción no válida\n", "red", false)
	}
}

// ========== GÉNEROS ==========

func (mc *Controller) handleGenres() {
	genres := mc.GetGenresFn()

	if len(genres.GenresObjArr) == 0 {
		util.ColorStringPrint("-> No hay géneros disponibles\n", "red", false)
		mc.state = StateMainMenu
		return
	}

	util.ColorStringPrint("\n--- GÉNEROS DISPONIBLES ---\n", "blue", true)

	// Imprimir géneros
	for i, genre := range genres.GenresObjArr {
		util.ColorIntPrint(i+1, "yellow", false)
		util.ColorStringPrint(". ", "yellow", false)
		fmt.Println(genre.GetName())
	}

	// Opción de volver
	backOption := len(genres.GenresObjArr) + 1
	util.ColorIntPrint(backOption, "yellow", false)
	util.ColorStringPrint(". ", "yellow", false)
	fmt.Println("Volver")

	option := util.Read("Selecciona género: ")

	genreIdx, err := strconv.Atoi(option)
	if err != nil {
		util.ColorStringPrint("-> Opción no válida\n", "red", false)
		return
	}

	// Volver al menú principal
	if genreIdx == backOption {
		mc.state = StateMainMenu
		return
	}

	// Seleccionar género
	if genreIdx > 0 && genreIdx <= len(genres.GenresObjArr) {
		mc.handler.CurrentGenre = genres.GenresObjArr[genreIdx-1].GetName()
		util.ColorStringPrint(fmt.Sprintf("\n-> Seleccionado: %s\n", mc.handler.CurrentGenre), "green", true)
		mc.state = StateSongs
	} else {
		util.ColorStringPrint("-> Opción no válida\n", "red", false)
	}
}

// ========== CANCIONES ==========

func (mc *Controller) handleSongs() {
	songs := mc.GetSongsByGenreFn(mc.handler.CurrentGenre)

	if len(songs.SongsObjArr) == 0 {
		util.ColorStringPrint(fmt.Sprintf("-> No hay canciones para: %s\n", mc.handler.CurrentGenre), "red", false)
		mc.state = StateGenres
		return
	}

	util.ColorStringPrint(fmt.Sprintf("\n--- CANCIONES DE %s ---\n", mc.handler.CurrentGenre), "blue", true)

	// Imprimir canciones
	for i, song := range songs.SongsObjArr {
		util.ColorIntPrint(i+1, "yellow", false)
		util.ColorStringPrint(". ", "yellow", false)
		fmt.Printf("%s - %s\n", song.GetTitle(), song.GetArtist())
	}

	// Opción de volver
	backOption := len(songs.SongsObjArr) + 1
	util.ColorIntPrint(backOption, "yellow", false)
	util.ColorStringPrint(". ", "yellow", false)
	fmt.Println("Volver")

	option := util.Read("Selecciona canción: ")

	songIdx, err := strconv.Atoi(option)
	if err != nil {
		util.ColorStringPrint("-> Opción no válida\n", "red", false)
		return
	}

	// Volver a géneros
	if songIdx == backOption {
		mc.state = StateGenres
		return
	}

	// Seleccionar canción
	if songIdx > 0 && songIdx <= len(songs.SongsObjArr) {
		song := mc.GetSongFn(songs.SongsObjArr[songIdx-1].GetTitle())
		if song == nil {
			util.ColorStringPrint("-> No se pudo cargar la canción\n", "red", false)
			return
		}
		mc.handler.CurrentSong = song
		util.ColorStringPrint(fmt.Sprintf("\n✓ Canción: %s\n", song.SongObj.Title), "green", true)
		mc.state = StateSongDetail
	} else {
		util.ColorStringPrint("-> Opción no válida\n", "red", false)
	}
}

// ========== DETALLES DE CANCIÓN ==========

func (mc *Controller) handleSongDetail() {
	util.ColorStringPrint("\n--- DETALLES DE CANCIÓN ---\n", "blue", true)
	util.ColorStringPrint("-> Título: ", "green", false)
	fmt.Println(mc.handler.CurrentSong.SongObj.Title)
	util.ColorStringPrint("-> Artista: ", "green", false)
	fmt.Println(mc.handler.CurrentSong.SongObj.Artist)
	util.ColorStringPrint("-> Género: ", "green", false)
	fmt.Println(mc.handler.CurrentSong.SongObj.Genre.GetName())
	util.ColorStringPrint("-> Año: ", "green", false)
	fmt.Println(mc.handler.CurrentSong.SongObj.Year)
	util.ColorStringPrint("-> Idioma: ", "green", false)
	fmt.Println(mc.handler.CurrentSong.SongObj.Language)
	util.ColorStringPrint(" Duracion: ", "green", false)
	fmt.Println(mc.handler.CurrentSong.SongObj.Duration)

	util.ColorStringPrint("\n1. ", "yellow", false)
	fmt.Println("Reproducir")
	util.ColorStringPrint("2. ", "yellow", false)
	fmt.Println("Volver")
	option := util.Read("Opción: ")

	switch option {
	case "1":
		mc.state = StatePlaying
	case "2":
		mc.state = StateSongs
	default:
		util.ColorStringPrint("-> Opción no válida\n", "red", false)
	}
}

// ========== REPRODUCCIÓN ==========

func (mc *Controller) handlePlaying() {
	util.ColorStringPrint("\n-> Iniciando reproducción...\n", "green", false)

	// Pasar el usuario dinámicamente a la función de streaming
	success := mc.GetStreamingSongFn(mc.handler.UserID, mc.handler.CurrentSong)

	if success {
		util.ColorStringPrint("\n-> Reproducción completada\n", "green", false)
		mc.state = StateSongDetail
	} else {
		util.ColorStringPrint("\n-> Error en la reproducción\n", "red", false)
		mc.state = StateSongs
	}
}

// ========== PREFERENCIAS ==========

func (mc *Controller) handlePreferences() {
	util.ColorStringPrint("\n--- PREFERENCIAS DE USUARIO ---\n", "blue", true)

	util.ColorStringPrint(fmt.Sprintf("\n-> Obteniendo preferencias para: %s...\n", mc.handler.UserID), "green", false)
	mc.GetPreferencesFn(mc.handler.UserID)

	mc.state = StateMainMenu
}
