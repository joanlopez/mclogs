package mclogs

import (
	"fmt"
	"github.com/brianvoe/gofakeit"
	"time"

	"github.com/mroth/weightedrand"
)

var (
	hasAccountPicker   *weightedrand.Chooser
	satisfactionPicker *weightedrand.Chooser
)

func init() {
	hasAccountPicker, _ = weightedrand.NewChooser(
		weightedrand.NewChoice(false, 70),
		weightedrand.NewChoice(true, 30),
	)

	satisfactionPicker, _ = weightedrand.NewChooser(
		weightedrand.NewChoice(0, 5),
		weightedrand.NewChoice(1, 5),
		weightedrand.NewChoice(2, 5),
		weightedrand.NewChoice(3, 5),
		weightedrand.NewChoice(4, 5),
		weightedrand.NewChoice(5, 10),
		weightedrand.NewChoice(6, 15),
		weightedrand.NewChoice(7, 30),
		weightedrand.NewChoice(8, 10),
		weightedrand.NewChoice(9, 5),
		weightedrand.NewChoice(10, 5),
	)
}

type CustomerSatisfaction struct {
	Name              string
	HasAccount        bool
	SatisfactionLevel int
	At                time.Time
}

func NewCustomerSatisfaction(at time.Time) Event {
	return CustomerSatisfaction{
		Name:              gofakeit.Name(),
		HasAccount:        hasAccountPicker.Pick().(bool),
		SatisfactionLevel: satisfactionPicker.Pick().(int),
		At:                at,
	}
}

func (cs CustomerSatisfaction) Format(f Format) string {
	var log string
	switch f {
	case FormatJson:
		log = `{"at": "%s", "name": "%s", "has_account": %t, "satisfaction_level": %d, "event": "%s"}`
	case FormatLogfmt:
		log = `at=%s name="%s" has_account=%t satisfaction_level=%d event="%s"`
	}

	const event = "Customer satisfaction survey submitted."
	return fmt.Sprintf(log, cs.At.Format(time.RFC3339), cs.Name, cs.HasAccount, cs.SatisfactionLevel, event)
}
