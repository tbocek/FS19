package main

import (
	//"bytes"
	"crypto/sha256"
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/hashicorp/go-version"
	"io"
	//"math/rand"
	"net"
	"strings"
	//"time"
)

func main() {
	remote, err := net.Dial("tcp", "node1.plutolo.gy:10333") //check: http://monitor.cityofzion.io/
	if err != nil {
		panic(err)
	}
	defer remote.Close()
	fmt.Println("Conneced to: %v", remote.RemoteAddr())

	payloadVersion := encodeVersion("/The HSR NEO client:0.0.1/")
	packetVersion := encodeHeader("version", payloadVersion)
	n, err := remote.Write(packetVersion)
	if err != nil {
		panic(err)
	}
	fmt.Printf("wrote version packet: %v, %d\n", packetVersion, n)

	//we get the version from the remote, display it
	read := make([]byte, 24)
	n, err = io.ReadFull(remote, read) //read header
	plen, rcvChecksum := decodeHeader(read)
	read = make([]byte, plen)
	n, err = io.ReadFull(remote, read) //read payload
	userAgent := decodeVersion(read)

	tmp := sha256.Sum256(read)
	hash := sha256.Sum256(tmp[:])
	checksum := binary.LittleEndian.Uint32(hash[0:4])
	fmt.Printf("read version payload: %v, %d\n", read, n)
	if rcvChecksum != checksum {
		panic(errors.New("checksum mismatch in version!"))
	}

	//check if we have a good version
	start := strings.Index(userAgent, ":")
	end := strings.Index(userAgent[start:], "/")
	if start < 0 && end < 0 {
		panic(errors.New(fmt.Sprintf("cannot parse version in %s", userAgent)))
	}
	semVer := userAgent[start+1 : start+end]
	fmt.Printf("parsed semver: %v\n", semVer)
	v1, err := version.NewVersion(semVer)
	min, err := version.NewVersion("2.10.1")
	if v1.LessThan(min) {
		panic(errors.New(fmt.Sprintf("%s is less than %s", v1, min)))
	}

	////////// got version, send ack
	packetVerack := encodeHeader("verack", []byte{})
	n, err = remote.Write(packetVerack)
	if err != nil {
		panic(err)
	}

	///////// wait for verack confirmation
	read = make([]byte, 24)
	n, err = io.ReadFull(remote, read)
	plen, rcvChecksum = decodeHeader(read)
	fmt.Printf("read verack array: %v, %d\n", read, plen)
	if rcvChecksum != 3806393949 {
		panic(errors.New("checksum mismatch in verack!"))
	}

	/////// send ping
	packet2 := encodeHeader("ping", encodePing())
	n, err = remote.Write(packet2)
	if err != nil {
		panic(err)
	}
	fmt.Printf("wrote ping: %v, %d\n", packet2, n)

	//////// receive pong
	read = make([]byte, 36)
	_, err = io.ReadFull(remote, read)
	_, rcvChecksum = decodeHeader(read)
	decodePing(read[24 : 24+12])
	fmt.Printf("read array: %v\n", read)

	tmp = sha256.Sum256(read[24 : 24+12])
	hash = sha256.Sum256(tmp[:])
	checksum = binary.LittleEndian.Uint32(hash[0:4])

	if rcvChecksum != checksum {
		panic(errors.New("checksum mismatch in pong!"))
	}

	remote.Close()
}

func encodeHeader(cmd string, payload []byte) []byte {
	b := make([]byte, 24+len(payload))

	//encoding here

	//payload
	copy(b[24:], payload)
	return b
}

func encodeVersion(userAgent string) []byte {
	userAgentLen := len(userAgent)
	b := make([]byte, 27+userAgentLen+1)
	// encoding here
	b[27+userAgentLen] = 0
	return b
}

func encodePing() []byte {
	b := make([]byte, 12)
	// encoding here
	return b
}

func decodeHeader(b []byte) (uint32, uint32) {
	// decoding here
	//return len, checksum

	return 0, 0
}

func decodeVersion(b []byte) string {
	// decoding here
	//return userAgent

	return ""
}

func decodePing(b []byte) {
	// decoding here
}
