package cli

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"unsafe"

	"github.com/luizbranco/pilgrims"
)

func init() {
	var state pilgrims.GameState
	state.Board = pilgrims.NewBoard()
	fmt.Println(unsafe.Sizeof(state))
	fmt.Println(state.Board)

	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, state)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("% x", buf.Bytes())
	fmt.Println(unsafe.Sizeof(buf))

	var s pilgrims.GameState

	err = binary.Read(buf, binary.LittleEndian, &s)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(unsafe.Sizeof(s))
	fmt.Println(s.Board)
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


                 ,---.
                /  a  \
            P  /       \  P
					,---(         )---.
         /  b  \       /  c  \
        /       \     /       \
	 ,---(         )---(         )---.
P /  d  \       /  e  \       /  f  \ P
 /       \     /       \     /       \
(         )---(         )---(         )
 \       /  h  \       /  i  \       /
  \     /       \     /       \     /
   )---(         )---(         )---(
  /  g  \       /  j  \       /  k  \
 /       \     /       \     /       \
(         )---(         )---(         )
 \       /  l  \       /  m  \       /
P \     /       \     /       \     / P
   )---(         )---(         )---(
  /  n  \       /  o  \       /  p  \
 /       \     /       \     /       \
(         )---(         )---(         )
 \       /  q  \       /  r  \       /
  \     /       \     /       \     /
	 '---(         )---(         )---'
        \       /  s  \       /
       P \     /       \     / P
					'---(         )---'
               \       /
                \     /
                 '---'
                   P

build road a b c
build city 10 9 6
`
