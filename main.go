package main

import (
	"fmt"
	"log"
	"net"
)

type Curry struct {
	addr   string
	ln     net.Listener
	quitch chan struct{}
}

func New(addr string) *Curry {
	return &Curry{addr: addr, quitch: make(chan struct{})}
}

func (c *Curry) Start() error {
	ln, err := net.Listen("tcp", c.addr)
	if err != nil {
		fmt.Printf("Error listening on address: %v", c.addr)
		fmt.Println(err)
		return err
	}
	defer ln.Close()
	c.ln = ln
	go c.Accept()
	<-c.quitch
	return nil
}

func (c *Curry) Accept() {
	for {
		conn, err := c.ln.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go c.Read(conn)
	}
}

func (c *Curry) Read(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 1024)
	_, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading from connection:", err)
		return
	}
	fmt.Printf("Got message from client: %v\n", string(buf))
	conn.Write(buf)
	return
}

func main() {
	fmt.Println("Starting Curry server")

	curry := New(":3000")
	log.Fatalln(curry.Start())
}
