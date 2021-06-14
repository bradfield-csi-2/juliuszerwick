package main

import (
	"fmt"
	"log"
	"net"
	"net/textproto"
	"syscall"
)

var (
	proxyPort int    = 8000
	localHost net.IP = net.ParseIP("127.0.0.1")
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
	// Create a socket.
	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, 0)
	if err != nil {
		log.Fatalf("failed to create socket - err: %v\n", err)
	}
	defer syscall.Close(fd)

	// Bind to socket.
	sAddr := &syscall.SockaddrInet4{Port: proxyPort}
	copy(sAddr.Addr[:], localHost)
	if err = syscall.Bind(fd, sAddr); err != nil {
		log.Fatalf("failed to bind to socket - err: %v\n", err)
	}

	// Listen for connections on socket.
	if err = syscall.Listen(fd, syscall.SOMAXCONN); err != nil {
		log.Fatalf("failed to listen on socket - err: %v\n", err)
	}

	fmt.Printf("=======================\n")
	fmt.Printf("Server is started!\n")
	fmt.Printf("=======================\n\n")

	// Wait for incoming connections.
	for {
		// Accept connection.
		nfd, _, err := syscall.Accept(fd)
		if err != nil {
			log.Fatalf("failed to accept connection - err: %v\n", err)
		}
		fmt.Printf("Accepted connection!\n\n")

		// Read request and print.
		// Allocate enough bytes to receive data sent over socket.
		buf := make([]byte, 1000)
		r, err := readRequest(nfd, buf)
		if err != nil {
			log.Fatalf("failed to read request - err: %v\n", err)
		}
		fmt.Printf("Read request...\n\n")

		// Write response.
		fmt.Printf("Writing response...\n\n")
		// Echo back bytes received from start of byte array up to number of bytes read.
		_, err = syscall.Write(nfd, buf[:r])
		if err != nil {
			log.Fatalf("failed to write response - err: %v\n", err)
		}
	}
}
