package main

import (
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"log"
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
	magic_num uint32
	// 2-byte number that should equal 2
	major_version uint16
	// 2-byte number that should equal 4
	minor_version uint16
	// 4-byte value that is always 0
	time_zone_offset uint32
	// 4-byte value that is always 0
	time_stamp_accuracy uint32
	// 4-byte number -> ea05 0000
	snapshot_length uint32
	// 4-byte number -> 0100 0000 (little-endian -> 0000 0001)
	// This is a value of 1 which indicates LINKTYPE_ETHERNET
	link_layer_header_type uint32
}

// Each packet in a pcap file following the pcap
// file header will have a packet header.
// The header is 16-bytes.
type pcap_packet_header struct {
	// 4-byte value -> 57d09840
	time_stamp_large uint32
	// 4-byte value -> 00031f0a
	time_stamp_small uint32
	// 4-byte value -> 0000 004e -> 18 bytes?
	length uint32
	// 4-byte value -> same value as length, so not truncated?
	ut_length uint32
}

func parsePcapHeader(data []byte) pcap_file_header {
	ph := pcap_file_header{}

	ph.magic_num = binary.BigEndian.Uint32(data[0:4])
	ph.major_version = binary.LittleEndian.Uint16(data[4:6])
	ph.minor_version = binary.LittleEndian.Uint16(data[6:8])
	ph.time_zone_offset = binary.LittleEndian.Uint32(data[8:12])
	ph.time_stamp_accuracy = binary.LittleEndian.Uint32(data[12:16])
	ph.snapshot_length = binary.LittleEndian.Uint32(data[16:20])
	ph.link_layer_header_type = binary.LittleEndian.Uint32(data[20:24])

	return ph
}

func parsePacketHeader(data []byte) pcap_packet_header {
	pph := pcap_packet_header{}

	pph.time_stamp_large = binary.LittleEndian.Uint32(data[0:4])
	pph.time_stamp_small = binary.LittleEndian.Uint32(data[4:8])
	pph.length = binary.LittleEndian.Uint32(data[8:12])
	pph.ut_length = binary.LittleEndian.Uint32(data[12:16])

	return pph
}

func countPackets(data []byte) int {
	var i uint32
	length := uint32(len(data))
	num := 0

	for i = 0; i < length; {
		// Parse the data in the pcap per packet header.
		packetHeader := parsePacketHeader(data[i : i+16])
		//fmt.Printf("packet_header: %#v\n\n", packetHeader)

		// Verify that packet lengths are the same.
		if packetHeader.length != packetHeader.ut_length {
			fmt.Printf("Packet lengths not equal!\nlength = %d\nut_length = %d\n", packetHeader.length, packetHeader.ut_length)
		}

		i += packetHeader.length + 16
		num += 1
	}

	return num
}

func main() {
	// Read the entire contents of the file into a byte array.
	data, err := ioutil.ReadFile("./net.cap")
	if err != nil {
		log.Fatal(err)
	}

	// Parse and store the data in the pcap file header.
	pcapHeader := parsePcapHeader(data[0:24])
	fmt.Printf("pcap_file_header: %#v\n\n", pcapHeader)

	// Count packets in file.
	numPackets := countPackets(data[24:])
	if numPackets == 99 {
		fmt.Printf("99 packets in file!\n\n")
	} else {
		fmt.Printf("ERROR: Only counted %d packets in file.\n\n", numPackets)
	}
}
