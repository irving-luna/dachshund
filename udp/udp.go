package udp

import (
	"encoding/json"
	"fmt"
	"net"
	"positrace/models"
	"positrace/usecase"
)

type UDPServer struct {
	Port    string
	Session *usecase.Session
}

func NewUDPServer(port string, session *usecase.Session) *UDPServer {
	return &UDPServer{
		Port:    port,
		Session: session,
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

	buffer := make([]byte, 1024)
	for {
		n, _, err := udpConn.ReadFrom(buffer)
		if err != nil {
			return err
		}

		var lr models.LoginRequest
		if err = json.Unmarshal(buffer[:n], &lr); err != nil {
			continue
		}

		u.Session.Start(lr)

	}

}
