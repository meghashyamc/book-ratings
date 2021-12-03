package logs

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
)

func Setup() *os.File {

	log.SetFormatter(&log.JSONFormatter{})
	f, err := os.OpenFile("book-ratings-log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		fmt.Printf("error opening file: %v", err)
	}
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(f)
	return f
}
