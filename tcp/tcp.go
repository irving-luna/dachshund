package tcp

import (
	"fmt"
	"net"
)

type TCPServer struct {
	Listeners []net.Listener
}

func NewTCPServer() *TCPServer {
	return &TCPServer{
		Listeners: make([]net.Listener, 5),
	}
}

func (u *TCPServer) Listen() error {

	for i := 0; i < 5; i++ {
		listener, _ := net.Listen("tcp", ":0")
		u.Listeners[i] = listener
		defer listener.Close()
		port := listener.Addr().(*net.TCPAddr).Port
		fmt.Printf("TCP listener in port: %d\n", port)
	}

	for _, v := range u.Listeners {
		for {
			conn, _ := v.Accept()
	
			buffer := make([]byte, 1024)
			conn.Read(buffer)
	
			// Read session ID from the client
			// Retrieve session details 
			// Handle client's data and store transmitted bytes
		}
		
	}
	return nil
}
