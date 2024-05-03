package main

import (
	"positrace/api"
	"positrace/tcp"
	"positrace/udp"
	"positrace/usecase"
)

func main() {

	tcpServer := tcp.NewTCPServer()
	go tcpServer.Listen()

	apiServer := api.NewAPIServer("8088")
	go apiServer.Start()

	sessionUsecase := usecase.NewSessionUsecase(60, &tcpServer.Listeners)
	udpServer := udp.NewUDPServer("10000", sessionUsecase)
	go udpServer.Listen()


	select {}
}
