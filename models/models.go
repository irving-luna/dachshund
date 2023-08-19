package models

import "time"

type Client struct{
	Name string
	TransmitedBytes int32
	SessionStartedAt time.Time
}