package udp

import (
	"fmt"
	"net"
)

type UDPServer struct {
	Port string
}

func NewUDPServer(port string) *UDPServer {
	return &UDPServer{
		Port: port,
	}
}

func (u *UDPServer) Listen() error {
	// UDP listener
	udpAddr, err := net.ResolveUDPAddr("udp", fmt.Sprintf(":%s", u.Port))
	if err != nil {
		return err
	}
	udpConn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		return err
	}

	fmt.Printf("UDP server listening in port: %s\n", u.Port)
 
	for {
		buffer := make([]byte, 1024)
		udpConn.ReadFromUDP(buffer)

		// TODO: Handle UDP login request, generate session ID, and respond with TCP address:port
		// TODO: Store session ID in Redis
	}

}
