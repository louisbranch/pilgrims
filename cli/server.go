package cli

import (
	"log"
	"net"
)

type message int

const (
	ping message = iota
	pong
)

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

	/*








				    __
				 __/2 \__
				/1 \__/	 \
				\__/10\__/
				   \__/

					____
				 /    \      /
				/      \____/
				\      /    \
				 \____/ 10   \
				      \ brick/
							 \____/


		     		    _____
		           /     \
		          /       \
		    ,----s         )----.
		   /      \       /      \
		  /        C_____/        \
		  \        /     \        /
		   \      /       \      /
		    )----(         )----(
		   /      \       /      \
		  /        \_____/        \
		  \        /     \        /
		   \      /       \      /
		    `----(         )----'
		          \       /
		           \_____/


			build road 10 9
			build city 10 9 6

	*/

	return nil
}
