package main

import (
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"log"
	_ "os"
	_ "strconv"
)

// Magic number of net.cap file is 0xd4c3b2a1
// This means the host that wrote the file uses the opposite
// byte order from my machine.
// Thus, we can change the byte order in Go with an endian
// function defined in the standard library.

// pcap file header that appears before packets.
// Header is 24 bytes.
// Values recorded in comments below are taken from using xxd:
// `xxd -l 32 -e net.cap`
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
	// 4-byte value -> 0000 004e
	length uint32
	// 4-byte value -> same value as length, so not truncated?
	ut_length uint32
}

type ethernet_frame struct {
	// 6 byte value
	mac_dest []byte
	// 6 byte value
	mac_src []byte
	// 2 byte value; indicates encapsulated IP protocol
	// IPv4: 0x0800
	// IPv6: 0x86DD
	ethertype uint16
	// 46-1500 byte value; data payload
	payload []byte
}

type ip_header struct {
	version uint8
	ihl     uint8
	//dscp            []byte
	//ecn             []byte
	total_length []byte
	//id              []byte
	//flags           []byte
	//fragment_offset []byte
	//ttl             uint8
	protocol []byte
	//header_checksum uint16
	src_ip []byte
	dst_ip []byte
	//options         []byte
}

type tcp_header struct {
	src_port []byte
	dst_port []byte
	seq_num  []byte
	//ack_num     uint32
	data_offset uint8
	//reserved    []byte
	flags []byte
	syn   uint8
	//window_size uint16
	//checksum    uint16
	//urg_pointer uint16
	//options     []byte
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

func parseEthernetFrame(data []byte) ethernet_frame {
	ef := ethernet_frame{}

	//ef.mac_dest = binary.BigEndian.Uint64(data[0:6])
	ef.mac_dest = data[0:6]
	//ef.mac_src = binary.BigEndian.Uint64(data[6:12])
	ef.mac_src = data[6:12]
	ef.ethertype = binary.BigEndian.Uint16(data[12:14])
	//ef.ethertype = data[12:14]
	//ef.payload = binary.BigEndian.Uint64(data[14:])
	// Payload below is WRONG!
	// Should be payload = packet_header.length - (14 bytes from ethernet headers except payload)
	//ef.payload = data[14:]

	return ef
}

func parseIPHeader(data []byte) ip_header {
	ipHeader := ip_header{}

	firstByte := data[0:1]

	// Below bitwise operation obtains high order 4 bits.
	ipHeader.version = firstByte[0] >> 4
	// Below bitwise operation obtains low order 4 bits.
	ipHeader.ihl = firstByte[0] & 0x0f
	ipHeader.total_length = data[2:4]
	ipHeader.protocol = data[9:10]
	ipHeader.src_ip = data[12:16]
	ipHeader.dst_ip = data[16:20]

	return ipHeader
}

func parseTCPHeader(data []byte) tcp_header {
	tcpHeader := tcp_header{}

	tcpHeader.src_port = data[0:2]
	tcpHeader.dst_port = data[2:4]
	tcpHeader.seq_num = data[4:8]
	// Data offset is the high order 4 bits of byte.
	tcpHeader.data_offset = data[12:13][0] >> 4
	// SYN flag will be the second to last bit in byte.
	// If bit is 1 then SYN is set and packet is the first in the sequence.
	// If bit is 0 then SYN is not set and packet is not the first in the sequence.
	// Use bitwise operator to check if SYN flag is set?
	tcpHeader.flags = data[13:14]

	if (tcpHeader.flags[0] & 0x2) == 2 {
		tcpHeader.syn = 1
	} else {
		tcpHeader.syn = 0
	}

	return tcpHeader
}

func parseHTTPData(data []byte) string {

	// Combine bytes into a single binary string.
	//str := ""
	//for _, b := range data {
	//	byteArr := []byte{b}
	//	bin := binary.BigEndian.Uint16(byteArr)
	//	binStr := strconv.FormatUint(uint64(bin), 2)
	//	str += binStr
	//}

	bin := Bytes2Bits(data)
	str := fmt.Sprint(bin)
	//str := string(bin)

	// Grab HTTP status line and headers by reading up to CR LF CR LF
	// CR has ASCII value of 13 -> 0b00001101
	// LF has ASCII value of 10 -> 0b00001010

	// Grab HTTP response body containing data by getting all data after CR LF CR LF

	return str
}

func Bytes2Bits(data []byte) []int {
	dst := make([]int, 0)
	for _, v := range data {
		for i := 0; i < 8; i++ {
			move := uint(7 - i)
			dst = append(dst, int((v>>move)&1))
		}
	}
	fmt.Println(len(dst))
	return dst
}

func countPackets(data []byte) int {
	var i uint32
	length := uint32(len(data))
	num := 0

	for i = 0; i < length; {
		// Parse the data in the pcap per packet header.
		packetHeader := parsePacketHeader(data[i : i+16])

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

	// Parse packets in a loop.
	//for i := 0; i < 99; i++ {

	//httpData := make(map[int]
	// Use below rudimentary storage to just work on http parsing logic
	//	- work on ordering by sequence number later
	httpData := make([]byte, 2)
	packetStart := 24
	sliceSeqNums := make([]uint32, 99)
	//for i := 0; i < 6; i++ {
	for i := 1; i < 100; i++ {
		fmt.Printf("PACKET #%d\n\n", i)
		// Parse per-packet pcap header.
		pph := parsePacketHeader(data[packetStart:(packetStart + 16)])
		packetLength := pph.length
		fmt.Printf("pph: %#v\n\n", pph)
		fmt.Printf("packetLength: %v\n\n", packetLength)

		ethernetStart := packetStart + 16
		ethernetFrame := parseEthernetFrame(data[ethernetStart:(ethernetStart + 14)])
		fmt.Printf("ethernetFrame\nmac_dest:  %#v\nmac_src: %#v\nethertype: %#v\n\n", ethernetFrame.mac_dest, ethernetFrame.mac_src, ethernetFrame.ethertype)
		//fmt.Printf("ethernetFrame payload length: %d\n", len(ethernetFrame.payload))

		ipStart := ethernetStart + 14
		ipHeader := parseIPHeader(data[ipStart:])
		fmt.Printf("ipHeader: %#v\n\n", ipHeader)

		tcpStart := ipStart + (int(ipHeader.ihl) * 4)
		fmt.Printf("IHL: %v\n\n", ipHeader.ihl)
		tcpHeader := parseTCPHeader(data[tcpStart:])
		fmt.Printf("tcpHeader: %#v\n\n", tcpHeader)

		if tcpHeader.syn == 1 {
			//packetStart = tcpStart + ((int(tcpHeader.data_offset) * 32) / 8)
			packetStart += 16 + int(packetLength)
			fmt.Printf("loop end packetStart = %v\n\n", packetStart)
			continue
		}

		ipTotalLength := binary.BigEndian.Uint16(ipHeader.total_length)
		fmt.Printf("ipTotalLength = %v\n\n", ipTotalLength)
		seqNum := binary.BigEndian.Uint32(tcpHeader.seq_num)
		fmt.Printf("seqNum = %v\n\n", seqNum)
		sliceSeqNums = append(sliceSeqNums, seqNum)
		httpStart := tcpStart + ((int(tcpHeader.data_offset) * 32) / 8)
		//httpEnd := int(ipTotalLength) - (int(ipHeader.ihl) * 4) - httpStart
		httpEnd := httpStart + int(packetLength)
		httpData = append(httpData, data[httpStart:httpEnd]...)
		//packetStart = httpEnd
		packetStart += 16 + int(packetLength)
	}

	fmt.Printf("sliceSeqNums: %v\n\n", sliceSeqNums)

	//fmt.Printf("httpData: %#v\n\n", httpData)
	//httpDataString := parseHTTPData(httpData)
	//fmt.Printf("httpHeader: %v\n", httpDataString)

	// Write HTTP data to a file.
	//	f, err := os.Create("packet.jpg")
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	defer f.Close()
	//
	//	_, err = f.WriteString(httpDataString)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
}
