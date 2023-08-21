package usecase

import (
	"net"
	"positrace/models"
	"time"

	"github.com/google/uuid"
)

type Session struct {
	Expiration int
	TcpPool    *[]net.Listener
}

func NewSessionUsecase(exp int, pool *[]net.Listener) *Session {
	return &Session{
		Expiration: exp,
		TcpPool:    pool,
	}
}

func (s *Session) Start(client models.LoginRequest) models.Session {

	// TODO: Get active sessions from database
	tcpAddress := s.getTCPAddress()
	sessionID := uuid.New()
	expiration := time.Now().Add(time.Duration(s.Expiration) * time.Millisecond * 1000)

	session := models.Session{
		ID:          sessionID.String(),
		ClientName:  &client.ClientName,
		TCPAddrress: &tcpAddress,
		ExpiresAt:   &expiration,
	}

	// TODO: store session in database

	return session

}

func (s *Session) getTCPAddress() string {

	// TODO: dispatcher count-based balancing by count of active sessions.

	return "localhost:port"
}
