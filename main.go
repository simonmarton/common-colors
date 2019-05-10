package main

import (
	"github.com/simonmarton/common-colors/server"
)

func main() {
	h := ProcessHandler{}

	server.Initialize(h)
}
