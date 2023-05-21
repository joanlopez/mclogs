package mclogs

import (
	"time"

	"github.com/mroth/weightedrand"
)

var (
	eventPicker *weightedrand.Chooser
)

func init() {
	eventPicker, _ = weightedrand.NewChooser(
		weightedrand.NewChoice(NewCustomerOrder, 65),
		weightedrand.NewChoice(NewCustomerSatisfaction, 35),
	)
}

type Event interface {
	Format(Format) string
}

func NewEvent(at time.Time) Event {
	return eventPicker.Pick().(func(time.Time) Event)(at)
}
