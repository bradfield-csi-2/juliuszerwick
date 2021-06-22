package main

import (
	"fmt"
	"log"
	"os"
	"syscall"
)

/*
Steps to complete:
	1) Create a raw socket and connect to it
	2) Construct a ping request
	3) Read and print the response back for each ping
*/

func readResponse(fd int, buf []byte) (int, error) {
	n, err := syscall.Read(fd, buf)
	if err != nil {
		return 0, err
	}

	return n, nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("A destination hostname or IP must be provided, please try again.")
		os.Exit(1)
	}

	dest := os.Args[1]
	fmt.Printf("dest: %v\n", dest)

	// Create a socket to use for ping requests.
	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_RAW, 0)
	if err != nil {
		log.Fatalf("failed to create socket - err: %v\n", err)
	}

	// Port number 33434 is the default port number used by the traceroute program.
	sAddr := &syscall.SockaddrInet4{Port: 33434}
	copy(sAddr.Addr[:], dest)

	err = syscall.Connect(fd, sAddr)
	if err != nil {
		log.Fatalf("failed to connect - err: %v\n", err)
	}

	// Create buffer to store ping request?
	buf := make([]byte, 1000)

	// Write request data to connection to send to web server.
	_, err = syscall.Write(fd, buf)
	if err != nil {
		log.Fatalf("failed to write over socket - err: %v\n", err)
	}

	// Read data that was returned by router.
	buf2 := make([]byte, 1000)
	_, err = readResponse(fd, buf2)
	if err != nil {
		log.Fatalf("failed to read response over socket - err: %v\n", err)
	}
	fmt.Printf("Reading response over socket...\n\n")
}
