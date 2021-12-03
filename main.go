package main

import (
	"os"

	"github.com/joho/godotenv"

	"github.com/meghashyamc/book-ratings/logs"
	"github.com/meghashyamc/book-ratings/ui"
)

func main() {

	godotenv.Load()
	logFile := logs.Setup()
	defer logFile.Close()
	screen, err := ui.NewScreen()
	if err != nil {
		os.Exit(1)
	}
	screen.ShowAndRun()

}
