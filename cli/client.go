package cli

import (
	"log"
	"net"
)

type Client struct {
	remote *net.UDPAddr
	conn   *net.UDPConn
	in     chan []byte
	out    chan []byte
}

func NewClient() (*Client, error) {
	remote := &net.UDPAddr{
		Port: 8080,
		IP:   net.ParseIP("127.0.0.1"),
	}

	conn, err := net.DialUDP("udp", nil, remote)
	if err != nil {
		return nil, err
	}

	c := &Client{
		remote: remote,
		conn:   conn,
		in:     make(chan []byte),
		out:    make(chan []byte),
	}

	go c.loop()

	return c, nil
}

func (c *Client) loop() {
	for {
		select {
		case msg := <-c.in:
			log.Printf("in: %s\n", msg)
		case msg := <-c.out:
			_, err := c.conn.Write(msg)
			if err != nil {
				log.Print(err)
			}
		}
	}
}

func (c *Client) Message(msg []byte) {
	c.out <- msg
}
