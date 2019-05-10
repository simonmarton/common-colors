package main

import (
	"github.com/simonmarton/common-colors/calculator"
	"github.com/simonmarton/common-colors/server"
)

func main() {
	h := ProcessHandler{
		calculator: calculator.New(
			calculator.Config{
				IterationCount:       6,
				TransparencyTreshold: 10,
				MinLuminance:         .3,
				MaxLuminance:         .9,
			},
		),
	}

	server.Initialize(h)
}
