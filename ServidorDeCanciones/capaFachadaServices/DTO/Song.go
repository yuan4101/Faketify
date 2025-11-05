package dto

type Genre struct {
	ID   int32
	NAME string
}

type Song struct {
	ID       int32
	TITLE    string
	ARTIST   string
	YEAR     string
	DURATION string
	LANGUAGE string
	GENRE    Genre
}
