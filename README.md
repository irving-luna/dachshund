# Positrace project

Service to store data from client. Server receives requests by UDP and inresponse to the client sends TCP endpoint(address and port) and sessionID. Client connects to the server by TCP and sends some data to the server. Server stores data from a client to persistent storage.

Prerequisites
- Go


Local execution

``` sh
go mod download
```

``` sh
 go run main.go
```
