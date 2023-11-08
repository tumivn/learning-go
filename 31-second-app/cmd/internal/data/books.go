package data

import "time"

type Book struct {
	ID        int
	Title     string
	Author    string
	Price     float32
	CreatedAt time.Time
	UpdatedAt time.Time
	Pages     int
	Genres    []string
	Rating    float32
	Version   int
}
