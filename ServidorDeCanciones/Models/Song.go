package models

type Genre struct {
	ID   int32
	NAME string
}

type Song struct {
	ID       int32
	TITLE    string
	ARTIST   string
	YEAR     int32
	DURATION string
	GENRE    Genre
}
