package models

type ResponseAllGenresDTO struct {
	GENRES_ARR []Genre
	CODE       int32
	MESSAGE    string
}

type ResponseSongsByGenreDTO struct {
	SONGS_ARR []Song
	CODE      int32
	MESSAGE   string
}

type ResponseSongDTO struct {
	SONG_OBJ Song
	CODE     int32
	MESSAGE  string
}
