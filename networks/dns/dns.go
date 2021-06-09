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
		7) Refactor code to allow user to specify the domain name and query type (turn it into a CLI tool).
*/

import "fmt"

var (
	domainName = "https://en.wikipedia.org/wiki/Main_Page"
	queryType  = "A"
)

type dns_message struct {
	header     header
	question   question
	answer     []resource_record
	authority  []resource_record
	additional []resource_record
}

type header struct {
	id      uint16
	qr      uint8
	opcode  uint8
	aa      uint8
	tc      uint8
	rd      uint8
	ra      uint8
	z       uint8
	rcode   uint8
	qdcount uint16
	ancount uint16
	nscount uint16
	arcount uint16
}

type question struct {
	qname  uint32
	qtype  uint16
	qclass uint16
}

type resource_record struct {
	name     uint32
	typ      uint16
	class    uint16
	ttl      uint32
	rdlength uint16
	rdata    uint32
}

func main() {
	fmt.Printf("Welcome to our DNS client!\n\n")
}
