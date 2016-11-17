package main

import (
	"bufio"
	"log"
	"os"

	"github.com/luizbranco/pilgrims/cli"
)

func main() {
	c, err := cli.NewClient()
	if err != nil {
		log.Fatal(err)
	}

	r := bufio.NewReader(os.Stdin)

	for {
		txt, err := r.ReadSlice('\n')
		if err != nil {
			log.Fatal(err)
		}
		c.Message(txt)
	}
}
