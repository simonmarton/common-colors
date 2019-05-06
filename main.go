package main

import (
	"github.com/simonmarton/common-colors/calculator"
	"github.com/simonmarton/common-colors/server"
)

func main() {
	h := ProcessHandler{
		calculator: calculator.New(
			calculator.Config{
				IterationCount:       2,
				TransparencyTreshold: 10,
			},
		),
	}

	server.Initialize(h)
}
