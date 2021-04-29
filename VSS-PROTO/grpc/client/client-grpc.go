package main

//go get -u google.golang.org/grpc
//sudo apt install protobuf-compiler
//go get github.com/golang/protobuf/protoc-gen-go
//wget https://github.com/protocolbuffers/protobuf/releases/download/v3.7.1/protobuf-all-3.7.1.zip
//unzip
//protoc -I/tmp/protobuf-3.7.1/src -I. --go_out=plugins=grpc:. grpc/*.grpc

//go build json/client/*;go build json/server/*;go build grpc/server/*;go build grpc/client/*
//scp -P 2222 *-* root@localhost:/tmp

//R1=`cat /sys/class/net/nat_wan/statistics/rx_bytes`;T1=`cat /sys/class/net/nat_wan/statistics/tx_bytes`;/tmp/client-
//json;R2=`cat /sys/class/net/nat_wan/statistics/rx_bytes`;T2=`cat /sys/class/net/nat_wan/statistics/tx_bytes`;TXPPS=`expr $T2 - $
//T1`;RXPPS=`expr $R2 - $R1`;echo "TX $1: $TXPPS bytes RX $1: $RXPPS bytes"

import (
	"context"
	"fmt"
	pb "github.com/tbocek/VSS-PROTO/grpc"
	"google.golang.org/grpc"
	"time"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial("10.0.2.16:4000", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	u := pb.User{}
	u.Username = "Thomas"
	u.Password = "hsr12345"
	u2 := &pb.User{}

	start := time.Now()
	for i := 0; i < 100; i++ {
		c := pb.NewUserServiceClient(conn)

		u2, err = c.UserRPC(context.Background(), &u)
		if err != nil {
			panic(err)
		}
	}
	elapsed := time.Since(start)
	fmt.Printf("%+v, took %s\n", u2, elapsed)
}
