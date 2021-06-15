package main

import (
	"fmt"
	"log"
	"net"
	"net/textproto"
	"syscall"
)

var (
	clientPort int    = 8000
	serverPort int    = 9000 //8001
	localHost  net.IP = net.ParseIP("127.0.0.1")
)

//type netSocket struct {
//	fd int
//}

type request struct {
	method string
	header textproto.MIMEHeader
	body   []byte
	uri    string
	proto  string
}

//func (ns netSocket) Read(p []byte) (int, error) {
//	return 0, nil
//}

//func (ns netSocket) Write(p []byte) (int, error) {
//	return 0, nil
//}

//func (ns *netSocket) Close() error {
//	return syscall.Close(ns.fd)
//}

func readRequest(fd int, buf []byte) (int, error) {
	n, err := syscall.Read(fd, buf)
	if err != nil {
		return 0, err
	}

	return n, nil
}

func main() {
	// Create a socket to use for receiving pending connections from clients.
	fdClient, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, 0)
	if err != nil {
		log.Fatalf("failed to create socket - err: %v\n", err)
	}
	defer syscall.Close(fdClient)

	// Bind to socket for pending connections from clients.
	sAddrClient := &syscall.SockaddrInet4{Port: clientPort}
	copy(sAddrClient.Addr[:], localHost)
	if err = syscall.Bind(fdClient, sAddrClient); err != nil {
		log.Fatalf("failed to bind to socket on port %d - err: %v\n", clientPort, err)
	}

	// Listen for incoming connections from clients on socket.
	if err = syscall.Listen(fdClient, syscall.SOMAXCONN); err != nil {
		log.Fatalf("failed to listen on socket - err: %v\n", err)
	}

	fmt.Printf("=======================\n")
	fmt.Printf("Proxy is started!\n")
	fmt.Printf("=======================\n\n")

	// Wait for incoming connections.
	for {
		// Accept incoming connection from a client.
		nfdClient, _, err := syscall.Accept(fdClient)
		if err != nil {
			log.Fatalf("failed to accept connection - err: %v\n", err)
		}
		fmt.Printf("Accepted connection!\n\n")

		// Read request from client and store into a byte array.
		// Allocate enough bytes to receive data sent over socket.
		buf := make([]byte, 1000)
		r, err := readRequest(nfdClient, buf)
		if err != nil {
			log.Fatalf("failed to read request - err: %v\n", err)
		}
		fmt.Printf("Reading request from client...\n\n")

		// Connect to web server.
		fmt.Printf("Writing to web server...\n\n")
		// Create addr for outbound connections to the web server.
		sAddrServer := &syscall.SockaddrInet4{Port: serverPort}
		copy(sAddrServer.Addr[:], localHost)

		// Create a socket to use for outbound connections to the web server.
		fdServer, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, 0)
		if err != nil {
			log.Fatalf("failed to create socket - err: %v\n", err)
		}

		err = syscall.Connect(fdServer, sAddrServer)
		if err != nil {
			log.Fatalf("failed to connect to web server - err: %v\n", err)
		}

		// Write request data to connection to send to web server.
		_, err = syscall.Write(fdServer, buf[:r])
		if err != nil {
			log.Fatalf("failed to write response - err: %v\n", err)
		}

		// Read data send back by server.
		buf2 := make([]byte, 1000)
		r2, err := readRequest(fdServer, buf2)
		if err != nil {
			log.Fatalf("failed to read response from web server - err: %v\n", err)
		}
		fmt.Printf("Reading response from web server...\n\n")

		//fmt.Printf("%v\n\n", buf2)
		//fmt.Printf("%s\n\n", string(buf2))

		// Write response data to connection to send to client.
		_, err = syscall.Write(nfdClient, buf2[:r2])
		if err != nil {
			log.Fatalf("failed to write response to client - err: %v\n", err)
		}

		fmt.Printf("data sent back to client\n")

		syscall.Close(fdServer)
	}
}
