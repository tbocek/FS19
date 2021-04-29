package main

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	bf "github.com/tbocek/VSS-PROTO/bloomfilter"
	"github.com/willf/bloom"
	"math/rand"
	"net/http"
)

func main() {

	var a []int64
	for i := 0; i < 1000; i++ {
		a[i] = rand.Int63n(2000)
	}
	filter := bloom.New(100, 5)

	for i := 0; i < 1000; i++ {
		n1 := make([]byte, 8)
		binary.BigEndian.PutUint64(n1, uint64(a[i]))
		filter.Add(n1)
	}

	jsonBytes, err := filter.MarshalJSON()
	req, err := http.NewRequest(http.MethodGet, "http://10.0.2.16:5000/nr", bytes.NewBuffer(jsonBytes))
	if err != nil {
		panic(err)
	}
	client := http.DefaultClient

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	nr := bf.Numbers{}

	err = json.NewDecoder(resp.Body).Decode(&nr)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v, took %s\n", nr)
}
