package Views

import (
	"fmt"
	"localServer/grpc-songServer/songServices"
	"strconv"
)

func ShowMenu(prmGenres *songServices.ResponseGenresDTO, prmSongs *songServices.ResponseSongsDTO, prmGenre string, prmSong *songServices.ResponseSongDTO, prmPlay bool) string {
	ColorStringPrint("\n\t Faketify \t\n\n", "green", true)

	genres := prmGenres.GetGenresObjArr()
	songs := prmSongs.GetSongsObjArr()
	if len(genres) > 0 {
		var lastIndex int = 0
		for i, genre := range genres {
			fmt.Print("  ")
			ColorIntPrint(i+1, "yellow", false)
			ColorStringPrint(". ", "yellow", false)
			ColorStringPrint(genre.GetName()+"\n", "white", true)
			lastIndex = i + 2
		}
		fmt.Print("  ")
		ColorIntPrint(lastIndex, "red", true)
		ColorStringPrint(". Atras", "red", true)
		return strconv.Itoa(lastIndex)
	} else if len(songs) > 0 {
		var lastIndex int = 0
		ColorStringPrint("Genero: ", "blue", false)
		ColorStringPrint(prmGenre+"\n\n", "white", true)
		for i, song := range songs {
			fmt.Print("  ")
			ColorIntPrint(i+1, "yellow", false)
			ColorStringPrint(". ", "yellow", false)
			ColorStringPrint(song.Artist+" - "+song.Title+"\n", "white", true)
			lastIndex = i + 2
		}
		fmt.Print("  ")
		ColorIntPrint(lastIndex, "red", true)
		ColorStringPrint(". Atras", "red", true)
		return strconv.Itoa(lastIndex)
	} else if prmSong != nil {
		if prmPlay {
			ColorStringPrint("Cancion: ", "blue", false)
			ColorStringPrint(prmSong.SongObj.Artist+" - "+prmSong.SongObj.Title+"\n\n", "white", true)
			ColorStringPrint("\t Reproduciendo cancion... \n\n\n\n", "green", true)
			ColorStringPrint("  1. Ir atras y detener la reproduccion\n", "red", true)
			return "1"
		} else {
			ColorStringPrint("Cancion: ", "blue", false)
			ColorStringPrint(prmSong.SongObj.Artist+" - "+prmSong.SongObj.Title+"\n\n", "white", true)

			ColorStringPrint("\t• Titulo: ", "blue", false)
			ColorStringPrint(prmSong.SongObj.Title+"\n", "white", true)
			ColorStringPrint("\t• Artista: ", "blue", false)
			ColorStringPrint(prmSong.SongObj.Artist+"\n", "white", true)
			ColorStringPrint("\t• Año de lanzamiento: ", "blue", false)
			ColorIntPrint(int(prmSong.SongObj.Year), "white", true)
			ColorStringPrint("\n\t• Duracion: ", "blue", false)
			ColorStringPrint(prmSong.SongObj.Duration+"\n\n", "white", true)

			ColorStringPrint("  1. Reproducir\n", "green", true)
			ColorStringPrint("  2. Atras", "red", true)
			return "2"
		}
	} else {
		ColorStringPrint("  1. Ver Generos\n", "green", true)
		ColorStringPrint("  2. Salir", "red", true)
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
