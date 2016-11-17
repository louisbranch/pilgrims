package main

import (
	"log"

	"github.com/luizbranco/pilgrims/cli"
)

func main() {
	err := cli.Listen()
	if err != nil {
		log.Fatal(err)
	}
}
