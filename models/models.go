package models

import "time"

type Client struct {
	Name             string
	TransmitedBytes  int32
	SessionStartedAt time.Time
}

type LoginRequest struct {
	ClientName string
}
type Session struct {
	ID          string
	ClientName  *string
	TCPAddrress *string
	ExpiresAt   *time.Time
}
