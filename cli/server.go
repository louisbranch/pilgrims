package cli

import (
	"log"
	"net"
	"unsafe"

	"github.com/luizbranco/pilgrims"
)

func init() {
	var state pilgrims.GameState
	state.Board = pilgrims.NewBoard()
	log.Println(unsafe.Sizeof(state))
	log.Println(state.Board)
}

func Listen() error {
	addr := net.UDPAddr{
		Port: 8080,
		IP:   net.ParseIP("127.0.0.1"),
	}

	ln, err := net.ListenUDP("udp", &addr)
	if err != nil {
		return err
	}

	buf := make([]byte, 512)

	for {
		n, remote, err := ln.ReadFromUDP(buf)
		if err != nil {
			log.Printf("connection error: %s", err)
			continue
		}
		log.Printf("%+v: %s", remote, buf[:n])
	}

	return nil
}

const board string = `

                 ,----.
                /      \
          _____/        \_____
         /     \        /     \
        /       \      /       \
  ,----(         )----(         )----.
 /      \       /      \       /      \
/        \_____/        \_____/        \
\        /     \        /     \        /
 \      /       \      /       \      /
  )----(         )----(         )----(
 /      \       /      \       /      \
/        \_____/        \_____/        \
\        /     \        /     \        /
 \      /       \      /       \      /
  )----(         )----(         )----(
 /      \       /      \       /      \
/        \_____/        \_____/        \
\        /     \        /     \        /
 \      /       \      /       \      /
  '----(         )----(         )----'
        \       /      \       /
				 \_____/        \_____/
							 \        /
			          \      /
								 '----'

build road a b c
build city 10 9 6
`
