package main

import (
	"fmt"
	"log"
	"net"
)

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

	dnsMsg := dns_message{
		header: {
			id:     uint16(1),
			qr:     uint8(0),
			opcode: 0,
		},
	}

	// Creating a socket for UDP and IPv4.
	c, err := net.Dial("udp4", domainName)
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	// Write DNS request to Conn instance c.
	_, err := c.Write()
	if err != nil {
		log.Fatal(err)
	}

	// Read DNS response from Conn instance c.
	resp, err := ioutil.ReadAll(c)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(resp))
}
