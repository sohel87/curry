package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	fmt.Println("Starting Curry server")

	listener, err := net.Listen("tcp", "0.0.0.0:6378")
	if err != nil {
		fmt.Println("Failed to bind server to port 6378")
		os.Exit(1)
	}

	_, err = listener.Accept()
	if err != nil {
		fmt.Println("Error accepting connections: ", err.Error())
		os.Exit(1)
	}
	fmt.Println("Started listening on port: 6378")
}
