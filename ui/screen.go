package ui

import (
	"github.com/gdamore/tcell/v2"
	log "github.com/sirupsen/logrus"
)

var (
	usedBoxes    = []box{goodreadsBox, libraryThingBox}
	usedButtons  = []button{getRatingsButton}
	usedMessages = []message{welcomeMessage}
)

type screen struct {
	screen   tcell.Screen
	boxes    []box
	buttons  []button
	messages []message
}

func NewScreen() (*screen, error) {

	sc := &screen{}

	if err := sc.init(); err != nil {
		return nil, err
	}
	sc.boxes = usedBoxes
	sc.buttons = usedButtons
	sc.messages = usedMessages
	return sc, nil

}

func (sc *screen) init() error {

	s, err := tcell.NewScreen()
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error(initFailedErr)
		return err
	}
	if err := s.Init(); err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error(initFailedErr)
		return err
	}
	screenStyle := tcell.StyleDefault.Foreground(defaultForeground).Background(defaultBackground)
	s.SetStyle(screenStyle)
	s.Clear()

	sc.screen = s
	return nil
}

func (sc *screen) ShowAndRun() {

}
