package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"os"
)

/*
TODOS:
	1) Define pcap file format with a struct.
	2) Define packet structure with a struct.
	3) Define HTTP, TCP, IP, Ethernet structures with structs.
	4) Write helper functions meant to parse each header:
			- HTTP
			- TCP
			- IP
			- Ethernet
*/

// Magic number of net.cap file is 0xd4c3b2a1
// This means the host that wrote the file uses the opposite
// byte order from my machine.
// Thus, we can change the byte order in Go with an endian
// function defined in the standard library.

// pcap file header that appears before packets.
// Header is 24 bytes.
// Values recorded in comments below are taken from using xxd:
// `xxd -l 32 net.cap`
type pcap_file_header struct {
	// 4-byte number -> d4c3b2a1
	magic_num int
	// 2-byte number that should equal 2
	major_version int
	// 2-byte number that should equal 4
	minor_version int
	// 4-byte value that is always 0
	time_zone_offset int
	// 4-byte value that is always 0
	time_stamp_accuracy int
	// 4-byte number -> ea05 0000
	snapshot_length int
	// 4-byte number -> 0100 0000 (little-endian -> 0000 0001)
	// This is a value of 1 which indicates LINKTYPE_ETHERNET
	link_layer_header_type int
}

// Each packet in a pcap file following the pcap
// file header will have a packet header.
// The header is 16-bytes.
type pcap_packet_header struct {
	// 4-byte value -> 57d09840
	time_stamp_large int
	// 4-byte value -> 00031f0a
	time_stamp_small int
	// 4-byte value -> 0000 004e -> 18 bytes?
	length int
	// 4-byte value -> same value as length, so not truncated?
	ut_length int
}

func main() {
	f, err := os.Open("./net.cap")
	if err != nil {
		log.Fatal(err)
	}

	// Read enough bytes to obtain the pcap file header.
	b := make([]byte, 24)
	_, err = f.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Magic number: %x\n", binary.LittleEndian.Uint32(b[0:4]))

	fmt.Printf("Major version: %x\n", binary.LittleEndian.Uint16(b[4:6]))

	fmt.Printf("Minor version: %x\n", binary.LittleEndian.Uint16(b[6:8]))

	fmt.Printf("Time zone offset: %x\n", binary.LittleEndian.Uint32(b[8:12]))

	fmt.Printf("Time stamp accuracy: %x\n", binary.LittleEndian.Uint32(b[12:16]))

	fmt.Printf("Snapshot length: %x\n", binary.LittleEndian.Uint32(b[16:20]))

	fmt.Printf("Link layer type: %x\n", binary.LittleEndian.Uint32(b[20:]))
}
