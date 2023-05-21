package main

import (
	"github.com/joanlopez/mclogs"
)

func main() {
	cfg := configFromFlags()
	if err := mclogs.Run(cfg); err != nil {
		panic(err)
	}
}
