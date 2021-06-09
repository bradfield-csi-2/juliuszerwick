package main

/*
	This is a DNS client that constructs a DNS request and transmits it through a socket
	to ensure the proper encapsulation by the transport, network, and link layer protocols.
	It will also parse the response back from the DNS server.

	TODOS:
		1) Write a struct to represent the DNS query.
		2) Write a struct to represent the DNS response.
		3) Add logic to construct the DNS query and print to verify.
		4) Add logic to create a datagram socket and write the DNS query to.
		5) Add logic to receive the response and parse it into the DNS response struct.
		6) Print the DNS response to verify.
*/

import "fmt"

func main() {
	fmt.Println("Welcome to our DNS client!\n")
}
