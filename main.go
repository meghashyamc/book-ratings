package main

import (
	"os"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"

	"github.com/meghashyamc/book-ratings/ratings"
	"github.com/meghashyamc/book-ratings/ui"
)

func main() {

	godotenv.Load()
	log.SetFormatter(&log.JSONFormatter{})

	ui.ShowWelcome()
	bookName, err := ui.InputBookName()
	if err != nil {
		os.Exit(1)
	}

	ratings, err := ratings.GetAllRatings(bookName)
	if err != nil {
		os.Exit(1)
	}
	ui.DisplayBookRatings(ratings)

}
