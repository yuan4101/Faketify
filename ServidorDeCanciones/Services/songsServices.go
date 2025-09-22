package services

import (
	model "localServer/grpc-songsServer/Models"
)

func LoadSongsMetadata(songsArr *[]model.Song, genresArr *[]model.Genre) {
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
		// Rock
		{ID: 0, TITLE: "Californication", ARTIST: "Red Hot Chili Peppers", YEAR: 1999, DURATION: "5:21 Minutos", GENRE: genresObjArr[0]},
		{ID: 1, TITLE: "De Música Ligera", ARTIST: "Soda Stereo", YEAR: 1990, DURATION: "3:33 Minutos", GENRE: genresObjArr[0]},

		// Pop
		{ID: 2, TITLE: "As It Was", ARTIST: "Harry Styles", YEAR: 2022, DURATION: "2:43 Minutos", GENRE: genresObjArr[1]},
		{ID: 3, TITLE: "Flowers", ARTIST: "Miley Cyrus", YEAR: 2023, DURATION: "3:17 Minutos", GENRE: genresObjArr[1]},

		// Jazz
		{ID: 4, TITLE: "Take Five", ARTIST: "Dave Brubeck", YEAR: 1959, DURATION: "5:24 Minutos", GENRE: genresObjArr[2]},
		{ID: 5, TITLE: "What A Wonderful World", ARTIST: "Louis Armstrong", YEAR: 1967, DURATION: "2:16 Minutos", GENRE: genresObjArr[2]},

		// Classical
		{ID: 6, TITLE: "Four Seasons Winter", ARTIST: "Vivaldi", YEAR: 1725, DURATION: "9:26 Minutos", GENRE: genresObjArr[3]},
		{ID: 7, TITLE: "Moonlight Sonata", ARTIST: "Beethoven", YEAR: 1801, DURATION: "6:14 Minutos", GENRE: genresObjArr[3]},

		// Hip Hop
		{ID: 8, TITLE: "In Da Club", ARTIST: "50 Cent", YEAR: 2003, DURATION: "3:13 Minutos", GENRE: genresObjArr[4]},
		{ID: 9, TITLE: "Still D.R.E.", ARTIST: "Dr. Dre ft. Snoop Dogg", YEAR: 1999, DURATION: "4:30 Minutos", GENRE: genresObjArr[4]},

		// Electronic
		{ID: 10, TITLE: "Faded", ARTIST: "Alan Walker", YEAR: 2015, DURATION: "3:32 Minutos", GENRE: genresObjArr[5]},
		{ID: 11, TITLE: "Titanium", ARTIST: "David Guetta ft. Sia", YEAR: 2011, DURATION: "4:27 Minutos", GENRE: genresObjArr[5]},

		// Salsa
		{ID: 12, TITLE: "Idilio", ARTIST: "Willie Colon", YEAR: 1993, DURATION: "5:08 Minutos", GENRE: genresObjArr[6]},
		{ID: 13, TITLE: "Rebelion", ARTIST: "Joe Arroyo", YEAR: 1986, DURATION: "6:18 Minutos", GENRE: genresObjArr[6]},

		// Reggae
		{ID: 14, TITLE: "Could you be loved", ARTIST: "Bob Marley", YEAR: 1980, DURATION: "3:50 Minutos", GENRE: genresObjArr[7]},
		{ID: 15, TITLE: "Sweat (A La La Long)", ARTIST: "Inner Circle", YEAR: 1992, DURATION: "4:27 Minutos", GENRE: genresObjArr[7]},

		// Blues
		{ID: 16, TITLE: "Sweet Home Chicago", ARTIST: "The Blues Brothers", YEAR: 1980, DURATION: "7:47 Minutos", GENRE: genresObjArr[8]},
		{ID: 17, TITLE: "The thrill is gone", ARTIST: "B.B. King", YEAR: 1969, DURATION: "5:23 Minutos", GENRE: genresObjArr[8]},

		// Metal
		{ID: 18, TITLE: "Toxicity", ARTIST: "System of a Down", YEAR: 2001, DURATION: "3:39 Minutos", GENRE: genresObjArr[9]},
		{ID: 19, TITLE: "Warriors Of Time", ARTIST: "BlackTide", YEAR: 2008, DURATION: "4:12 Minutos", GENRE: genresObjArr[9]},
	}

	*songsArr = append(*songsArr, songsObjArr...)
	*genresArr = append(*genresArr, genresObjArr...)
}

func GetSong(prmTitle string, songsArr []model.Song) model.ResponseSongDTO {
	var response model.ResponseSongDTO

	for i := 0; i < len(songsArr); i++ {
		if prmTitle == songsArr[i].TITLE {
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

func GetGenres(genresArr []model.Genre) model.ResponseAllGenresDTO {
	var response model.ResponseAllGenresDTO

	if len(genresArr) > 0 {
		response.GENRES_ARR = genresArr
		response.CODE = 200
		response.MESSAGE = "Genres found"
		return response
	}

	response.CODE = 404
	response.MESSAGE = "Genres not found"
	return response
}

func GetSongsByGenre(prmGenre string, songsArr []model.Song) model.ResponseSongsByGenreDTO {
	var response model.ResponseSongsByGenreDTO
	var filteredSongs []model.Song
	var genreName string

	for i := 0; i < len(songsArr); i++ {
		if prmGenre == songsArr[i].GENRE.NAME {
			genreName = songsArr[i].GENRE.NAME
			filteredSongs = append(filteredSongs, songsArr[i])
		}
	}

	if len(filteredSongs) > 0 {
		response.SONGS_ARR = filteredSongs
		response.CODE = 200
		response.MESSAGE = genreName + " Songs found"
		return response
	}

	response.CODE = 404
	response.MESSAGE = genreName + " Songs not found"
	return response
}
