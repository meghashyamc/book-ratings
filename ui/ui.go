package ui

import (
	"bufio"
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
)

func ShowWelcome() {
	fmt.Println("Welcome to 'Book Ratings'")

}

func InputBookName() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the name of the book and I'll tell you it's ratings.")
	bookName, err := reader.ReadString('\n')
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("Error in reading book name (input)")
		return "", err
	}
	return bookName, nil

}

func DisplayBookRatings(ratings map[string]float32) {
	fmt.Println("These are the ratings:")
	for k, v := range ratings {
		fmt.Println(k, ":", v)
	}
}
