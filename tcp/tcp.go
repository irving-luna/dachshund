package tcp

import (
	"fmt"
	"net"
)

type TCPServer struct {
}

func NewTCPServer() *TCPServer {
	return &TCPServer{}
}

func (u *TCPServer) Listen() error {
	// TCP listener pool
	listeners := make([]net.Listener, 5)

	for i := 0; i < 5; i++ {
		listener, _ := net.Listen("tcp", ":0")
		listeners[i] = listener
		defer listener.Close()
		port := listener.Addr().(*net.TCPAddr).Port
		fmt.Printf("TCP listener in port: %d\n", port)
	}

	for _, v := range listeners {
		for {
			conn, _ := v.Accept()
	
			buffer := make([]byte, 1024)
			conn.Read(buffer)
	
			// Read session ID from the client
			// Retrieve session details from Redis
			// Handle client's data and store transmitted bytes
		}
		
	}
	return nil
}
