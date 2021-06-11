package main

import (
	"fmt"
	"log"
	"net"

	"github.com/pkg/errors"
)

var (
	domainName = "https://en.wikipedia.org/wiki/Main_Page"
	queryType  = "A"
)

type dns_message struct {
	header   header
	question question
	answer   []*rr
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
			id: 1,
			// Value of 0 specifies a query.
			qr: 0,
			// Value of 0 specifies a standard query.
			opcode: 0,
			// Number of entries in the question section.
			qdcount: 1,
			// Value of 1 indicates that we desire a recursive query.
			rd: 1,
		},
		question: {
			qname: domainName,
			// QYTPE A - host address - value of 1
			qtype: 1,
			// QCLASS IN - Internet - value of 1
			qclass: 1,
		},
	}

	// Creating a socket for UDP and IPv4.
	c, err := net.Dial("udp4", domainName)
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	payload, err := dnsMsg.Marshal()
	if err != nil {
		err = errors.Wrapf(err, "failure to marshal dnsMsg into paylaod %+v", dnsMsg)
		log.Fatal(err)
	}

	fmt.Printf("dnsMsg sent %s\n", string(payload))

	// Write DNS request to Conn instance c.
	_, err := c.Write()
	if err != nil {
		err = errors.Wrapf(err, "failure to write payload to conn%+v", dnsMsg)
		log.Fatal(err)
	}

	// Read DNS response from Conn instance c.
	resp, err := ioutil.ReadAll(c)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(resp))
}
