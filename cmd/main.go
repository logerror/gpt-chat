package main

import (
	"github.com/logerror/gpt-chat/cmd/app"
	"log"
)

func main() {
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
