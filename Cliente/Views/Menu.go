package Views

import (
	"fmt"
	util "localClient/grpc-client/Utilities"
	"localServer/grpc-songServer/songServices"
	"strconv"
)

func ShowMenu(prmGenres *songServices.ResponseGenresDTO, prmSongs *songServices.ResponseSongsDTO, prmGenre string, prmSong *songServices.ResponseSongDTO, prmPlay bool) string {
	util.ColorStringPrint("\n\t Faketify \t\n\n", "green", true)

	genres := prmGenres.GetGenresObjArr()
	songs := prmSongs.GetSongsObjArr()
	if len(genres) > 0 {
		var lastIndex int = 0
		for i, genre := range genres {
			fmt.Print("  ")
			util.ColorIntPrint(i+1, "yellow", false)
			util.ColorStringPrint(". ", "yellow", false)
			util.ColorStringPrint(genre.GetName()+"\n", "white", true)
			lastIndex = i + 2
		}
		fmt.Print("  ")
		util.ColorIntPrint(lastIndex, "red", true)
		util.ColorStringPrint(". Atras", "red", true)
		return strconv.Itoa(lastIndex)
	} else if len(songs) > 0 {
		var lastIndex int = 0
		util.ColorStringPrint("Genero: ", "blue", false)
		util.ColorStringPrint(prmGenre+"\n\n", "white", true)
		for i, song := range songs {
			fmt.Print("  ")
			util.ColorIntPrint(i+1, "yellow", false)
			util.ColorStringPrint(". ", "yellow", false)
			util.ColorStringPrint(song.Artist+" - "+song.Title+"\n", "white", true)
			lastIndex = i + 2
		}
		fmt.Print("  ")
		util.ColorIntPrint(lastIndex, "red", true)
		util.ColorStringPrint(". Atras", "red", true)
		return strconv.Itoa(lastIndex)
	} else if prmSong != nil {
		if prmPlay {
			util.ColorStringPrint("Cancion: ", "blue", false)
			util.ColorStringPrint(prmSong.SongObj.Artist+" - "+prmSong.SongObj.Title+"\n\n", "white", true)
			util.ColorStringPrint("\t Reproduciendo cancion... \n\n", "green", true)
			util.ColorStringPrint("  1. Ir atras y detener la reproduccion", "red", true)
			return "1"
		} else {
			util.ColorStringPrint("Cancion: ", "blue", false)
			util.ColorStringPrint(prmSong.SongObj.Artist+" - "+prmSong.SongObj.Title+"\n\n", "white", true)

			util.ColorStringPrint("\t• Titulo: ", "blue", false)
			util.ColorStringPrint(prmSong.SongObj.Title+"\n", "white", true)
			util.ColorStringPrint("\t• Artista: ", "blue", false)
			util.ColorStringPrint(prmSong.SongObj.Artist+"\n", "white", true)
			util.ColorStringPrint("\t• Año de lanzamiento: ", "blue", false)
			util.ColorIntPrint(int(prmSong.SongObj.Year), "white", true)
			util.ColorStringPrint("\n\t• Duracion: ", "blue", false)
			util.ColorStringPrint(prmSong.SongObj.Duration+"\n\n", "white", true)

			util.ColorStringPrint("  1. Reproducir\n", "green", true)
			util.ColorStringPrint("  2. Atras", "red", true)
			return "2"
		}
	} else {
		util.ColorStringPrint("  1. Ver Generos\n", "green", true)
		util.ColorStringPrint("  2. Salir", "red", true)
		return "2"
	}
}

func ShowMainMenu() string {
	return ShowMenu(nil, nil, "", nil, false)
}

func ShowGenresMenu(prmGenres *songServices.ResponseGenresDTO) string {
	return ShowMenu(prmGenres, nil, "", nil, false)
}

func ShowSongsMenu(prmGenre string, prmSongs *songServices.ResponseSongsDTO) string {
	return ShowMenu(nil, prmSongs, prmGenre, nil, false)
}

func ShowSongMenu(prmSong *songServices.ResponseSongDTO) string {
	return ShowMenu(nil, nil, "", prmSong, false)
}

func ShowSongPlayMenu(prmSong *songServices.ResponseSongDTO, prmPlay bool) string {
	return ShowMenu(nil, nil, "", prmSong, true)
}
