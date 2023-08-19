package main

import (
	"positrace/api"
	"positrace/tcp"
	"positrace/udp"
)

func main() {
	udpServer := udp.NewUDPServer("10000")
	go udpServer.Listen()

	tcpServer := tcp.NewTCPServer()
	go tcpServer.Listen()

	apiServer := api.NewAPIServer("8088")
	go apiServer.Start()

	select {}
}
