package main

import (
	"encoding/hex"
	"fmt"
	"net"
	"os"
	"time"
)

//compile for Alpine with
//CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo -ldflags '-w' -o vss *.go

func main() {
	fmt.Println("connecting...")
	conn, err := net.DialTimeout("udp", "172.20.0.1:5351", 1*time.Second)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("connected.")
	n, err := conn.Write([]byte{
		0x0, 0x2, 0x0, 0x0,
		0x20, 0x0, 0x10, 0x0, //external 4096, internal 8192
		0x0, 0x0, 0x20, 0x0}) //lifetime 8192 seconds

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("wrote %d bytes", n)

	tmp := make([]byte, 512)

	n, err = conn.Read(tmp)

	if err != nil {
		fmt.Println("read error:", err)
		os.Exit(1)
	}

	str := hex.EncodeToString(tmp[:n])
	fmt.Println(str)
	fmt.Println("read %d bytes", n)

	conn.Close()
}
