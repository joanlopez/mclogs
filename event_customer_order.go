package mclogs

import (
	"fmt"
	"time"

	"github.com/brianvoe/gofakeit"
)

type CustomerOrder struct {
	Name        string
	HasAccount  bool
	TotalAmount float64
	At          time.Time
}

func NewCustomerOrder(at time.Time) Event {
	return CustomerOrder{
		Name:        gofakeit.Name(),
		HasAccount:  hasAccountPicker.Pick().(bool),
		TotalAmount: gofakeit.Price(0.35, 150),
		At:          at,
	}
}

func (co CustomerOrder) Format(f Format) string {
	var log string
	switch f {
	case FormatJson:
		log = `{"at": "%s", "name": "%s", "has_account": %t, "total_amount": %.2f, "event": "%s"}`
	case FormatLogfmt:
		log = `at=%s name="%s" has_account=%t total_amount=%.2f event="%s"`
	}

	const event = "New customer order."
	return fmt.Sprintf(log, co.At.Format(time.RFC3339), co.Name, co.HasAccount, co.TotalAmount, event)
}
