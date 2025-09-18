package services

import (
	model "localServer/grpc-songsServer/Models"
	"regexp"
	"strings"
)

func cleanSongTitle(title string) string {
	cleaned := strings.ToLower(strings.TrimSpace(title))

	// Convierte acentos a letras normales
	replacer := strings.NewReplacer(
		"á", "a", "à", "a", "â", "a", "ä", "a", "ã", "a",
		"é", "e", "è", "e", "ê", "e", "ë", "e",
		"í", "i", "ì", "i", "î", "i", "ï", "i",
		"ó", "o", "ò", "o", "ô", "o", "ö", "o", "õ", "o",
		"ú", "u", "ù", "u", "û", "u", "ü", "u",
		"ñ", "n", "ç", "c",
	)
	cleaned = replacer.Replace(cleaned)

	// Ahora limpia manteniendo solo a-z, 0-9 y &
	re := regexp.MustCompile(`[^a-z0-9&]+`)
	return re.ReplaceAllString(cleaned, "")
}

func LoadSongsMetadata(songsArr *[]model.Song) {
	// Genres
	genresObjArr := []model.Genre{
		{ID: 0, NAME: "Rock"},
		{ID: 1, NAME: "Pop"},
		{ID: 2, NAME: "Jazz"},
		{ID: 3, NAME: "Classical"},
		{ID: 4, NAME: "Hip Hop"},
		{ID: 5, NAME: "Electronic"},
		{ID: 6, NAME: "Salsa"},
		{ID: 7, NAME: "Reggae"},
		{ID: 8, NAME: "Blues"},
		{ID: 9, NAME: "Metal"},
	}

	// Songs
	songsObjArr := []model.Song{
		{ID: 0, TITLE: "Test", ARTIST: "Test Artist", YEAR: 9999, DURATION: "9:99", GENRE: genresObjArr[3]},
		{ID: 1, TITLE: "Bohemian Rhapsody", ARTIST: "Queen", YEAR: 1975, DURATION: "5:54", GENRE: genresObjArr[0]},
	}

	*songsArr = append(*songsArr, songsObjArr...)
}

func GetSong(prmTitle string, songsArr []model.Song) model.ResponseSongDTO {
	var response model.ResponseSongDTO

	cleanSearchTitle := cleanSongTitle(prmTitle)

	for i := 0; i < len(songsArr); i++ {
		cleanSongTitle := cleanSongTitle(songsArr[i].TITLE)
		//log.Print("Comparacion ", i+1, ": ", cleanSearchTitle, " | ", cleanSongTitle, "\n")
		if cleanSongTitle == cleanSearchTitle {
			response.SONG_OBJ = songsArr[i]
			response.CODE = 200
			response.MESSAGE = "Song found"
			return response
		}
	}
	response.CODE = 404
	response.MESSAGE = "Song not found"
	return response
}
