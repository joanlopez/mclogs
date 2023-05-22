package main

import (
	"github.com/joanlopez/mclogs/cmd/mclogs/bootstrap"
)

func main() {
	if err := bootstrap.Run(); err != nil {
		panic(err)
	}
}
