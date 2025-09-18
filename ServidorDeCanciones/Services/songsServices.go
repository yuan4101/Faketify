package services

import (
	model "localServer/grpc-songsServer/Models"
	"strings"
)

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
		{ID: 0, TITLE: "test", ARTIST: "Test artist", YEAR: 1975, DURATION: "9:99", GENRE: genresObjArr[3]},
		{ID: 1, TITLE: "Bohemian Rhapsody", ARTIST: "Queen", YEAR: 1975, DURATION: "5:54", GENRE: genresObjArr[0]},
	}

	*songsArr = append(*songsArr, songsObjArr...)
}

func GetSong(prmTitle string, songsArr []model.Song) model.ResponseSongDTO {
	var response model.ResponseSongDTO

	cleanSearchTitle := strings.ToLower(strings.ReplaceAll(strings.TrimSpace(prmTitle), " ", ""))

	for i := 0; i < len(songsArr); i++ {
		cleanSongTitle := strings.ToLower(strings.ReplaceAll(strings.TrimSpace(songsArr[i].TITLE), " ", ""))
		print(cleanSearchTitle, " | ", cleanSongTitle, "\n")
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
