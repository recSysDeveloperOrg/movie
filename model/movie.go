package model

import "time"

type Movie struct {
	Id           string
	Title        string
	PicUrl       string
	Introduction string
	Participants string
	ReleaseDate  time.Time
	Language     string
}
