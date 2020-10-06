package main

import (
	"log"

	"github.com/copilot-example-voting-app/results"
)

func main() {
	if err := results.Run(); err != nil {
		log.Fatalf("run vote server: %v\n", err)
	}
}